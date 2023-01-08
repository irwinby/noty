package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/go-chi/chi/v5"
	testifymock "github.com/stretchr/testify/mock"

	"github.com/hereisajvi/noty/internal/mock"
	"github.com/hereisajvi/noty/internal/model"
)

func TestMessageHandler_Create(t *testing.T) {
	type request struct {
		body CreateMessageRequest
	}

	type response struct {
		body   model.MessageDataTransferObject
		status int
		err    error
	}

	type behavior struct {
		fn func(service *mock.MessageService, request request, response response)
	}

	tests := map[string]struct {
		request  request
		behavior behavior
		response response
	}{
		"OK": {
			request: request{
				body: CreateMessageRequest{
					Text: "<text>",
				},
			},
			behavior: behavior{
				fn: func(service *mock.MessageService, request request, response response) {
					service.
						EXPECT().
						Create(
							testifymock.Anything,
							model.MessageBusinessObject{
								Text: request.body.Text,
							},
						).
						Return(
							response.body,
							response.err,
						)
				},
			},
			response: response{
				body: model.MessageDataTransferObject{
					Text: "<text>",
				},
				status: http.StatusCreated,
			},
		},
	}

	service := mock.NewMessageService(t)

	router := chi.NewRouter()

	handler := NewMessageHandler(service)
	handler.Routes(router)

	server := httptest.NewServer(router)
	defer server.Close()

	expect := httpexpect.New(t, server.URL)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.behavior.fn(service, test.request, test.response)

			expect.
				POST("/v1/messages").
				WithJSON(test.request.body).
				Expect().
				Status(test.response.status).
				JSON().
				Equal(test.response.body)
		})
	}
}

func TestMessageHandler_Update(t *testing.T) {
	type request struct {
		id   string
		body UpdateMessageRequest
	}

	type response struct {
		body   model.MessageDataTransferObject
		status int
		err    error
	}

	type behavior struct {
		fn func(service *mock.MessageService, request request, response response)
	}

	tests := map[string]struct {
		request  request
		behavior behavior
		response response
	}{
		"OK": {
			request: request{
				id: "c23432zy5pk8",
				body: UpdateMessageRequest{
					Text: "<text>",
				},
			},
			behavior: behavior{
				fn: func(service *mock.MessageService, request request, response response) {
					service.
						EXPECT().
						Update(
							testifymock.Anything,
							request.id,
							model.MessageBusinessObject{
								Text: request.body.Text,
							},
						).
						Return(
							response.body,
							response.err,
						)
				},
			},
			response: response{
				body: model.MessageDataTransferObject{
					Text: "<text>",
				},
				status: http.StatusOK,
			},
		},
	}

	service := mock.NewMessageService(t)

	router := chi.NewRouter()

	handler := NewMessageHandler(service)
	handler.Routes(router)

	server := httptest.NewServer(router)
	defer server.Close()

	expect := httpexpect.New(t, server.URL)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.behavior.fn(service, test.request, test.response)

			expect.
				PUT("/v1/messages/{id}").
				WithPath("id", test.request.id).
				WithJSON(test.request.body).
				Expect().
				Status(test.response.status).
				JSON().
				Equal(test.response.body)
		})
	}
}

func TestMessageHandler_Delete(t *testing.T) {
	type request struct {
		id string
	}

	type response struct {
		status int
		err    error
	}

	type behavior struct {
		fn func(service *mock.MessageService, request request, response response)
	}

	tests := map[string]struct {
		request  request
		behavior behavior
		response response
	}{
		"OK": {
			request: request{
				id: "c23432zy5pk8",
			},
			behavior: behavior{
				fn: func(service *mock.MessageService, request request, response response) {
					service.
						EXPECT().
						Delete(
							testifymock.Anything,
							request.id,
						).
						Return(
							response.err,
						)
				},
			},
			response: response{
				status: http.StatusNoContent,
			},
		},
	}

	service := mock.NewMessageService(t)

	router := chi.NewRouter()

	handler := NewMessageHandler(service)
	handler.Routes(router)

	server := httptest.NewServer(router)
	defer server.Close()

	expect := httpexpect.New(t, server.URL)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.behavior.fn(service, test.request, test.response)

			expect.
				DELETE("/v1/messages/{id}").
				WithPath("id", test.request.id).
				Expect().
				Status(test.response.status)
		})
	}
}

func TestMessageHandler_List(t *testing.T) {
	type request struct{}

	type response struct {
		body   []model.MessageDataTransferObject
		status int
		err    error
	}

	type behavior struct {
		fn func(service *mock.MessageService, request request, response response)
	}

	tests := map[string]struct {
		request  request
		behavior behavior
		response response
	}{
		"OK": {
			behavior: behavior{
				fn: func(service *mock.MessageService, request request, response response) {
					service.
						EXPECT().
						List(testifymock.Anything).
						Return(
							response.body,
							response.err,
						)
				},
			},
			response: response{
				body: []model.MessageDataTransferObject{
					{
						Text: "<text>",
					},
				},
				status: http.StatusOK,
			},
		},
	}

	service := mock.NewMessageService(t)

	router := chi.NewRouter()

	handler := NewMessageHandler(service)
	handler.Routes(router)

	server := httptest.NewServer(router)
	defer server.Close()

	expect := httpexpect.New(t, server.URL)

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.behavior.fn(service, test.request, test.response)

			expect.
				GET("/v1/messages").
				Expect().
				Status(test.response.status).
				JSON().
				Equal(test.response.body)
		})
	}
}
