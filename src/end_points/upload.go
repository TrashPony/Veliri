package end_points

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)

		floreTexture, handler := parseFile(r, "terrainTexture") // файлы земли
		if floreTexture != nil {
			addFile(floreTexture, handler, "./src/static/assets/map/terrain/")
		}
		objectTexture, handler := parseFile(r, "objectTexture") // файлы обьекта
		if objectTexture != nil {
			addFile(objectTexture, handler, "./src/static/assets/map/objects/")
		}

		animateSprite, handler := parseFile(r, "animateSprite") // файлы анимации
		if animateSprite != nil {
			addFile(animateSprite, handler, "./src/static/assets/map/animate/")
		}
	}
}

func parseFile(r *http.Request, fileName string) (multipart.File, *multipart.FileHeader) {
	floreTexture, handler, err := r.FormFile(fileName)
	if err != nil {
		return nil, nil
	}
	defer floreTexture.Close()

	return floreTexture, handler
}

func addFile(file multipart.File, handler *multipart.FileHeader, path string) {
	f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	io.Copy(f, file)
}
