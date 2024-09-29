package repository

import (
	"context"

	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/models"
)

type Repository interface {
	CreateTickets(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error)
	GetTicketByID(ctx context.Context, id int) (*models.Ticket, error)
	PurchaseTickets(ctx context.Context, quantity int, ticketID int) error
}
