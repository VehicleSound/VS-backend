package domain

type Sound struct {
	Id          string
	AuthorId    string
	Name        string
	Description string
	PictureUrl  string
	SoundUrl    string
	Tags        []Tag
	Vehicle     *Vehicle
}
