package postgres

import (
	"errors"
	"github.com/timickb/transport-sound/internal/domain"
)

func (p PqRepository) AddTagToSound(soundId, tagId string) error {
	query := `INSERT INTO sound_tags (sound_id, tag_id) VALUES($1, $2)`

	if _, err := p.db.Exec(query, soundId, tagId); err != nil {
		return err
	}

	return nil
}

func (p PqRepository) CreateSound(sound *domain.Sound) error {
	query := `INSERT INTO sounds
    		(id, name, description, author_id, vehicle_id, sound_file_id, picture_file_id) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)`

	if _, err := p.db.Exec(query,
		sound.Id,
		sound.Name,
		sound.Description,
		sound.AuthorId,
		sound.VehicleId,
		sound.SoundFileId,
		sound.PictureFileId); err != nil {
		return err
	}

	return nil
}

func (p PqRepository) GetSoundById(id string) (*domain.Sound, error) {
	query := `SELECT s.id, s.name, s.description, s.author_id, s.vehicle_id, s.picture_file_id, s.sound_file_id, u.login, v.name 
			FROM sounds s
			JOIN users u on u.id = s.author_id
			JOIN vehicles v on v.id = s.vehicle_id
			WHERE s.id=$1`

	row := p.db.QueryRow(query, id)

	if row == nil {
		return nil, errors.New("sound not found")
	}

	sound := &domain.Sound{}

	if err := row.Scan(
		&sound.Id,
		&sound.Name,
		&sound.Description,
		&sound.AuthorId,
		&sound.VehicleId,
		&sound.PictureFileId,
		&sound.SoundFileId,
		&sound.AuthorLogin,
		&sound.VehicleName); err != nil {
		return nil, err
	}

	return sound, nil
}

func (p PqRepository) GetAllSounds() ([]*domain.Sound, error) {
	sQuery := `SELECT s.id, s.name, s.description, s.author_id, s.vehicle_id, s.picture_file_id, s.sound_file_id, u.login, v.name FROM sounds s 
    	JOIN vehicles v on v.id = s.vehicle_id
    	JOIN users u on s.author_id = u.id`

	sRows, err := p.db.Query(sQuery)
	if err != nil {
		return nil, err
	}

	tagsMap, err := p.getTagsMap()

	var sounds []*domain.Sound

	for sRows.Next() {
		sound := &domain.Sound{}
		err := sRows.Scan(
			&sound.Id,
			&sound.Name,
			&sound.Description,
			&sound.AuthorId,
			&sound.VehicleId,
			&sound.PictureFileId,
			&sound.SoundFileId,
			&sound.AuthorLogin,
			&sound.VehicleName)
		if err != nil {
			return nil, err
		}

		if tagsMap[sound.Id] != nil {
			sound.Tags = tagsMap[sound.Id]
		} else {
			sound.Tags = []*domain.Tag{}
		}

		sounds = append(sounds, sound)
	}

	if err := sRows.Err(); err != nil {
		return nil, err
	}

	return sounds, nil
}

func (p PqRepository) GetSounds(limit, offset int) ([]*domain.Sound, error) {
	//TODO implement me
	panic("implement me")
}

func (p PqRepository) GetSoundsNameLike(name string) ([]*domain.Sound, error) {
	rows, err := p.db.Query(
		`SELECT s.id, s.name, s.description, s.author_id, s.vehicle_id, s.sound_file_id, s.picture_file_id, u.login, v.name
				FROM sounds s 
    			JOIN users u on s.author_id = u.id 
         		JOIN vehicles v on v.id = s.vehicle_id
         		WHERE lower(s.name) LIKE '%' || $1 || '%'`, name)

	if err != nil {
		return nil, err
	}

	var result []*domain.Sound

	for rows.Next() {
		s := &domain.Sound{}
		err := rows.Scan(
			&s.Id,
			&s.Name,
			&s.Description,
			&s.AuthorId,
			&s.VehicleId,
			&s.SoundFileId,
			&s.PictureFileId,
			&s.AuthorLogin,
			&s.VehicleName)
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (p PqRepository) GetSoundsByTagId(tagId string) ([]*domain.Sound, error) {
	rows, err := p.db.Query(
		`SELECT s.id, 
       				  s.name, 
       				  s.description, 
       				  s.vehicle_id, 
       				  s.author_id, 
       				  s.picture_file_id, 
       				  s.sound_file_id, 
       				  u.login, 
       				  v.name 
				FROM sound_tags st 
    			JOIN sounds s on s.id = st.sound_id 
				JOIN vehicles v on v.id = s.vehicle_id
				JOIN users u on u.id = s.author_id
                WHERE tag_id=$1`, tagId)
	if err != nil {
		return nil, err
	}

	sounds := make([]*domain.Sound, 0)

	for rows.Next() {
		s := &domain.Sound{}
		if err := rows.Scan(
			&s.Id,
			&s.Name,
			&s.Description,
			&s.VehicleId,
			&s.AuthorId,
			&s.PictureFileId,
			&s.SoundFileId,
			&s.AuthorLogin,
			&s.VehicleName,
		); err != nil {
			return nil, err
		}
		sounds = append(sounds, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sounds, nil
}

func (p PqRepository) GetSoundsByVehicleId(vehicleId string) ([]*domain.Sound, error) {
	rows, err := p.db.Query(
		`SELECT s.id, 
       				  s.name, 
       				  s.description, 
       				  s.vehicle_id, 
       				  s.author_id, 
       				  s.picture_file_id, 
       				  s.sound_file_id, 
       				  u.login, 
       				  v.name FROM sounds s
         JOIN vehicles v on v.id = s.vehicle_id
         JOIN users u on u.id = s.author_id
         WHERE vehicle_id=$1`, vehicleId)

	if err != nil {
		return nil, err
	}

	sounds := make([]*domain.Sound, 0)

	for rows.Next() {
		s := &domain.Sound{}
		if err := rows.Scan(
			&s.Id,
			&s.Name,
			&s.Description,
			&s.VehicleId,
			&s.AuthorId,
			&s.PictureFileId,
			&s.SoundFileId,
			&s.AuthorLogin,
			&s.VehicleName,
		); err != nil {
			return nil, err
		}
		sounds = append(sounds, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sounds, nil
}

func (p PqRepository) getTagsMap() (map[string][]*domain.Tag, error) {
	tQuery := `SELECT st.sound_id, st.tag_id, t.title FROM sound_tags st
		JOIN tags t ON t.id = st.tag_id`

	tRows, err := p.db.Query(tQuery)
	if err != nil {
		return nil, err
	}

	tagsMap := make(map[string][]*domain.Tag)
	var soundId string

	for tRows.Next() {
		st := &domain.Tag{}
		err := tRows.Scan(&soundId, &st.Id, &st.Title)
		if err != nil {
			return nil, err
		}
		tagsMap[soundId] = append(tagsMap[soundId], st)
	}

	if err := tRows.Err(); err != nil {
		return nil, err
	}

	return tagsMap, nil
}
