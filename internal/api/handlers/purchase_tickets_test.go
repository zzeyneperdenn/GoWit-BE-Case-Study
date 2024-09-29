package handlers

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/repository"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/services"
)

type purchaseTicketsSuite struct {
	suite.Suite
	mux *chi.Mux
	// mockTicketsService *mocks.ITicketsService
}

type MockTicketsService struct {
	mock.Mock
	services.TicketsService
}

type MockDB struct {
	mock.Mock
	repository.Repository
}

func initHandlers(ticketsService *services.TicketsService) struct {
	*Handler
} {
	return struct{ *Handler }{
		Handler: &Handler{
			ticketsService: ticketsService,
		},
	}
}

func (s *purchaseTicketsSuite) SetupSubTest() {
	s.mux = chi.NewRouter()

	//	mockTicketsService := &MockTicketsService{}

	//server.HandlerFromMux(initHandlers(&mockTicketsService), s.mux)
}

func TestPurchaseTickets(t *testing.T) {
	suite.Run(t, new(purchaseTicketsSuite))
}

// func (s *purchaseTicketsSuite) Test_V1ActivateFlexiCredit() {
// 	path := "/v1/flexi-credit/accounts/activate"

// 	requestBody := "./testdata/mocks/activate-flexi-credit/activate_flexi_credit_request_body.json"

// 	userID := "12345"
// 	accountID := "4cbc6747-9a43-4de3-81ad-8f87c23a8f8b"
// 	parsedID, _ := uuid.Parse("4cbc6747-9a43-4de3-81ad-8f87c23a8f8b")
// 	var gtResp *gotron.V1GetPaymentSourcesResponse

// 	s.Run("return 400 when json body cannot be decoded", func() {
// 		res := apitest.New().Mocks().Handler(s.mux).
// 			Post(path).
// 			Body("InvalidRequestBody").
// 			Header("Content-Type", httpUtils.ApplicationJSONType).
// 			Header("X-User-Id", userID).
// 			Expect(s.T()).
// 			Status(http.StatusBadRequest).
// 			End()

// 		s.Assert().Equal(http.StatusBadRequest, res.Response.StatusCode)

// 	})

// 	s.Run("return 404 when cannot get cl account by userID", func() {
// 		mockFlexiCreditAccountService.On("GetFlexiCreditAccountByUserID", mock.Anything, mock.Anything).Return(&models.FlexiCreditAccount{}, errors.New("error")).Once()

// 		res := apitest.New().Mocks().Handler(s.mux).
// 			Post(path).
// 			BodyFromFile(requestBody).
// 			Header("Content-Type", httpUtils.ApplicationJSONType).
// 			Header("X-User-Id", userID).
// 			Header("X-Tenant", "sample_tenant").
// 			QueryParams(map[string]string{
// 				"filters[user_id]": userID,
// 			}).
// 			Expect(s.T()).
// 			Status(http.StatusNotFound).
// 			End()

// 		var response server.ErrorResponse
// 		res.JSON(&response)

// 		s.Assert().Equal(http.StatusNotFound, res.Response.StatusCode)

// 		s.Assert().Equal("FlexiCredit account not found for user: 12345", response.Error())

// 	})

// 	s.Run("return 400 when cl account has accountID", func() {
// 		mockFlexiCreditAccountService.On("GetFlexiCreditAccountByUserID", mock.Anything, mock.Anything).Return(&models.FlexiCreditAccount{
// 			UserID:    userID,
// 			AccountID: &accountID,
// 		}, nil).Once()

// 		res := apitest.New().Mocks().Handler(s.mux).
// 			Post(path).
// 			BodyFromFile(requestBody).
// 			Header("Content-Type", httpUtils.ApplicationJSONType).
// 			Header("X-User-Id", userID).
// 			Header("X-Tenant", "sample_tenant").
// 			QueryParams(map[string]string{
// 				"filters[user_id]": userID,
// 			}).
// 			Expect(s.T()).
// 			Status(http.StatusBadRequest).
// 			End()

// 		var response server.ErrorResponse
// 		res.JSON(&response)

// 		s.Assert().Equal(http.StatusBadRequest, res.Response.StatusCode)

// 		s.Assert().Equal("FlexiCredit Account already activated", response.Error())

// 	})

// 	s.Run("return 500 when gotron returns error", func() {
// 		mockFlexiCreditAccountService.On("GetFlexiCreditAccountByUserID", mock.Anything, userID).Return(&models.FlexiCreditAccount{
// 			ID:        &parsedID,
// 			UserID:    userID,
// 			AccountID: nil,
// 		}, nil).Once()

// 		mockGotronService.On("V1GetPaymentSources", mock.Anything, &gotron.V1GetPaymentSourcesParams{
// 			Filters: &gotron.FetchPaymentSourcesFilters{
// 				UserId: &userID,
// 			},
// 		}).Return(gtResp, errors.New("error")).Once()

// 		res := apitest.New().Mocks().Handler(s.mux).
// 			Post(path).
// 			BodyFromFile(requestBody).
// 			Header("Content-Type", httpUtils.ApplicationJSONType).
// 			Header("X-User-Id", userID).
// 			Header("X-Tenant", "sample_tenant").
// 			QueryParams(map[string]string{
// 				"filters[user_id]": userID,
// 			}).
// 			Expect(s.T()).
// 			Status(http.StatusInternalServerError).
// 			End()

// 		var response server.ErrorResponse
// 		res.JSON(&response)

// 		s.Assert().Equal(http.StatusInternalServerError, res.Response.StatusCode)

// 		s.Assert().Equal("cannot get payment sources", response.Error())

// 	})

// 	s.Run("return 404 when there is no cl offer", func() {
// 		mockFlexiCreditAccountService.On("GetFlexiCreditAccountByUserID", mock.Anything, userID).Return(&models.FlexiCreditAccount{
// 			ID:        &parsedID,
// 			UserID:    userID,
// 			AccountID: nil,
// 		}, nil).Once()

// 		paymentSourcesWithoutCard := "./testdata/mocks/activate-flexi-credit/payment_sources_with_card.json"

// 		mockGotronService.On("V1GetPaymentSources", mock.Anything, &gotron.V1GetPaymentSourcesParams{
// 			Filters: &gotron.FetchPaymentSourcesFilters{
// 				UserId: &userID,
// 			},
// 		}).Return(decodeFileToPaymentSources(paymentSourcesWithoutCard), nil).Once()

// 		mockFlexiCreditAccountService.On("GetCurrentFlexiCreditOfferByAccountID", mock.Anything, parsedID.String()).Return(&models.FlexiCreditOffer{}, errors.New("error")).Once()

// 		res := apitest.New().Mocks().Handler(s.mux).
// 			Post(path).
// 			BodyFromFile(requestBody).
// 			Header("Content-Type", httpUtils.ApplicationJSONType).
// 			Header("X-User-Id", "12345").
// 			Header("X-Tenant", "sample_tenant").
// 			QueryParams(map[string]string{
// 				"filters[user_id]": "12345",
// 			}).
// 			Expect(s.T()).
// 			Status(http.StatusNotFound).
// 			End()

// 		var response server.ErrorResponse
// 		res.JSON(&response)

// 		s.Assert().Equal(http.StatusNotFound, res.Response.StatusCode)

// 		s.Assert().Equal("cannot get current FlexiCredit offer with AccountID:4cbc6747-9a43-4de3-81ad-8f87c23a8f8b", response.Error())
// 	})

// 	s.Run("return 202 with SUCCESS status when card is not connected", func() {
// 		mockFlexiCreditAccountService.On("GetFlexiCreditAccountByUserID", mock.Anything, userID).Return(&models.FlexiCreditAccount{
// 			ID:        &parsedID,
// 			UserID:    userID,
// 			AccountID: nil,
// 		}, nil).Once()

// 		mockFlexiCreditAccountService.On("GetPlatformAccount", mock.Anything, mock.Anything, userID).Return(&server.PlatformAccount{
// 			Id:     parsedID,
// 			UserId: userID,
// 		}, nil).Once()

// 		paymentSourcesWithoutCard := "./testdata/mocks/activate-flexi-credit/payment_sources_without_card.json"

// 		mockGotronService.On("V1GetPaymentSources", mock.Anything, &gotron.V1GetPaymentSourcesParams{
// 			Filters: &gotron.FetchPaymentSourcesFilters{
// 				UserId: &userID,
// 			},
// 		}).Return(decodeFileToPaymentSources(paymentSourcesWithoutCard), nil).Once()

// 		res := apitest.New().Mocks().Handler(s.mux).
// 			Post(path).
// 			BodyFromFile(requestBody).
// 			Header("Content-Type", httpUtils.ApplicationJSONType).
// 			Header("X-User-Id", "12345").
// 			Header("X-Tenant", "sample_tenant").
// 			QueryParams(map[string]string{
// 				"filters[user_id]": "12345",
// 			}).
// 			Expect(s.T()).
// 			Status(http.StatusAccepted).
// 			Assert(jsonpath.Equal(`$.data.status`, string(server.ActivationStatusAWAITINGCARDCONNECTION))).
// 			End()

// 		var response server.ErrorResponse
// 		res.JSON(&response)

// 		s.Assert().Equal(http.StatusAccepted, res.Response.StatusCode)

// 	})
// }
