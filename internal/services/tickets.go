package services

import (
	"context"

	"github.com/samber/lo"

	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api/server"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/models"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/repository"
)

type TicketsService struct {
	db repository.Repository
}

type ITicketsService interface {
	CreateTickets(ctx context.Context, name string, description string, allocation int) (*models.Ticket, error)
	GetTicketByID(ctx context.Context, id int) (*models.Ticket, error)
	PurchaseTickets(ctx context.Context, quantity int, ticketID int) error
	MapToTicketResponse(ticket *models.Ticket) server.TicketResponse
}

func NewTicketsService(db repository.Repository) *TicketsService {
	return &TicketsService{db}
}

func (s *TicketsService) CreateTickets(ctx context.Context, name string, description string, allocation int) (*models.Ticket, error) {
	createTicket := &models.Ticket{
		Name:        name,
		Description: lo.ToPtr(description),
		Allocation:  allocation,
	}
	ticket, err := s.db.CreateTickets(ctx, createTicket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *TicketsService) GetTicketByID(ctx context.Context, id int) (*models.Ticket, error) {
	ticket, err := s.db.GetTicketByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *TicketsService) PurchaseTickets(ctx context.Context, quantity int, ticketID int) error {
	err := s.db.PurchaseTickets(ctx, quantity, ticketID)
	if err != nil {
		return err
	}

	return nil
}

func (s *TicketsService) MapToTicketResponse(ticket *models.Ticket) server.TicketResponse {
	return server.TicketResponse{
		Allocation: ticket.Allocation,
		Desc:       ticket.Description,
		Id:         ticket.ID,
		Name:       ticket.Name,
	}
}
