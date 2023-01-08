package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/hereisajvi/noty/internal/delivery/http/status"
	"github.com/hereisajvi/noty/internal/model"
	"github.com/hereisajvi/noty/internal/service"
)

type (
	CreateMessageRequest struct {
		Text string `json:"text"`
	}

	UpdateMessageRequest struct {
		Text string `json:"text"`
	}
)

type MessageHandler struct {
	service MessageService
}

func NewMessageHandler(service MessageService) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

func (h *MessageHandler) Routes(router chi.Router) {
	router.Route("/v1", func(r chi.Router) {
		r.Post("/messages", h.Create)
		r.Put("/messages/{id}", h.Update)
		r.Delete("/messages/{id}", h.Delete)
		r.Get("/messages", h.List)
	})
}

// Create 		godoc
// @Summary 	Creates a new message
// @Description Creates a new message
// @Tags 		messages
// @Accept      json
// @Produce     json
// @Success 	201 	{object} model.MessageDataTransferObject
// @Failure 	400,500 {object} status.ErrorResponse
// @Router 		/v1/messages [post]
func (h *MessageHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var request CreateMessageRequest

	err := decoder.Decode(&request)
	if err != nil {
		status.Error(w, status.BadRequestError, err.Error())
		return
	}

	message, err := h.service.Create(r.Context(), model.MessageBusinessObject{
		Text: request.Text,
	})
	if err != nil {
		status.Error(w, status.InternalServerError, err.Error())
		return
	}

	status.OK(w, &message, http.StatusCreated)
}

// Update 		godoc
// @Summary 	Updates an existing message by its id
// @Description Updates an existing message by its id
// @Tags 		messages
// @Accept      json
// @Produce     json
// @Param       id   		path     string  true  "Message ID"
// @Success 	200 		{object} model.MessageDataTransferObject
// @Failure 	400,404,500 {object} status.ErrorResponse
// @Router 		/v1/messages/{id} [put]
func (h *MessageHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	decoder := json.NewDecoder(r.Body)

	var request UpdateMessageRequest

	err := decoder.Decode(&request)
	if err != nil {
		status.Error(w, status.BadRequestError, err.Error())
		return
	}

	message, err := h.service.Update(r.Context(), id, model.MessageBusinessObject{
		Text: request.Text,
	})
	if err != nil {
		if errors.Is(err, service.ErrNoMessage) {
			status.Error(w, status.NotFoundError)
			return
		}

		status.Error(w, status.InternalServerError, err.Error())

		return
	}

	status.OK(w, &message, http.StatusOK)
}

// Delete 		godoc
// @Summary 	Deletes an existing message by its id
// @Description Deletes an existing message by its id
// @Tags 		messages
// @Param       id   		 path    string  true  "Message ID"
// @Success 	204
// @Failure 	400,404,500 {object} status.ErrorResponse
// @Router 		/v1/messages/{id} [delete]
func (h *MessageHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNoMessage) {
			status.Error(w, status.NotFoundError)
			return
		}

		status.Error(w, status.InternalServerError, err.Error())

		return
	}

	status.OK(w, http.NoBody, http.StatusNoContent)
}

// List 		godoc
// @Summary 	Returns a list of all messages
// @Description Returns a list of all messages
// @Tags 		messages
// @Produce     json
// @Success 	200 {array}  model.MessageDataTransferObject
// @Failure 	500 {object} status.ErrorResponse
// @Router 		/v1/messages [get]
func (h *MessageHandler) List(w http.ResponseWriter, r *http.Request) {
	messages, err := h.service.List(r.Context())
	if err != nil {
		if errors.Is(err, service.ErrNoMessage) {
			status.Empty(w, http.StatusOK)
			return
		}

		status.Error(w, status.InternalServerError, err.Error())

		return
	}

	status.OK(w, &messages, http.StatusOK)
}
