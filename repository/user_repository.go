package repository

import (
	"database/sql"

	"enigmacamp.com/be-enigma-laundry/model"
)

type UserRepository interface {
	Get(id string) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Get(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`SELECT id,name,email,username,role,created_at,updated_at FROM users WHERE id = $1`, id).
		Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Username,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
