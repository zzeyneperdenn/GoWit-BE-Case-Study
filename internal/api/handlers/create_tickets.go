package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api/server"
)

func (h *Handler) CreateTickets(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var requestBody server.CreateTicketsRequestBody

	if err := decoder.Decode(&requestBody); err != nil {
		api.RenderHTTPError(http.StatusBadRequest, api.BadRequestError(), r, w)
		return
	}

	ticket, err := h.ticketsService.CreateTickets(r.Context(), requestBody.Name, requestBody.Desc, requestBody.Allocation)
	if err != nil {
		api.RenderHTTPError(http.StatusNotFound, api.NotFoundError(), r, w)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, h.ticketsService.MapToTicketResponse(ticket))
}
