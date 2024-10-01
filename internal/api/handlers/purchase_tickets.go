package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api/server"
)

func (h *Handler) PurchaseTickets(w http.ResponseWriter, r *http.Request, id int) {
	decoder := json.NewDecoder(r.Body)

	var requestBody server.PurchaseTicketsRequestBody

	if err := decoder.Decode(&requestBody); err != nil {
		api.RenderHTTPError(http.StatusBadRequest, api.BadRequestError(), r, w)
		log.Printf("invalid request body %+v\n", requestBody)
		return
	}

	if requestBody.Quantity <= 0 || requestBody.UserId == "" {
		api.RenderHTTPError(http.StatusBadRequest, api.BadRequestError(), r, w)
		log.Printf("request is not valid %+v\n", requestBody)
		return
	}

	err := h.ticketsService.PurchaseTickets(r.Context(), requestBody.Quantity, id)
	if err != nil {
		api.RenderHTTPError(http.StatusBadRequest, api.BadRequestError(), r, w)
		log.Printf("could not purchase any tickets. %s\n", err.Error())
		return
	}

	render.Status(r, http.StatusOK)
}
