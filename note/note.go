package note

import (
	"context"
	"database/sql"
	"time"
)

type Note struct {
	Id        string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// GetNotesCount функция из модуля 1, где мы проверяли количество записей в таблице.
// Для модуля 2, данная функция не понадобится. Используется для совместимости.
func GetNotesCount(ctx context.Context, db *sql.DB) (int, error) {
	row := db.QueryRowContext(ctx, "SELECT COUNT(*) as count FROM notes")

	var totalNotes int
	if err := row.Scan(&totalNotes); err != nil {
		return 0, err
	}

	return totalNotes, nil
}
