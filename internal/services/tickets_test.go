package services

import (
	"context"
	"errors"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/models"
)

type TicketsServiceTestSuite struct {
	suite.Suite
	mockRepo      *MockRepository
	ticketService *TicketsService
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateTickets(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	args := m.Called(ctx, ticket)
	return args.Get(0).(*models.Ticket), args.Error(1)
}

func (m *MockRepository) GetTicketByID(ctx context.Context, id int) (*models.Ticket, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Ticket), args.Error(1)
}

func (m *MockRepository) PurchaseTickets(ctx context.Context, quantity int, ticketID int) error {
	args := m.Called(ctx, quantity, ticketID)
	return args.Error(0)
}

func TestTicketsServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TicketsServiceTestSuite))
}

func (s *TicketsServiceTestSuite) SetupTest() {
	s.mockRepo = new(MockRepository)
	s.ticketService = NewTicketsService(s.mockRepo)
}

func (s *TicketsServiceTestSuite) Test_CreateTickets() {
	s.Run("return error when cannot create ticket", func() {
		s.mockRepo.On("CreateTickets", mock.Anything, mock.AnythingOfType("*models.Ticket")).Return(&models.Ticket{}, errors.New("database error")).Once()

		result, err := s.ticketService.CreateTickets(context.TODO(), "Test Ticket", "Test Description", 100)

		assert.Nil(s.T(), result)
		assert.Error(s.T(), err)
		s.mockRepo.AssertExpectations(s.T())
	})

	s.Run("return ticket when success", func() {
		ticket := &models.Ticket{
			ID:          1,
			Name:        "Test Ticket",
			Description: lo.ToPtr("Test Desc"),
			Allocation:  100,
		}

		s.mockRepo.On("CreateTickets", mock.Anything, mock.AnythingOfType("*models.Ticket")).Return(ticket, nil).Once()

		result, err := s.ticketService.CreateTickets(context.TODO(), ticket.Name, *ticket.Description, ticket.Allocation)

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), ticket, result)
		s.mockRepo.AssertExpectations(s.T())
	})
}

func (s *TicketsServiceTestSuite) Test_GetTicketByID() {
	s.Run("return error when cannot find ticket", func() {
		s.mockRepo.On("GetTicketByID", mock.Anything, mock.AnythingOfType("int")).Return(&models.Ticket{}, errors.New("not found")).Once()

		result, err := s.ticketService.GetTicketByID(context.TODO(), 1)

		assert.Nil(s.T(), result)
		assert.Error(s.T(), err)
		s.mockRepo.AssertExpectations(s.T())
	})

	s.Run("return ticket when success", func() {
		ticket := &models.Ticket{
			ID:          1,
			Name:        "Test Ticket",
			Description: lo.ToPtr("desc"),
			Allocation:  100,
		}

		s.mockRepo.On("GetTicketByID", mock.Anything, ticket.ID).Return(ticket, nil).Once()

		result, err := s.ticketService.GetTicketByID(context.TODO(), ticket.ID)

		assert.NoError(s.T(), err)
		assert.Equal(s.T(), ticket, result)
		s.mockRepo.AssertExpectations(s.T())
	})
}

func (s *TicketsServiceTestSuite) Test_PurchaseTickets() {
	s.Run("return error when cannot purchase", func() {
		s.mockRepo.On("PurchaseTickets", mock.Anything, 10, 1).Return(errors.New("not enough tickets")).Once()

		err := s.ticketService.PurchaseTickets(context.TODO(), 10, 1)

		assert.Error(s.T(), err)
		s.mockRepo.AssertExpectations(s.T())
	})

	s.Run("return nil when success", func() {
		s.mockRepo.On("PurchaseTickets", mock.Anything, 10, 1).Return(nil).Once()

		err := s.ticketService.PurchaseTickets(context.TODO(), 10, 1)

		assert.NoError(s.T(), err)
		s.mockRepo.AssertExpectations(s.T())
	})
}
