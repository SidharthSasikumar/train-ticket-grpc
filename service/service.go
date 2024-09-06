package service

import (
	"context"
	"fmt"
	"github.com/SidharthSasikumar/train-ticket-grpc/ticket"
	"github.com/SidharthSasikumar/train-ticket-grpc/utils"
	"strings"
	"sync"
)

type Server struct {
	ticket.UnimplementedTicketingServiceServer
	Users map[string][]*ticket.Receipt // Changed to store multiple receipts per user
	Seats map[string]string
	mu    sync.Mutex
}

func (s *Server) PurchaseTicket(ctx context.Context, req *ticket.PurchaseRequest) (*ticket.PurchaseResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	seat, err := utils.AllocateSeat(s.Seats)
	if err != nil {
		return nil, err
	}

	if receipts, userExist := s.Users[req.User.Email]; userExist {
		for _, receipt := range receipts { // Correctly iterating over the user's receipts
			if receipt.User.FirstName == req.User.FirstName && receipt.User.LastName == req.User.LastName {
				return nil, fmt.Errorf("the user already has a seat for the journey: seatNumber: %v", receipt.Seat)
			}
		}
	}

	receipt := &ticket.Receipt{
		From:      req.From,
		To:        req.To,
		User:      req.User,
		PricePaid: 20,
		Seat:      seat,
	}
	s.Users[req.User.Email] = append(s.Users[req.User.Email], receipt)
	s.Seats[seat] = req.User.Email
	fmt.Printf("Ticket purchased successfully %s\n", receipt)
	return &ticket.PurchaseResponse{
		Message: "Ticket purchased successfully",
		Receipt: receipt,
	}, nil
}
func (s *Server) GetReceipt(ctx context.Context, req *ticket.GetReceiptRequest) (*ticket.GetReceiptResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt, exists := s.Users[req.Email]
	if !exists {
		return nil, fmt.Errorf("no receipt found for email: %s", req.Email)
	}
	fmt.Printf("receipt:%s", receipt)
	return &ticket.GetReceiptResponse{Receipt: receipt}, nil
}
func (s *Server) ViewUsers(ctx context.Context, req *ticket.ViewUsersRequest) (*ticket.ViewUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	users := []*ticket.UserSeat{}
	sectionPrefix := req.Section

	for seat, email := range s.Seats {
		if strings.HasPrefix(seat, sectionPrefix) {
			users = append(users, &ticket.UserSeat{
				Email: email,
				Seat:  seat,
			})
		}
	}
	fmt.Printf("User Deitals are %s\n", &ticket.ViewUsersResponse{Users: users})
	return &ticket.ViewUsersResponse{Users: users}, nil
}

func (s *Server) RemoveUser(ctx context.Context, req *ticket.RemoveUserRequest) (*ticket.RemoveUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipts, exists := s.Users[req.Email]
	if !exists {
		return nil, fmt.Errorf("user with email %v Sasikumars not exist", req.Email)
	}

	// Free the seat
	for _, receipt := range receipts {
		delete(s.Seats, receipt.Seat)
	}

	// Remove the user from the users map
	delete(s.Users, req.Email)

	fmt.Printf("User with email %v has been removed\n", req.Email)
	return &ticket.RemoveUserResponse{
		Message: fmt.Sprintf("User with email %v has been removed", req.Email),
	}, nil
}

func (s *Server) ModifySeat(ctx context.Context, req *ticket.ModifySeatRequest) (*ticket.ModifySeatResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipts, exists := s.Users[req.Email]
	if !exists || len(receipts) == 0 {
		return nil, fmt.Errorf("user with email %v does not exist", req.Email)
	}

	// Check if the new seat is already taken
	if _, seatTaken := s.Seats[req.NewSeat]; seatTaken {
		return nil, fmt.Errorf("seat %v is already taken", req.NewSeat)
	}

	// Modify the most recent receipt
	receipt := receipts[len(receipts)-1] // Get the most recent receipt

	// Free the old seat
	delete(s.Seats, receipt.Seat)

	// Assign the new seat
	s.Seats[req.NewSeat] = req.Email
	receipt.Seat = req.NewSeat

	fmt.Printf("User with email %v has been moved to seat %v\n", req.Email, req.NewSeat)
	return &ticket.ModifySeatResponse{
		Message: fmt.Sprintf("User with email %v has been moved to seat %v", req.Email, req.NewSeat),
	}, nil
}
