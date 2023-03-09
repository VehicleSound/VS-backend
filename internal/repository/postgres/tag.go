package postgres

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/timickb/transport-sound/internal/domain"
)

func (p PqRepository) CreateTag(title string) (*domain.Tag, error) {
	st := `INSERT INTO tags (id, title) VALUES ($1, $2)`
	id := uuid.NewString()

	_, err := p.db.Exec(st, id, title)
	if err != nil {
		return nil, fmt.Errorf("db insertion err: %w", err)
	}

	return &domain.Tag{
		Id:    id,
		Title: title,
	}, nil
}

func (p PqRepository) GetTagById(id string) (*domain.Tag, error) {
	row := p.db.QueryRow("SELECT * FROM tags WHERE id=$1", id)

	if row == nil {
		return nil, errors.New("no such tag")
	}

	tag := &domain.Tag{}
	err := row.Scan(&tag.Id, &tag.Title)

	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (p PqRepository) GetTagByTitle(title string) (*domain.Tag, error) {
	row := p.db.QueryRow("SELECT * FROM tags WHERE title=$1", title)

	if row == nil {
		return nil, errors.New("no such tag")
	}

	tag := &domain.Tag{}
	err := row.Scan(&tag.Id, &tag.Title)

	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (p PqRepository) GetAllTags() ([]*domain.Tag, error) {
	rows, err := p.db.Query("SELECT * FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*domain.Tag
	var id, title string

	for rows.Next() {
		err := rows.Scan(&id, &title)
		if err != nil {
			return nil, err
		}
		result = append(result, &domain.Tag{Id: id, Title: title})
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p PqRepository) GetTagsForSound(soundId string) ([]*domain.Tag, error) {
	tRows, err := p.db.Query("SELECT t.id, t.title FROM sound_tags st JOIN tags t ON t.id = st.tag_id WHERE st.sound_id=$1", soundId)

	if err != nil {
		return nil, err
	}

	var tags []*domain.Tag
	for tRows.Next() {
		tag := &domain.Tag{}

		err := tRows.Scan(&tag.Id, &tag.Title)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	if err := tRows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}
