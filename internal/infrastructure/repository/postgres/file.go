package postgres

import (
	"errors"
	"fmt"
)

func (p PqRepository) CreateFile(id, ext, sum string) error {
	_, err := p.db.Exec(`INSERT INTO files (id, ext, sum) VALUES ($1, $2, $3)`, id, ext, sum)
	if err != nil {
		return fmt.Errorf("err save file info: %w", err)
	}

	return nil
}

func (p PqRepository) GetFileIdBySum(sum string) (string, error) {
	row := p.db.QueryRow(`SELECT id FROM files WHERE sum=$1`, sum)

	var id string
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
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
