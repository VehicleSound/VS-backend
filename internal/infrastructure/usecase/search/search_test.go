package search

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/timickb/transport-sound/internal/infrastructure/domain"
	"github.com/timickb/transport-sound/internal/infrastructure/repository/memory"
	"github.com/timickb/transport-sound/internal/infrastructure/repository/mock"
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

func TestSearchByName(t *testing.T) {
	r := mock.NewRepository()
	searchService := New(r, logrus.New())

	names := []string{"sound 1", "sound 2", "sound 3", "sound 4"}

	if err := r.CreateTestSounds(names); err != nil {
		t.Fatal(err)
	}

	sounds, err := searchService.Search(&Request{
		Name:       "sound",
		TagIds:     make([]string, 0),
		VehicleIds: make([]string, 0),
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(sounds) != len(names) {
		t.Error("expected result size 4")
	}

	sounds, err = searchService.Search(&Request{
		Name:       "1",
		TagIds:     make([]string, 0),
		VehicleIds: make([]string, 0),
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(sounds) != 1 {
		t.Error("expected result size 1")
	}
}

func TestSearchByTags(t *testing.T) {
	r := memory.NewRepository()

	// create sounds author
	uid, err := r.CreateUser("login", "email", "hash")
	if err != nil {
		t.Fatal(err)
	}

	// prepare tags
	tagIds := []string{"1", "2", "3"}
	tags := []*domain.Tag{
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

	err = r.CreateSound(&domain.Sound{
		Id:       sid1,
		Name:     "test sound",
		AuthorId: uid,
		Tags:     []*domain.Tag{tags[0], tags[1]},
	})
	if err != nil {
		t.Error(err)
	}

	err = r.CreateSound(&domain.Sound{
		Id:       sid2,
		Name:     "test sound 2",
		AuthorId: uid,
		Tags:     []*domain.Tag{tags[1], tags[2]},
	})
	if err != nil {
		t.Error(err)
	}

	err = r.CreateSound(&domain.Sound{
		Id:       sid3,
		Name:     "test sound 3",
		AuthorId: uid,
		Tags:     []*domain.Tag{tags[0], tags[2]},
	})
	if err != nil {
		t.Error(err)
	}

}
