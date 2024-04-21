package oteldemo

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

var (
	ErrDuplicatedLead = errors.New("email already in use")
	ErrLeadNotFound   = errors.New("lead not found")
)

type LeadRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type Lead struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Group       string    `json:"group"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

type LeadService interface {
	Create(ctx context.Context, newLead Lead) error
	GetByID(ctx context.Context, id string) (Lead, error)
}

func CreateLeadByParam(request LeadRequest) Lead {
	now := time.Now().UTC()
	return Lead{
		Name:        request.Name,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		ID:          uuid.NewString(),
		Group:       getRandomGroup(),
		CreatedAt:   now,
		ModifiedAt:  now,
	}
}

func getRandomGroup() string {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(4)

	for index, value := range []string{"junior", "middle", "master", "genius", "slot"} {
		if index == number {
			return value
		}
	}
	return "random"
}
