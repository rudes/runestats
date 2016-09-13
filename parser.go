package main

import (
	"html/template"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	_oldSchoolURL = "http://services.runescape.com/m=hiscore_oldschool/hiscorepersonal.ws?user1="
)

// Stat structure for housing Rune Stat data
type Stat struct {
	Type, Picture template.HTML
	Value         string
}

func oldSchoolHandler(p string) []Stat {
	if p == "" {
		return nil
	}
	doc, err := goquery.NewDocument(_oldSchoolURL + p[1:])
	if err != nil {
		logIt("Couldn't query webpage : ", err)
		return nil
	}
	var rows []string
	doc.Find("td").Each(func(i int, s *goquery.Selection) {
		if i > 12 && i < 131 {
			res, _ := s.Html()
			rows = append(rows, res)
		}
	})
	if rows == nil {
		return nil
	}
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
