package controller

import (
	"net/http"

	"github.com/olahol/go-imageupload"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	img, err := imageupload.Process(r, "file")
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	thumb, err := imageupload.ThumbnailJPEG(img, 300, 300, 80)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	thumb.Write(w)
}
