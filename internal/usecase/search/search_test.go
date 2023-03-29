package search

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	domain2 "github.com/timickb/transport-sound/internal/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/repository/memory"
	"testing"
)

func TestSearchNothingFound(t *testing.T) {
	r := memory.NewRepository()
	searchService := New(r, logrus.New())

	sounds, err := searchService.Search(&Request{
		Name:       "name",
		TagIds:     make([]string, 0),
		VehicleIds: make([]string, 0),
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(sounds) != 0 {
		t.Error("expected sounds len = 0")
	}
}

func TestSearchByTags(t *testing.T) {
	r := memory.NewRepository()

	user := domain2.User{
		Login:        "login",
		Email:        "email",
		PasswordHash: "hash",
		Id:           uuid.NewString(),
	}

	// create sounds author
	err := r.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	// prepare tags
	tagIds := []string{"1", "2", "3"}
	tags := []*domain2.Tag{
		{
			Id:    tagIds[0],
			Title: "tag1",
		},
		{
			Id:    tagIds[1],
			Title: "tag2",
		},
		{
			Id:    tagIds[2],
			Title: "tag3",
		},
	}

	sid1 := uuid.NewString()
	sid2 := uuid.NewString()
	sid3 := uuid.NewString()

	err = r.CreateSound(&domain2.Sound{
		Id:       sid1,
		Name:     "test sound",
		AuthorId: user.Id,
		Tags:     []*domain2.Tag{tags[0], tags[1]},
	})
	if err != nil {
		t.Error(err)
	}

	err = r.CreateSound(&domain2.Sound{
		Id:       sid2,
		Name:     "test sound 2",
		AuthorId: user.Id,
		Tags:     []*domain2.Tag{tags[1], tags[2]},
	})
	if err != nil {
		t.Error(err)
	}

	err = r.CreateSound(&domain2.Sound{
		Id:       sid3,
		Name:     "test sound 3",
		AuthorId: user.Id,
		Tags:     []*domain2.Tag{tags[0], tags[2]},
	})
	if err != nil {
		t.Error(err)
	}

}
