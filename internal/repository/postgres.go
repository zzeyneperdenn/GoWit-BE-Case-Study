package repository

import (
	"context"
	"fmt"

	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type postgresRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &postgresRepository{db: db}
}

func (p *postgresRepository) CreateTickets(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	result := p.db.WithContext(ctx).Create(ticket)
	if result.Error != nil {
		return nil, result.Error
	}

	return ticket, nil
}

func (p *postgresRepository) GetTicketByID(ctx context.Context, id int) (*models.Ticket, error) {
	var ticket models.Ticket

	result := p.db.WithContext(ctx).
		Where("id = ?", id).
		First(&ticket)
	if result.Error != nil {
		return nil, result.Error
	}

	return &ticket, nil
}

func (p *postgresRepository) PurchaseTickets(ctx context.Context, quantity int, ticketID int) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		var ticket models.Ticket

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", ticketID).
			First(&ticket).Error; err != nil {
			return fmt.Errorf("failed to lock the tickets row: %v", err)
		}

		if ticket.Allocation < quantity {
			return fmt.Errorf("not enough tickets available")
		}

		ticket.Allocation -= quantity
		if err := tx.Save(&ticket).Error; err != nil {
			return fmt.Errorf("failed to update tickets: %v", err)
		}

		fmt.Printf("Successfully purchased %d tickets for event %d\n", quantity, ticketID)
		return nil
	})
}
