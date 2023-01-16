package mikerybkadev

import (
	"fmt"
	"net/http"
)

type App struct{}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("mikerybka.dev"))
	} else {
		path := r.URL.Path[1:]
		fmt.Fprintf(w, "<meta name=\"go-import\" content=\"mikerybka.dev/%s git https://github.com/mikerybka/go-%s\" />", path, path)
	}
}
