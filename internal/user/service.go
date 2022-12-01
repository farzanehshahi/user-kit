package user

import (
	"context"
	"fmt"
	"github.com/farzanehshahi/user-kit/internal/customErrors"
	"github.com/farzanehshahi/user-kit/internal/entity"
	"github.com/farzanehshahi/user-kit/pkg/validator"
	"github.com/go-kit/log"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo      Repository
	validator validator.Validator
	logger    log.Logger
}

// NewService creates a new service.
func NewService(repo Repository, validator validator.Validator, logger log.Logger) Service {
	return service{repo, validator, logger}
}

// create a new user
func (s service) Create(ctx context.Context, reqUser *entity.User) error {

	// validate request --> it seems the go-kit validates the requests ==> TODO: need check
	//if ok, err := s.validator.Validate(reqUser); !ok {
	//	return err
	//}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqUser.Password), 10)
	if err != nil {
		return customErrors.ErrInvalidCredentials
	}

	reqUser.Password = string(hashedPassword)

	if err := s.repo.Create(ctx, reqUser); err != nil {
		return err
	}

	fmt.Println("user in service:", reqUser)

	s.logger.Log("user created successfully.")
	return nil
}

// Get
func (s service) Get(ctx context.Context, id string) (entity.User, error) {

	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s service) Update(ctx context.Context, id string, username, password string) error {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		// s.logger.Error("error while hashing password", err)
		return err
	}

	updateUser := &entity.User{
		Username: username,
		Password: string(hashedPassword),
	}
	return s.repo.Update(ctx, id, updateUser)
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
