package repository

import (
	"database/sql"
	"time"

	"enigmacamp.com/be-enigma-laundry/model"
)

type UserRepository interface {
	Get(id string) (model.User, error)
	// get by username ketika login
	GetByUsername(username string) (model.User, error)
	Create(payload model.User) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Create(payload model.User) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`
	INSERT INTO users (name, email, username, password, role, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, email, username, role, created_at, updated_at
	`, payload.Name, payload.Email, payload.Username, payload.Password, payload.Role, time.Now()).Scan(
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

func (u *userRepository) GetByUsername(username string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`SELECT id, name, email, username, password, role FROM users WHERE username = $1 OR email $1`, username).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.Role)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
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
