package handlers

import (
	"encoding/json"
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
		return
	}

	err := h.ticketsService.PurchaseTickets(r.Context(), *requestBody.Quantity, id)
	if err != nil {
		api.RenderHTTPError(http.StatusBadRequest, api.BadRequestError(), r, w)
		return
	}

	render.Status(r, http.StatusOK)
}
