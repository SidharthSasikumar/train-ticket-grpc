package main

import (
	"context"
	"fmt"
	"github.com/SidharthSasikumar/train-ticket-grpc/ticket"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"sync"
)

type server struct {
	ticket.UnimplementedTicketingServiceServer
	users map[string][]*ticket.Receipt // Changed to store multiple receipts per user
	seats map[string]string
	mu    sync.Mutex
}

func (s *server) PurchaseTicket(ctx context.Context, req *ticket.PurchaseRequest) (*ticket.PurchaseResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	seat, err := s.allocateSeat()
	if err != nil {
		return nil, err
	}
	receipt := &ticket.Receipt{
		From:      req.From,
		To:        req.To,
		User:      req.User,
		PricePaid: 20,
		Seat:      seat,
	}
	s.users[req.User.Email] = append(s.users[req.User.Email], receipt)
	s.seats[seat] = req.User.Email
	fmt.Printf("Ticket purchased successfully %s\n", receipt)
	return &ticket.PurchaseResponse{
		Message: "Ticket purchased successfully",
		Receipt: receipt,
	}, nil
}
func (s *server) GetReceipt(ctx context.Context, req *ticket.GetReceiptRequest) (*ticket.GetReceiptResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt, exists := s.users[req.Email]
	if !exists {
		return nil, fmt.Errorf("no receipt found for email: %s", req.Email)
	}
	fmt.Printf("receipt:%s", receipt)
	return &ticket.GetReceiptResponse{Receipt: receipt}, nil
}

const (
	totalSeatsA = 50 // number of seats in Section A
	totalSeatsB = 50 // number of seats in Section B
)

func (s *server) ViewUsers(ctx context.Context, req *ticket.ViewUsersRequest) (*ticket.ViewUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	users := []*ticket.UserSeat{}
	sectionPrefix := req.Section

	for seat, email := range s.seats {
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

func (s *server) RemoveUser(ctx context.Context, req *ticket.RemoveUserRequest) (*ticket.RemoveUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipts, exists := s.users[req.Email]
	if !exists {
		return nil, fmt.Errorf("user with email %v Sasikumars not exist", req.Email)
	}

	// Free the seat
	for _, receipt := range receipts {
		delete(s.seats, receipt.Seat)
	}

	// Remove the user from the users map
	delete(s.users, req.Email)

	fmt.Printf("User with email %v has been removed\n", req.Email)
	return &ticket.RemoveUserResponse{
		Message: fmt.Sprintf("User with email %v has been removed", req.Email),
	}, nil
}

func (s *server) ModifySeat(ctx context.Context, req *ticket.ModifySeatRequest) (*ticket.ModifySeatResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipts, exists := s.users[req.Email]
	if !exists || len(receipts) == 0 {
		return nil, fmt.Errorf("user with email %v does not exist", req.Email)
	}

	// Check if the new seat is already taken
	if _, seatTaken := s.seats[req.NewSeat]; seatTaken {
		return nil, fmt.Errorf("seat %v is already taken", req.NewSeat)
	}

	// Modify the most recent receipt
	receipt := receipts[len(receipts)-1] // Get the most recent receipt

	// Free the old seat
	delete(s.seats, receipt.Seat)

	// Assign the new seat
	s.seats[req.NewSeat] = req.Email
	receipt.Seat = req.NewSeat

	fmt.Printf("User with email %v has been moved to seat %v\n", req.Email, req.NewSeat)
	return &ticket.ModifySeatResponse{
		Message: fmt.Sprintf("User with email %v has been moved to seat %v", req.Email, req.NewSeat),
	}, nil
}

func (s *server) allocateSeat() (string, error) {

	// Track the highest allocated seat in both sections
	for i := 1; i <= totalSeatsA; i++ {
		seat := fmt.Sprintf("A%d", i)
		if _, taken := s.seats[seat]; !taken {
			return seat, nil
		}
	}

	for i := 1; i <= totalSeatsB; i++ {
		seat := fmt.Sprintf("B%d", i)
		if _, taken := s.seats[seat]; !taken {
			return seat, nil
		}
	}

	// If no seats are available
	return "", fmt.Errorf("no seats available")
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	ticket.RegisterTicketingServiceServer(grpcServer, &server{
		users: make(map[string][]*ticket.Receipt),
		seats: make(map[string]string),
	})

	log.Printf("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
