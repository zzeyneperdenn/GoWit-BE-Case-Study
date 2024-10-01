package handlers

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/samber/lo"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api/server"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/models"
)

type GetTicketByIdSuite struct {
	suite.Suite
	mux                *chi.Mux
	mockTicketsService *MockTicketsService
	handler            *Handler
}

func (m *MockTicketsService) GetTicketByID(ctx context.Context, id int) (*models.Ticket, error) {
	args := m.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Ticket), args.Error(1)
	}
	return nil, args.Error(1)
}

func (s *GetTicketByIdSuite) SetupTest() {
	s.mockTicketsService = new(MockTicketsService)

	s.mux = chi.NewRouter()

	server.HandlerFromMux(&Handler{
		ticketsService: s.mockTicketsService,
	}, s.mux)
}

func TestGetTicketById(t *testing.T) {
	suite.Run(t, new(GetTicketByIdSuite))
}

func (s *GetTicketByIdSuite) Test_GetTicketById_NotFound() {
	s.Run("return 400 when json body cannot be decoded", func() {
		path := "/tickets/1"

		s.mockTicketsService.On("GetTicketByID", mock.Anything, 1).Return(nil, errors.New("ticket not found")).Once()

		res := apitest.New().Mocks().Handler(s.mux).
			Get(path).
			Expect(s.T()).
			Status(http.StatusNotFound).
			End()

		s.Assert().Equal(http.StatusNotFound, res.Response.StatusCode)
	})

	s.Run("return 201 when create ticket was successful", func() {
		path := "/tickets/1"

		ticket := &models.Ticket{
			ID:          1,
			Name:        "Sample Ticket",
			Description: lo.ToPtr("This is a sample ticket"),
			Allocation:  100,
		}

		ticketResponse := server.TicketResponse{
			Id:         ticket.ID,
			Name:       ticket.Name,
			Desc:       ticket.Description,
			Allocation: ticket.Allocation,
		}

		s.mockTicketsService.On("GetTicketByID", mock.Anything, 1).Return(ticket, nil).Once()

		s.mockTicketsService.On("MapToTicketResponse", ticket).Return(ticketResponse).Once()

		res := apitest.New().Mocks().Handler(s.mux).
			Get(path).
			Expect(s.T()).
			Status(http.StatusOK).
			End()

		s.Assert().Equal(http.StatusOK, res.Response.StatusCode)
	})
}
