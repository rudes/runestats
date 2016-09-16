package statapi

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	_oldSchoolURL    = "http://services.runescape.com/m=hiscore_oldschool/hiscorepersonal.ws?user1="
	_oldSchoolAPIURL = "http://services.runescape.com/m=hiscore_oldschool/index_lite.ws?player="
)

func OldSchoolAPIHandler(p string) []string {
	res, err := http.Get(_oldSchoolAPIURL + p)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	var stats []string
	for _, row := range strings.Split(string(body), "\n") {
		stat := newStatFromAPI(row)
		if stat != "" {
			stats = append(stats, stat)
		}
	}
	return stats
}

func newStatFromAPI(row string) string {
	if row == "" {
		return ""
	}
	stats := strings.Split(row, ",")
	if len(stats) > 0 {
		return stats[1]
	}
	return ""
}
