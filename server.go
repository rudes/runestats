package main

import (
	"html/template"
	"io"
	"net/http"
	"time"
)

// Context structure for rendering templates
type Context struct {
	Stats []Stat
}

const (
	_staticURL    = "/templates/static/"
	_staticRoot   = "/go/src/github.com/rudes/runestats/templates/static/"
	_templateRoot = "/go/src/github.com/rudes/runestats/templates/"
)

func main() {
	//:8080/niriviaa
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
		logIt("Unable to service content : ", err)
	}
	http.NotFound(w, req)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, _templateRoot+"index.html")
		} else {
			stats := oldSchoolHandler(r.URL.Path)
			if stats != nil {
				render(w, r, stats)
			}
		}
	}
}

func render(w http.ResponseWriter, r *http.Request, stats []Stat) {
	ctx := Context{Stats: stats}
	t, err := template.ParseFiles(_templateRoot+"base.tmpl", _templateRoot+"header.tmpl", _templateRoot+"content.tmpl")
	if err != nil {
		logIt("Unable to create template : ", err)
		return
	}
	err = t.Execute(w, ctx)
	if err != nil {
		logIt("Unable to execute template : ", err)
	}
}
