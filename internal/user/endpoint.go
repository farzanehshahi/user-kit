package user

import (
	"context"
	"github.com/farzanehshahi/user-kit/internal/entity"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Create endpoint.Endpoint
	Get    endpoint.Endpoint
	Update endpoint.Endpoint
	Delete endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}
}

// create user endpoint
type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	User *entity.User `json:"user,omitempty"`
	Err  error        `json:"error,omitempty"`
}

//
func makeCreateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*CreateUserRequest)
		user := entity.User{
			Username: req.Username,
			Password: req.Password,
		}

		err := s.Create(ctx, &user)
		return CreateUserResponse{User: &user}, err
	}
}

// get user endpoint
type GetUserRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetUserResponse struct {
	User *entity.User `json:"user,omitempty"`
	Err  error        `json:"error,omitempty"`
}

func makeGetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetUserRequest)

		user, err := s.Get(ctx, req.ID)
		return GetUserResponse{User: &user}, err
	}
}

// update user endpoint
type UpdateUserRequest struct {
	ID              string `json:"id" validate:"required"`
	UpdatedUsername string `json:"username" validate:"required"`
	UpdatedPassword string `json:"password" validate:"required"`
}

type UpdateUserResponse struct {
	// User *entity.User `json:"user,omitempty"`
	Message string `json:"message,omitempty"`
	Err     error  `json:"error,omitempty"`
}

// todo: user should be returned by update function from repo
func makeUpdateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*UpdateUserRequest)

		err := s.Update(ctx, req.ID, req.UpdatedUsername, req.UpdatedPassword)
		return UpdateUserResponse{Message: "user successfully updated."}, err
	}
}

// delete user endpoint
type DeleteUserRequest struct {
	ID string `json:"id" validate:"required"`
}

type DeleteUserResponse struct {
	Message string `json:"message,omitempty"`
	Err     error  `json:"error,omitempty"`
}

func makeDeleteEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*DeleteUserRequest)

		err := s.Delete(ctx, req.ID)
		return DeleteUserResponse{Message: "user successfully deleted."}, err
	}
}
