package main

import (
	"context"
	"fmt"
	"github.com/SidharthSasikumar/train-ticket-grpc/service"
	"github.com/SidharthSasikumar/train-ticket-grpc/ticket"
	"github.com/SidharthSasikumar/train-ticket-grpc/utils"
	"testing"
)

func TestPurchaseTicket(t *testing.T) {
	s := &service.Server{
		Users: make(map[string][]*ticket.Receipt),
		Seats: make(map[string]string),
	}

	req := &ticket.PurchaseRequest{
		From: "London",
		To:   "France",
		User: &ticket.User{
			FirstName: "Sidharth1",
			LastName:  "Sasikumar",
			Email:     "sidharth.sasikumar@example.com",
		},
	}

	res, err := s.PurchaseTicket(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.Receipt.Seat == "" {
		t.Fatalf("Expected seat allocation, got empty seat")
	}

	// Check that the user has 1 receipt
	if len(s.Users[req.User.Email]) != 1 {
		t.Fatalf("Expected 1 receipt for user, got %d", len(s.Users[req.User.Email]))
	}
	req2 := &ticket.PurchaseRequest{
		From: "London",
		To:   "France",
		User: &ticket.User{
			FirstName: "Sidharth2",
			LastName:  "Sasikumar",
			Email:     "sidharth.sasikumar@example.com",
		},
	}

	// Purchase another ticket for the same user
	res, err = s.PurchaseTicket(context.Background(), req2)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check that the user now has 2 receipts
	if len(s.Users[req.User.Email]) != 2 {
		t.Fatalf("Expected 2 receipts for user, got %d", len(s.Users[req.User.Email]))
	}
}

const (
	totalSeatsA = 50 // number of seats in Section A
	totalSeatsB = 50 // number of seats in Section B
)

func TestAllocateSeat(t *testing.T) {
	s := &service.Server{
		Users: make(map[string][]*ticket.Receipt),
		Seats: make(map[string]string),
	}

	// Allocate all seats in Section A
	for i := 1; i <= totalSeatsA; i++ {
		seat, err := utils.AllocateSeat(s.Seats)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		expected := fmt.Sprintf("A%d", i)
		if seat != expected {
			t.Fatalf("Expected seat %v, got %v", expected, seat)
		}
		fmt.Printf("Expected Seat Number %s Got %s\n", expected, seat)
		s.Seats[seat] = fmt.Sprintf("user%d@example.com", i) // Mark the seat as taken
	}

	// Allocate all seats in Section B
	for i := 1; i <= totalSeatsB; i++ {
		seat, err := utils.AllocateSeat(s.Seats)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		expected := fmt.Sprintf("B%d", i)
		if seat != expected {
			t.Fatalf("Expected seat %v, got %v", expected, seat)
		}
		fmt.Printf("Expected Seat Number %s Got %s\n", expected, seat)
		s.Seats[seat] = fmt.Sprintf("user%d@example.com", totalSeatsA+i) // Mark the seat as taken
	}

	// Ensure no seats are left
	_, err := utils.AllocateSeat(s.Seats)
	fmt.Printf("Expected: no seats available Got: %s\n", err)
	if err == nil {
		t.Fatalf("Expected an error when all seats are taken")
	}
}

func TestGetReceipt(t *testing.T) {
	// Initialize the server with a sample receipt stored
	s := &service.Server{
		Users: make(map[string][]*ticket.Receipt),
		Seats: make(map[string]string),
	}

	// Sample receipt to be stored in the server
	sampleReceipt := &ticket.Receipt{
		From: "London",
		To:   "France",
		User: &ticket.User{
			FirstName: "Sidharth2",
			LastName:  "Sasikumar",
			Email:     "sidharth.sasikumar@example.com",
		},
		PricePaid: 20,
		Seat:      "A1",
	}

	// Store the sample receipt in the server
	s.Users[sampleReceipt.User.Email] = []*ticket.Receipt{sampleReceipt}

	// Create a GetReceiptRequest with the email of the stored receipt
	req := &ticket.GetReceiptRequest{
		Email: "sidharth.sasikumar@example.com",
	}

	// Call the GetReceipt method
	res, err := s.GetReceipt(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check that the returned receipts match the stored one
	if len(res.Receipt) != 1 {
		t.Fatalf("Expected 1 receipt, got %d", len(res.Receipt))
	}

	// Validate the details of the returned receipt
	receipt := res.Receipt[0] // Get the first (and only) receipt

	if receipt.From != sampleReceipt.From {
		t.Errorf("Expected From %v, got %v", sampleReceipt.From, receipt.From)
	}
	if receipt.To != sampleReceipt.To {
		t.Errorf("Expected To %v, got %v", sampleReceipt.To, receipt.To)
	}
	if receipt.User.Email != sampleReceipt.User.Email {
		t.Errorf("Expected Email %v, got %v", sampleReceipt.User.Email, receipt.User.Email)
	}
	if receipt.Seat != sampleReceipt.Seat {
		t.Errorf("Expected Seat %v, got %v", sampleReceipt.Seat, receipt.Seat)
	}
}

func TestViewUsers(t *testing.T) {
	s := &service.Server{
		Users: make(map[string][]*ticket.Receipt),
		Seats: make(map[string]string),
	}

	// Add some users
	s.Users["sidharth.sasikumar@example.com"] = []*ticket.Receipt{
		{
			From: "London",
			To:   "France",
			User: &ticket.User{
				FirstName: "Sidharth3",
				LastName:  "Sasikumar",
				Email:     "sidharth.sasikumar@example.com",
			},
			PricePaid: 20,
			Seat:      "A1",
		},
	}
	s.Seats["A1"] = "sidharth.sasikumar@example.com"

	s.Users["jane.smith@example.com"] = []*ticket.Receipt{
		{
			From: "London",
			To:   "France",
			User: &ticket.User{
				FirstName: "Jane",
				LastName:  "Smith",
				Email:     "jane.smith@example.com",
			},
			PricePaid: 20,
			Seat:      "B1",
		},
	}
	s.Seats["B1"] = "jane.smith@example.com"

	// Request to view users in Section A
	req := &ticket.ViewUsersRequest{
		Section: "A",
	}

	res, err := s.ViewUsers(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(res.Users) != 1 || res.Users[0].Seat != "A1" {
		t.Fatalf("Expected 1 user in Section A with seat A1, got %v", res.Users)
	}

	// Request to view users in Section B
	req.Section = "B"
	res, err = s.ViewUsers(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(res.Users) != 1 || res.Users[0].Seat != "B1" {
		t.Fatalf("Expected 1 user in Section B with seat B1, got %v", res.Users)
	}
}

func TestRemoveUser(t *testing.T) {
	s := &service.Server{
		Users: make(map[string][]*ticket.Receipt),
		Seats: make(map[string]string),
	}

	// Add a user
	s.Users["sidharth.sasikumar@example.com"] = []*ticket.Receipt{
		{
			From: "London",
			To:   "France",
			User: &ticket.User{
				FirstName: "Sidharth4",
				LastName:  "Sasikumar",
				Email:     "sidharth.sasikumar@example.com",
			},
			PricePaid: 20,
			Seat:      "A1",
		},
	}
	s.Seats["A1"] = "sidharth.sasikumar@example.com"

	fmt.Printf("Current Seat allocation %s\n", s.Seats)
	// Remove the user
	req := &ticket.RemoveUserRequest{
		Email: "sidharth.sasikumar@example.com",
	}

	res, err := s.RemoveUser(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if _, exists := s.Users["sidharth.sasikumar@example.com"]; exists {
		t.Fatalf("Expected user to be removed, but still exists")
	}

	if _, seatTaken := s.Seats["A1"]; seatTaken {
		t.Fatalf("Expected seat A1 to be freed, but it is still taken")
	}

	if res.Message != "User with email sidharth.sasikumar@example.com has been removed" {
		t.Fatalf("Unexpected remove message: %v", res.Message)
	}
	fmt.Printf("After Test Seat allocation %s\n", s.Seats)

}

func TestModifySeat(t *testing.T) {
	s := &service.Server{
		Users: make(map[string][]*ticket.Receipt),
		Seats: make(map[string]string),
	}

	// Add a user
	s.Users["sidharth.sasikumar@example.com"] = []*ticket.Receipt{
		{
			From: "London",
			To:   "France",
			User: &ticket.User{
				FirstName: "Sidharth5",
				LastName:  "Sasikumar",
				Email:     "sidharth.sasikumar@example.com",
			},
			PricePaid: 20,
			Seat:      "A1",
		},
	}
	s.Seats["A1"] = "sidharth.sasikumar@example.com"
	fmt.Printf("Current Seat allocation %s\n", s.Seats)
	// Modify the seat
	req := &ticket.ModifySeatRequest{
		Email:   "sidharth.sasikumar@example.com",
		NewSeat: "B1",
	}

	res, err := s.ModifySeat(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Retrieve the user's most recent receipt
	receipts := s.Users["sidharth.sasikumar@example.com"]
	if receipts[len(receipts)-1].Seat != "B1" {
		t.Fatalf("Expected seat to be updated to B1, but got %v", receipts[len(receipts)-1].Seat)
	}

	if _, seatTaken := s.Seats["A1"]; seatTaken {
		t.Fatalf("Expected old seat A1 to be freed, but it is still taken")
	}

	if s.Seats["B1"] != "sidharth.sasikumar@example.com" {
		t.Fatalf("Expected new seat B1 to be taken by sidharth.sasikumar@example.com, but got %v", s.Seats["B1"])
	}

	if res.Message != "User with email sidharth.sasikumar@example.com has been moved to seat B1" {
		t.Fatalf("Unexpected modify message: %v", res.Message)
	}
	fmt.Printf("After Test Seat allocation %s\n", s.Seats)
}

func TestRemoveUserAndViewUsers(t *testing.T) {
	s := &service.Server{
		Users: make(map[string][]*ticket.Receipt),
		Seats: make(map[string]string),
	}

	// Add a user with a ticket in Section A
	s.Users["sidharth.sasikumar@example.com"] = []*ticket.Receipt{
		{
			From: "London",
			To:   "France",
			User: &ticket.User{
				FirstName: "Sidharth6",
				LastName:  "Sasikumar",
				Email:     "sidharth.sasikumar@example.com",
			},
			PricePaid: 20,
			Seat:      "A1",
		},
	}
	s.Seats["A1"] = "sidharth.sasikumar@example.com"

	// Add another user in Section B
	s.Users["jane.smith@example.com"] = []*ticket.Receipt{
		{
			From: "London",
			To:   "France",
			User: &ticket.User{
				FirstName: "Jane",
				LastName:  "Smith",
				Email:     "jane.smith@example.com",
			},
			PricePaid: 20,
			Seat:      "B1",
		},
	}
	s.Seats["B1"] = "jane.smith@example.com"

	// Remove the user in Section A
	removeReq := &ticket.RemoveUserRequest{
		Email: "sidharth.sasikumar@example.com",
	}

	removeRes, err := s.RemoveUser(context.Background(), removeReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if removeRes.Message != "User with email sidharth.sasikumar@example.com has been removed" {
		t.Fatalf("Unexpected remove message: %v", removeRes.Message)
	}

	// View users in Section A to ensure the user is removed
	viewReq := &ticket.ViewUsersRequest{
		Section: "A",
	}

	viewRes, err := s.ViewUsers(context.Background(), viewReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(viewRes.Users) != 0 {
		t.Fatalf("Expected no users in Section A after removal, but got %d", len(viewRes.Users))
	}

	// View users in Section B to ensure the other user still exists
	viewReq.Section = "B"
	viewRes, err = s.ViewUsers(context.Background(), viewReq)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(viewRes.Users) != 1 || viewRes.Users[0].Seat != "B1" {
		t.Fatalf("Expected 1 user in Section B with seat B1, but got %v", viewRes.Users)
	}
}
