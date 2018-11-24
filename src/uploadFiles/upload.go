package uploadFiles

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)

		floreTexture, handler := parseFile(r, "terrainTexture") // файлы земли
		addFile(floreTexture, handler, "./src/static/assets/map/terrain/")

		objectTexture, handler := parseFile(r, "objectTexture") // файлы обьекта
		addFile(objectTexture, handler, "./src/static/assets/map/objects/")

		animateSprite, handler := parseFile(r, "animateSprite") // файлы анимации
		addFile(animateSprite, handler, "./src/static/assets/map/animate/")
	}
}

func parseFile(r *http.Request, fileName string) (multipart.File, *multipart.FileHeader) {
	floreTexture, handler, err := r.FormFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer floreTexture.Close()

	return floreTexture, handler
}

func addFile(file multipart.File, handler *multipart.FileHeader, path string) {
	f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	io.Copy(f, file)
}
