package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	_url = "http://services.runescape.com/m=hiscore_oldschool/hiscorepersonal.ws?user1=niriviaa"
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
	overall := newStat(rows[0], "", rows[2])
	attack := newStat(rows[5], rows[4], rows[7])
	fmt.Printf("%s\n", overall.Type)
	fmt.Printf("%s\n", attack.Picture)
	// var i int
	// for _, a := range rows {
	// 	fmt.Println(i, a)
	// 	i++
	// }
}

func newStat(t string, p string, v string) stat {
	s := stat{Type: strings.Replace(t, "\n", "", -1),
		Picture: p,
		Value:   v}
	s.Type = strings.Replace(s.Type, "overall.ws",
		"http://services.runescape.com/m=hiscore_oldschool/overall.ws", -1)
	return s
}
