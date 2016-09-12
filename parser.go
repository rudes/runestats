package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	_oldSchoolURL = "http://services.runescape.com/m=hiscore_oldschool/hiscorepersonal.ws?user1="
	_staticURL    = "/templates/static/"
	_staticRoot   = "/home/rudes/go/src/github.com/rudes/runestats/templates/static/"
	_templateRoot = "/home/rudes/go/src/github.com/rudes/runestats/templates/"
)

// Stat structure for housing Rune Stat data
type Stat struct {
	Type, Picture template.HTML
	Value         string
}

// Context structure for rendering templates
type Context struct {
	Stats []Stat
}

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
	}
	http.NotFound(w, req)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, _templateRoot+"index.html")
			return
		}
		render(w, r, oldSchoolHandler(r.URL.Path))
	}
}

func oldSchoolHandler(p string) []Stat {
	doc, err := goquery.NewDocument(_oldSchoolURL + p[1:])
	if err != nil {
		log.Fatal(err)
	}
	var rows []string
	doc.Find("td").Each(func(i int, s *goquery.Selection) {
		if i > 12 && i < 131 {
			res, _ := s.Html()
			rows = append(rows, res)
		}
	})
	var stats []Stat
	stats = append(stats, newStat(rows[0], "", rows[2]))
	for i := 5; i < 118; i = i + 5 {
		stats = append(stats, newStat(rows[i], rows[i-1], rows[i+2]))
	}

	return stats
}

func newStat(t string, p string, v string) Stat {
	s := Stat{
		Type:    template.HTML(strings.Replace(t, "\n", "", -1)),
		Picture: template.HTML(p),
		Value:   v}
	s.Type = template.HTML(strings.Replace(string(s.Type), "overall.ws",
		"http://services.runescape.com/m=hiscore_oldschool/overall.ws", -1))
	s.Picture = template.HTML(strings.Replace(string(s.Picture),
		"http://www.runescape.com/img/rsp777/hiscores",
		"templates/static/images/", -1))
	return s
}

func render(w http.ResponseWriter, r *http.Request, stats []Stat) {
	ctx := Context{Stats: stats}
	t, err := template.ParseFiles(_templateRoot+"base.tmpl", _templateRoot+"header.tmpl", _templateRoot+"content.tmpl")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = t.Execute(w, ctx)
	if err != nil {
		log.Fatal(err)
	}
}
