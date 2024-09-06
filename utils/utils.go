package utils

import (
	"fmt"
)

const (
	totalSeatsA = 50 // number of seats in Section A
	totalSeatsB = 50 // number of seats in Section B
)

// AllocateSeat allocates a seat for the user based on availability
func AllocateSeat(seats map[string]string) (string, error) {
	// Track the highest allocated seat in both sections
	for i := 1; i <= totalSeatsA; i++ {
		seat := fmt.Sprintf("A%d", i)
		if _, taken := seats[seat]; !taken {
			return seat, nil
		}
	}

	for i := 1; i <= totalSeatsB; i++ {
		seat := fmt.Sprintf("B%d", i)
		if _, taken := seats[seat]; !taken {
			return seat, nil
		}
	}

	// If no seats are available
	return "", fmt.Errorf("no seats available")
}
