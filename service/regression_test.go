package service

import (
	"context"
	"github.com/SidharthSasikumar/train-ticket-grpc/ticket"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestInvalidRequestsAndDetachedReceipts(t *testing.T) {
	s := &Server{}
	ctx := context.Background()
	if _, err := s.PurchaseTicket(ctx, nil); status.Code(err) != codes.InvalidArgument {
		t.Fatal(err)
	}
	if _, err := s.PurchaseTicket(ctx, &ticket.PurchaseRequest{}); status.Code(err) != codes.InvalidArgument {
		t.Fatal(err)
	}
	req := &ticket.PurchaseRequest{User: &ticket.User{Email: "test@example.com", FirstName: "Test"}}
	got, err := s.PurchaseTicket(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	req.User.FirstName = "mutated"
	got.Receipt.Seat = "X999"
	read, err := s.GetReceipt(ctx, &ticket.GetReceiptRequest{Email: "test@example.com"})
	if err != nil {
		t.Fatal(err)
	}
	if read.Receipt[0].Seat != "A1" || read.Receipt[0].User.FirstName != "Test" {
		t.Fatal("caller mutated stored receipt")
	}
	read.Receipt[0].Seat = "B50"
	for _, seat := range []string{"", "X1", "A0", "A51", "A01"} {
		if _, err := s.ModifySeat(ctx, &ticket.ModifySeatRequest{Email: "test@example.com", NewSeat: seat}); status.Code(err) != codes.InvalidArgument {
			t.Fatalf("accepted %q: %v", seat, err)
		}
	}
	read, _ = s.GetReceipt(ctx, &ticket.GetReceiptRequest{Email: "test@example.com"})
	if read.Receipt[0].Seat != "A1" {
		t.Fatal("invalid update changed seat")
	}
}
