package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	err := os.MkdirAll(_staticRoot+"images/os_rs", os.ModeDir)
	if err != nil {
		logIt("Unable to setup environment : ", err)
		return
	}
	ticker := time.NewTicker(30 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				emptyDir(_staticRoot + "images/os_rs")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	http.HandleFunc("/", handler)
	http.HandleFunc(_staticURL, staticHandler)
	http.ListenAndServe(":8080", nil)
	close(quit)
}

func emptyDir(dir string) {
	d, err := os.Open(dir)
	if err != nil {
		logIt("Unable to open directory : ", err)
		return
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		logIt("Unable to read directory : ", err)
		return
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			logIt("Unable to remove files : ", name, err)
			return
		}
	}
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
			f, err := http.Dir(_staticRoot + "images/os_rs").Open(sf)
			if err != nil {
				logIt("Creating new player image : ", err)
				player := strings.TrimSuffix(sf, ".png")
				statimage.NewRuneStat(player,
					statapi.OldSchoolAPIHandler(player),
					_staticRoot)
			}
			if f == nil {
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
