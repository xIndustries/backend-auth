package repositories

import (
	"database/sql"

	"github.com/xIndustries/BandRoom/backend-auth/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// ✅ CreateUser - Inserts a new user into the database
func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (id, auth0_id, email, username, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.DB.Exec(query, user.ID, user.Auth0ID, user.Email, user.Username, user.CreatedAt)
	return err
}

// ✅ GetUser - Retrieves a user by their Auth0 ID
func (r *UserRepository) GetUser(auth0ID string) (*models.User, error) {
	query := `SELECT id, auth0_id, email, username, created_at FROM users WHERE auth0_id = $1`
	row := r.DB.QueryRow(query, auth0ID)

	var user models.User
	err := row.Scan(&user.ID, &user.Auth0ID, &user.Email, &user.Username, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ✅ UpdateUsername - Updates the username for a user
func (r *UserRepository) UpdateUsername(auth0ID, username string) error {
	query := `UPDATE users SET username = $1 WHERE auth0_id = $2`
	_, err := r.DB.Exec(query, username, auth0ID)
	return err
}

// ✅ UpdateUserEmail - Updates the email for a user
func (r *UserRepository) UpdateUserEmail(auth0ID, email string) error {
	query := `UPDATE users SET email = $1 WHERE auth0_id = $2`
	_, err := r.DB.Exec(query, email, auth0ID)
	return err
}

// ✅ DeleteUser - Removes a user by their Auth0 ID
func (r *UserRepository) DeleteUser(auth0ID string) error {
	query := `DELETE FROM users WHERE auth0_id = $1`
	_, err := r.DB.Exec(query, auth0ID)
	return err
}
