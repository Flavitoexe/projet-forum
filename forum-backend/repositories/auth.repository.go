package repositories

import "database/sql"

type AuthRepository struct {
	db *sql.DB
}

func InitAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}
