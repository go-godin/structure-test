package ticket

import (
	"context"

	"github.com/go-godin/log"
)

type Service interface {
	Create(ctx context.Context, title string) (Ticket, error)
}

type service struct {
	logger log.Logger
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (svc *service) Create(ctx context.Context, title string) (Ticket, error) {
	return Ticket{Title:title}, nil
}
