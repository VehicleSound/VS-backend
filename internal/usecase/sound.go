package usecase

import "net/http"

type SoundUseCase struct {
	r Repository
}

func (u *SoundUseCase) Upload(file http.File) {

}
