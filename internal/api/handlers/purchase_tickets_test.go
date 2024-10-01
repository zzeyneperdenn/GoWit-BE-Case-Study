package handlers

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api/server"
)

type purchaseTicketsSuite struct {
	suite.Suite
	mux                *chi.Mux
	mockTicketsService *MockTicketsService
}

func (m *MockTicketsService) PurchaseTickets(ctx context.Context, quantity int, ticketID int) error {
	args := m.Called(ctx, quantity, ticketID)
	return args.Error(0)
}

func (s *purchaseTicketsSuite) SetupSubTest() {
	s.mockTicketsService = new(MockTicketsService)

	s.mux = chi.NewRouter()

	server.HandlerFromMux(&Handler{
		ticketsService: s.mockTicketsService,
	}, s.mux)
}

func TestPurchaseTickets(t *testing.T) {
	suite.Run(t, new(purchaseTicketsSuite))
}

func (s *purchaseTicketsSuite) Test_PurchaseTickets() {
	path := "/tickets/1/purchases"
	requestBody := "./testdata/mocks/purchase_tickets_request_body.json"

	s.Run("return 400 when json body cannot be decoded", func() {
		s.mockTicketsService.On("PurchaseTickets", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		res := apitest.New().Mocks().Handler(s.mux).
			Post(path).
			Body("InvalidRequestBody").
			Expect(s.T()).
			Status(http.StatusBadRequest).
			End()

		s.Assert().Equal(http.StatusBadRequest, res.Response.StatusCode)
	})

	s.Run("return 400 when PurchaseTickets returns an error", func() {
		s.mockTicketsService.On("PurchaseTickets", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		res := apitest.New().Mocks().Handler(s.mux).
			Post(path).
			BodyFromFile(requestBody).
			Expect(s.T()).
			Status(http.StatusBadRequest).
			End()

		s.Assert().Equal(http.StatusBadRequest, res.Response.StatusCode)

		s.mockTicketsService.AssertExpectations(s.T())
	})

	s.Run("return 200 when purchase tickets was successful", func() {
		s.mockTicketsService.On("PurchaseTickets", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		res := apitest.New().Mocks().Handler(s.mux).
			Post(path).
			BodyFromFile(requestBody).
			Expect(s.T()).
			Status(http.StatusOK).
			End()

		s.Assert().Equal(http.StatusOK, res.Response.StatusCode)
	})
}
