package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	_player = "niriviaa"
	_url    = "http://services.runescape.com/m=hiscore_oldschool/hiscorepersonal.ws?user1=" + _player
)

type stat struct {
	Type, Picture, Value string
}

func main() {
	doc, err := goquery.NewDocument(_url)
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
	var stats []stat
	stats = append(stats, newStat(rows[0], "", rows[2]))
	for i := 5; i < 118; i = i + 5 {
		stats = append(stats, newStat(rows[i], rows[i-1], rows[i+2]))
	}

	fmt.Printf("%v", stats)
}

func newStat(t string, p string, v string) stat {
	s := stat{Type: strings.Replace(t, "\n", "", -1),
		Picture: p,
		Value:   v}
	s.Type = strings.Replace(s.Type, "overall.ws",
		"http://services.runescape.com/m=hiscore_oldschool/overall.ws", -1)
	return s
}
