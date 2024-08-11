package user

import (
	"context"
	"database/sql"
	"log"
	"newproject/domain/entities"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Create(ctx context.Context, user entities.User) error {
	query := `
	INSERT INTO user (name) VALUES (?)
	`

	_, err := r.db.ExecContext(ctx, query, user.Name)
	if err != nil {
		log.Println("[Create] Error ExecContext", err)
		return err
	}

	return nil
}
