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

type createTicketsSuite struct {
	suite.Suite
	mux                *chi.Mux
	mockTicketsService *MockTicketsService
}

type MockTicketsService struct {
	mock.Mock
}

func (m *MockTicketsService) CreateTickets(ctx context.Context, name string, desc string, allocation int) (*models.Ticket, error) {
	args := m.Called(ctx, name, desc, allocation)
	return args.Get(0).(*models.Ticket), args.Error(1)
}

func (m *MockTicketsService) MapToTicketResponse(ticket *models.Ticket) server.TicketResponse {
	args := m.Called(ticket)
	return args.Get(0).(server.TicketResponse)
}

func (s *createTicketsSuite) SetupSubTest() {
	s.mockTicketsService = new(MockTicketsService)

	s.mux = chi.NewRouter()

	server.HandlerFromMux(&Handler{
		ticketsService: s.mockTicketsService,
	}, s.mux)
}

func TestCreateTickets(t *testing.T) {
	suite.Run(t, new(createTicketsSuite))
}

func (s *createTicketsSuite) Test_CreateTickets() {
	path := "/tickets"
	requestBody := "./testdata/mocks/create_tickets_request_body.json"

	s.Run("return 400 when json body cannot be decoded", func() {
		s.mockTicketsService.On("CreateTickets", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()

		res := apitest.New().Mocks().Handler(s.mux).
			Post(path).
			Body("InvalidRequestBody").
			Expect(s.T()).
			Status(http.StatusBadRequest).
			End()

		s.Assert().Equal(http.StatusBadRequest, res.Response.StatusCode)
	})

	s.Run("return 400 when ticketsService returns any error", func() {
		s.mockTicketsService.On("CreateTickets", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&models.Ticket{}, errors.New("error")).Once()

		res := apitest.New().Mocks().Handler(s.mux).
			Post(path).
			BodyFromFile(requestBody).
			Expect(s.T()).
			Status(http.StatusBadRequest).
			End()

		s.Assert().Equal(http.StatusBadRequest, res.Response.StatusCode)
	})

	s.Run("return 201 when create ticket was successful", func() {
		ticket := &models.Ticket{
			ID:          1,
			Name:        "example",
			Description: lo.ToPtr("sample description"),
			Allocation:  100,
		}

		ticketResponse := server.TicketResponse{
			Id:         ticket.ID,
			Name:       ticket.Name,
			Desc:       ticket.Description,
			Allocation: ticket.Allocation,
		}

		s.mockTicketsService.On("CreateTickets", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(ticket, nil).Once()

		s.mockTicketsService.On("MapToTicketResponse", mock.Anything).Return(ticketResponse, nil).Once()

		res := apitest.New().Mocks().Handler(s.mux).
			Post(path).
			BodyFromFile(requestBody).
			Expect(s.T()).
			Status(http.StatusCreated).
			End()

		s.Assert().Equal(http.StatusCreated, res.Response.StatusCode)
	})
}
