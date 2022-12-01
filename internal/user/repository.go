package user

import (
	"context"
	"database/sql"
	"github.com/farzanehshahi/user-kit/internal/customErrors"
	"github.com/farzanehshahi/user-kit/internal/entity"
	"github.com/go-kit/log"
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return repository{db: db, logger: logger}
}

func (repo repository) Create(ctx context.Context, user *entity.User) error {

	// todo: first need to check the username not be exist!
	// q := `select exists(select 1 from users where id=$1)`

	query := `INSERT INTO "users" (username, password) 
			VALUES ($1, $2)
			returning id;`

	var insertedId string
	if err := repo.db.QueryRow(query, user.Username, user.Password).Scan(&insertedId); err != nil {
		return err
	}
	user.ID = insertedId
	return nil
}

func (repo repository) Get(ctx context.Context, id string) (entity.User, error) {

	query := `SELECT * FROM users WHERE id=$1;`

	var username, password string
	err := repo.db.QueryRow(query, id).Scan(&id, &username, &password)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, customErrors.ErrUserNotFound
		}
		return entity.User{}, err
	}

	return entity.User{
		ID:       id,
		Username: username,
		Password: password,
	}, nil
}

func (repo repository) Update(ctx context.Context, id string, updatedUser *entity.User) error {

	query := `UPDATE users
			SET username = $2, password = $3
			WHERE id = $1;`

	res, err := repo.db.Exec(query, id, updatedUser.Username, updatedUser.Password)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil && count != 0 {
		return err
	}
	if count == 0 {
		return customErrors.ErrUserNotFound
	}

	return nil
}

func (repo repository) Delete(ctx context.Context, id string) error {

	query := `DELETE FROM users WHERE id = $1;`
	res, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil && count != 0 {
		return err
	}
	if count == 0 {
		return customErrors.ErrUserNotFound
	}
	return nil
}
