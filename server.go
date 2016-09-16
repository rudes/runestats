package main

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rudes/runestats/statapi"
	"github.com/rudes/runestats/statimage"
)

const (
	_staticURL    = "/templates/static/"
	_staticRoot   = "/go/src/github.com/rudes/runestats/templates/static/"
	_templateRoot = "/go/src/github.com/rudes/runestats/templates/"
)

func main() {
	// TODO: Initialize app by building a directory
	//	    tree for the images to be stored
	http.HandleFunc("/", handler)
	http.HandleFunc(_staticURL, staticHandler)
	http.ListenAndServe(":8080", nil)
}

func staticHandler(w http.ResponseWriter, req *http.Request) {
	sf := req.URL.Path[len(_staticURL):]
	if len(sf) != 0 {
		f, err := http.Dir(_staticRoot).Open(sf)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, sf, time.Now(), content)
			return
		}
		logIt("Unable to serve content : ", err)
	}
	http.NotFound(w, req)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, _templateRoot+"index.html")
		} else if strings.Contains(r.URL.Path, "png") {
			sf := r.URL.Path[1:]
			f, err := http.Dir(_staticRoot + "images/").Open(sf)
			if err != nil {
				logIt("Creating new player image : ", err)
			}
			if f == nil {
				player := strings.TrimSuffix(sf, ".png")
				statimage.NewRuneStat(player, statapi.OldSchoolAPIHandler(player))
				f, err = os.Open(_staticRoot + "images/" + sf)
				if err != nil {
					http.NotFound(w, r)
				}
			}
			content := io.ReadSeeker(f)
			http.ServeContent(w, r, sf, time.Now(), content)

		} else {
			http.NotFound(w, r)
		}
	}
}
