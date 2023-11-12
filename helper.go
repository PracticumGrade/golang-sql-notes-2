package main

import (
	"encoding/csv"
	"io"
	"os"
	"time"

	"github.com/PracticumGrade/golang-sql-notes/note"
)

func readNotesCSV(csvFile string) ([]note.Note, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var notes []note.Note

	const (
		Id        = 0
		Title     = 1
		Content   = 2
		CreatedAt = 3
		UpdatedAt = 4
	)

	r := csv.NewReader(file)
	if _, err := r.Read(); err != nil {
		return nil, err
	}

	for {
		l, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		n := note.Note{
			Id:      l[Id],
			Title:   l[Title],
			Content: l[Content],
		}

		if n.CreatedAt, err = time.Parse(time.DateTime, l[CreatedAt]); err != nil {
			return nil, err
		}

		// значение UpdatedAt в базе может быть nil
		if l[UpdatedAt] == "" {
			n.UpdatedAt = nil
		} else {
			updatedAt, err := time.Parse(time.DateTime, l[UpdatedAt])
			if err != nil {
				return nil, err
			}
			n.UpdatedAt = &updatedAt
		}

		notes = append(notes, n)
	}

	return notes, nil
}
