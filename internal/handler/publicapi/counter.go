package publicapi

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/syth0le/counter-service/internal/clients/auth"
	"github.com/syth0le/counter-service/internal/model"
	"github.com/syth0le/counter-service/internal/service/counter"
)

func (h *Handler) GetUserCounter(w http.ResponseWriter, r *http.Request) {
	handleRequest := func() (*counterResponse, error) {
		ctx := r.Context()

		userIDStr := ctx.Value(auth.UserIDValue)
		if userIDStr == "" {
			return nil, fmt.Errorf("cannot recognize userID")
		}

		dialogID := chi.URLParamFromCtx(ctx, "dialogID")

		counterModel, err := h.CounterService.GetDialogCounterForUser(ctx, &counter.GetDialogCounterForUserParams{
			DialogID: model.DialogID(dialogID),
			UserID:   userIDStr.(model.UserID),
		})
		if err != nil {
			return nil, fmt.Errorf("create message: %w", err)
		}

		return counterModelToResponse(counterModel), nil
	}

	resp, err := handleRequest()
	if err != nil {
		h.writeError(r.Context(), w, fmt.Errorf("get user counter: %w", err))
		return
	}

	writeResponse(w, resp)
}

func (h *Handler) GetUserCountersList(w http.ResponseWriter, r *http.Request) {
	handleRequest := func() (*counterListResponse, error) {
		ctx := r.Context()

		userIDStr := ctx.Value(auth.UserIDValue)
		if userIDStr == "" {
			return nil, fmt.Errorf("cannot recognize userID")
		}

		counters, err := h.CounterService.GetUserCounters(ctx, &counter.GetUserCountersParams{
			UserID: userIDStr.(model.UserID),
		})
		if err != nil {
			return nil, fmt.Errorf("create message: %w", err)
		}

		return counterModelsToResponse(counters), nil
	}

	resp, err := handleRequest()
	if err != nil {
		h.writeError(r.Context(), w, fmt.Errorf("get user counters list: %w", err))
		return
	}

	writeResponse(w, resp)
}

type counterResponse struct {
	DialogID string `json:"dialog_id"`
	Counter  int64  `json:"counter"`
}

type counterListResponse struct {
	Counters []*counterResponse `json:"counters"`
}

func counterModelToResponse(counter *model.Counter) *counterResponse {
	return &counterResponse{
		DialogID: counter.DialogID.String(),
		Counter:  counter.Value,
	}
}

func counterModelsToResponse(counterModels []*model.Counter) *counterListResponse {
	counters := make([]*counterResponse, 0)
	for _, counterModel := range counterModels {
		counters = append(counters, counterModelToResponse(counterModel))
	}

	return &counterListResponse{
		Counters: counters,
	}
}
