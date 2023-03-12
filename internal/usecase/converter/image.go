package converter

import (
	"github.com/sunshineplan/imgconv"
	"image"
	"path/filepath"
)

func HandleAndSaveImage(savePath, id string, img image.Image) error {
	img = imgconv.Resize(img, &imgconv.ResizeOption{
		Width:  512,
		Height: 512,
	})

	path := filepath.Join(savePath, id+".jpg")

	if err := imgconv.Save(path, img, &imgconv.FormatOption{
		Format: imgconv.JPEG,
	}); err != nil {
		return err
	}

	return nil
}
