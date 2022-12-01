package user

import (
	"context"
	"encoding/json"
	"github.com/farzanehshahi/user-kit/internal/customErrors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHandlers(r *mux.Router, endpoints Endpoints, httpLogger log.Logger) {

	opts := []httpTransport.ServerOption{
		httpTransport.ServerErrorHandler(transport.NewLogErrorHandler(httpLogger)),
		httpTransport.ServerErrorEncoder(encodeError),
	}

	r.Methods(http.MethodPost).Path("/user").Handler(httpTransport.NewServer(
		endpoints.Create,
		decodeCreateUserRequest,
		encodeResponse,
		opts...,
	))

	r.Methods(http.MethodGet).Path("/user/{id}").Handler(httpTransport.NewServer(
		endpoints.Get,
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	))

	r.Methods(http.MethodPut).Path("/user").Handler(httpTransport.NewServer(
		endpoints.Update,
		decodeUpdateUserRequest,
		encodeResponse,
		opts...,
	))

	r.Methods(http.MethodDelete).Path("/user/{id}").Handler(httpTransport.NewServer(
		endpoints.Delete,
		decodeDeleteUserRequest,
		encodeResponse,
		opts...,
	))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// encoding methods
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// decoding methods
func decodeGetUserRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	req := &GetUserRequest{
		ID: vars["id"],
	}
	return req, nil
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := &CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeUpdateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := &UpdateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeDeleteUserRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	req := &DeleteUserRequest{
		ID: vars["id"],
	}
	return req, nil
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case customErrors.ErrUserNotFound:
		w.WriteHeader(http.StatusNotFound)
	case customErrors.ErrUsernameAlreadyExists:
		w.WriteHeader(http.StatusBadRequest)
	case customErrors.ErrInvalidCredentials:
		w.WriteHeader(http.StatusBadRequest)

	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
