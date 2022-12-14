// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package repository

import (
	"context"
)

type Querier interface {
	CreateTicket(ctx context.Context, arg CreateTicketParams) (Ticket, error)
	DeleteTicket(ctx context.Context, arg DeleteTicketParams) error
	GetTicket(ctx context.Context, arg GetTicketParams) (Ticket, error)
	ListTickets(ctx context.Context, username string) ([]Ticket, error)
}

var _ Querier = (*Queries)(nil)
