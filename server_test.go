package main

import (
	"context"
	"fmt"
	"github.com/SidharthSasikumar/train-ticket-grpc/ticket"
	"testing"
)

func TestPurchaseTicket(t *testing.T) {
	s := &server{
		users: make(map[string]*ticket.Receipt),
		seats: make(map[string]string),
	}

	req := &ticket.PurchaseRequest{
		From: "London",
		To:   "France",
		User: &ticket.User{
			FirstName: "Sidharth",
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
}

func TestAllocateSeat(t *testing.T) {
	s := &server{
		seats: make(map[string]string),
	}

	// Allocate all seats in Section A
	for i := 1; i <= totalSeatsA; i++ {
		seat, err := s.allocateSeat()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		expected := fmt.Sprintf("A%d", i)
		if seat != expected {
			t.Fatalf("Expected seat %v, got %v", expected, seat)
		}
		fmt.Printf("Expected Seat Number %s Got %s\n", expected, seat)
		s.seats[seat] = fmt.Sprintf("user%d@example.com", i) // Mark the seat as taken
	}

	// Allocate all seats in Section B
	for i := 1; i <= totalSeatsB; i++ {
		seat, err := s.allocateSeat()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		expected := fmt.Sprintf("B%d", i)
		if seat != expected {
			t.Fatalf("Expected seat %v, got %v", expected, seat)
		}
		fmt.Printf("Expected Seat Number %s Got %s\n", expected, seat)
		s.seats[seat] = fmt.Sprintf("user%d@example.com", totalSeatsA+i) // Mark the seat as taken
	}

	// Ensure no seats are left
	_, err := s.allocateSeat()
	fmt.Printf("Expected: no seats available Got: %s\n", err)
	if err == nil {
		t.Fatalf("Expected an error when all seats are taken")
	}
}

func TestGetReceipt(t *testing.T) {
	// Initialize the server with a sample receipt stored
	s := &server{
		users: make(map[string]*ticket.Receipt),
		seats: make(map[string]string),
	}

	// Sample receipt to be stored in the server
	sampleReceipt := &ticket.Receipt{
		From: "London",
		To:   "France",
		User: &ticket.User{
			FirstName: "Sidharth",
			LastName:  "Sasikumar",
			Email:     "sidharth.sasikumar@example.com",
		},
		PricePaid: 20,
		Seat:      "A1",
	}

	// Store the sample receipt in the server
	s.users[sampleReceipt.User.Email] = sampleReceipt

	// Create a GetReceiptRequest with the email of the stored receipt
	req := &ticket.GetReceiptRequest{
		Email: "sidharth.sasikumar@example.com",
	}

	// Call the GetReceipt method
	res, err := s.GetReceipt(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check that the returned receipt matches the stored one
	if res.Receipt.From != sampleReceipt.From {
		t.Errorf("Expected From %v, got %v", sampleReceipt.From, res.Receipt.From)
	}
	if res.Receipt.To != sampleReceipt.To {
		t.Errorf("Expected To %v, got %v", sampleReceipt.To, res.Receipt.To)
	}
	if res.Receipt.User.Email != sampleReceipt.User.Email {
		t.Errorf("Expected Email %v, got %v", sampleReceipt.User.Email, res.Receipt.User.Email)
	}
	if res.Receipt.Seat != sampleReceipt.Seat {
		t.Errorf("Expected Seat %v, got %v", sampleReceipt.Seat, res.Receipt.Seat)
	}
}

func TestViewUsers(t *testing.T) {
	s := &server{
		users: make(map[string]*ticket.Receipt),
		seats: make(map[string]string),
	}

	// Add some users
	s.users["sidharth.sasikumar@example.com"] = &ticket.Receipt{
		From: "London",
		To:   "France",
		User: &ticket.User{
			FirstName: "Sidharth",
			LastName:  "Sasikumar",
			Email:     "sidharth.sasikumar@example.com",
		},
		PricePaid: 20,
		Seat:      "A1",
	}
	s.seats["A1"] = "sidharth.sasikumar@example.com"

	s.users["jane.smith@example.com"] = &ticket.Receipt{
		From: "London",
		To:   "France",
		User: &ticket.User{
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane.smith@example.com",
		},
		PricePaid: 20,
		Seat:      "B1",
	}
	s.seats["B1"] = "jane.smith@example.com"

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
	s := &server{
		users: make(map[string]*ticket.Receipt),
		seats: make(map[string]string),
	}

	// Add a user
	s.users["sidharth.sasikumar@example.com"] = &ticket.Receipt{
		From: "London",
		To:   "France",
		User: &ticket.User{
			FirstName: "Sidharth",
			LastName:  "Sasikumar",
			Email:     "sidharth.sasikumar@example.com",
		},
		PricePaid: 20,
		Seat:      "A1",
	}
	s.seats["A1"] = "sidharth.sasikumar@example.com"

	fmt.Printf("Current Seat allocation %s\n", s.seats)
	// Remove the user
	req := &ticket.RemoveUserRequest{
		Email: "sidharth.sasikumar@example.com",
	}

	res, err := s.RemoveUser(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if _, exists := s.users["sidharth.sasikumar@example.com"]; exists {
		t.Fatalf("Expected user to be removed, but still exists")
	}

	if _, seatTaken := s.seats["A1"]; seatTaken {
		t.Fatalf("Expected seat A1 to be freed, but it is still taken")
	}

	if res.Message != "User with email sidharth.sasikumar@example.com has been removed" {
		t.Fatalf("Unexpected remove message: %v", res.Message)
	}
	fmt.Printf("After Test Seat allocation %s\n", s.seats)

}

func TestModifySeat(t *testing.T) {
	s := &server{
		users: make(map[string]*ticket.Receipt),
		seats: make(map[string]string),
	}

	// Add a user
	s.users["sidharth.sasikumar@example.com"] = &ticket.Receipt{
		From: "London",
		To:   "France",
		User: &ticket.User{
			FirstName: "Sidharth",
			LastName:  "Sasikumar",
			Email:     "sidharth.sasikumar@example.com",
		},
		PricePaid: 20,
		Seat:      "A1",
	}
	s.seats["A1"] = "sidharth.sasikumar@example.com"
	fmt.Printf("Current Seat allocation %s\n", s.seats)
	// Modify the seat
	req := &ticket.ModifySeatRequest{
		Email:   "sidharth.sasikumar@example.com",
		NewSeat: "B1",
	}

	res, err := s.ModifySeat(context.Background(), req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if s.users["sidharth.sasikumar@example.com"].Seat != "B1" {
		t.Fatalf("Expected seat to be updated to B1, but got %v", s.users["sidharth.sasikumar@example.com"].Seat)
	}

	if _, seatTaken := s.seats["A1"]; seatTaken {
		t.Fatalf("Expected old seat A1 to be freed, but it is still taken")
	}

	if s.seats["B1"] != "sidharth.sasikumar@example.com" {
		t.Fatalf("Expected new seat B1 to be taken by sidharth.sasikumar@example.com, but got %v", s.seats["B1"])
	}

	if res.Message != "User with email sidharth.sasikumar@example.com has been moved to seat B1" {
		t.Fatalf("Unexpected modify message: %v", res.Message)
	}
	fmt.Printf("After Test Seat allocation %s\n", s.seats)
}
