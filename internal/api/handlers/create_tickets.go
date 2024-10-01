package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/samber/lo"

	"github.com/go-chi/render"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api/server"
)

func (h *Handler) CreateTickets(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var requestBody server.CreateTicketsRequestBody

	if err := decoder.Decode(&requestBody); err != nil {
		api.RenderHTTPError(http.StatusBadRequest, api.BadRequestError(), r, w)
		log.Printf("invalid request body %+v\n", requestBody)
		return
	}

	if requestBody.Allocation <= 0 || requestBody.Name == "" {
		api.RenderHTTPError(http.StatusBadRequest, api.BadRequestError(), r, w)
		log.Printf("request is not valid %+v\n", requestBody)
		return
	}

	ticket, err := h.ticketsService.CreateTickets(r.Context(), requestBody.Name, lo.FromPtr(requestBody.Desc), requestBody.Allocation)
	if err != nil {
		api.RenderHTTPError(http.StatusBadRequest, api.NotFoundError(), r, w)
		log.Printf("could not create ticket. %s\n", err.Error())
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, h.ticketsService.MapToTicketResponse(ticket))
}
