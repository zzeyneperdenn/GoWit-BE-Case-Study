package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/zzeyneperdenn/GoWit-BE-Case-Study/internal/api"
)

func (h *Handler) GetTicketById(w http.ResponseWriter, r *http.Request, id int) {
	ticket, err := h.ticketsService.GetTicketByID(r.Context(), id)
	if err != nil {
		api.RenderHTTPError(http.StatusNotFound, api.NotFoundError(), r, w)
		log.Printf("cannot find the ticket with id: %v\n", id)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, h.ticketsService.MapToTicketResponse(ticket))
}
