package postgres

import (
	"errors"
	"fmt"
)

func (p PqRepository) CreateFile(id, ext string) error {
	_, err := p.db.Exec(`INSERT INTO files (id, ext) VALUES ($1, $2)`, id, ext)
	if err != nil {
		return fmt.Errorf("err save file info: %w", err)
	}

	return nil
}

func (p PqRepository) GetFileExtById(id string) (string, error) {
	row := p.db.QueryRow(`SELECT ext FROM files WHERE id=$1`, id)
	if row == nil {
		return "", errors.New("err read file info: no such file")
	}

	var ext string
	err := row.Scan(&ext)
	if err != nil {
		return "", fmt.Errorf("err read file info: %w", err)
	}
	return ext, nil
}
