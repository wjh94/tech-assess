package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
)

func (r *Repository) GetUserByID(ctx context.Context, userID int64) (*User, error) {
	var user User
	err := r.db.QueryRowContext(ctx, "SELECT id, username, email FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("no user found") // User not found
	}
	if err != nil {
		logger.Error("Failed to get user by ID: ", err)
		return nil, err
	}
	return &user, nil
}
