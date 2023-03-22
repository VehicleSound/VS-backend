package sound

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/repository/memory"
	"testing"
)

func TestCreateSoundWithoutTags(t *testing.T) {
	r := memory.NewRepository()

	uid, err := r.CreateUser("login", "email", "pwd_hash")
	if err != nil {
		t.Fatal(err)
	}

	if err := r.CreateFile("1", ".jpg"); err != nil {
		t.Fatal(err)
	}
	if err := r.CreateFile("2", ".mp3"); err != nil {
		t.Fatal(err)
	}

	soundService := New(r, logrus.New())

	sound := &domain.Sound{
		Id:            uuid.NewString(),
		Name:          "sound",
		Description:   "desc",
		AuthorId:      uid,
		PictureFileId: "1",
		SoundFileId:   "2",
	}
	sid, err := soundService.CreateSound(sound, []string{})

	if err != nil {
		t.Fatal(err)
	}

	savedSound, err := soundService.GetSoundById(sid)
	if err != nil {
		t.Fatal(err)
	}

	if sound.Name != savedSound.Name {
		t.Error("name field corrupted")
	}
	if sound.Description != savedSound.Description {
		t.Error("description field corrupted")
	}
	if sound.AuthorId != savedSound.AuthorId {
		t.Error("author_id field corrupted")
	}
	if sound.PictureFileId != savedSound.PictureFileId {
		t.Error("picture_file_id field corrupted")
	}
	if sound.SoundFileId != savedSound.SoundFileId {
		t.Error("sound_file_id field corrupted")
	}
}
