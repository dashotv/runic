package parser

import (
	"regexp"
	"strings"

	"github.com/samber/lo"
)

func isUncensored(title string) bool {
	return uncensored.MatchString(title)
}
func isBluray(title string) bool {
	return bluray.MatchString(title)
}
func getResolution(title string) string {
	results := resolutions.FindStringSubmatch(title)
	if len(results) == 0 {
		return ""
	}
	return results[1]
}
func getEncodings(title string) []string {
	results := encodings.FindAllString(title, -1)
	return lo.Map(results, func(s string, i int) string {
		return strings.ToLower(s)
	})
}
func getQuality(title string) string {
	results := qualities.FindAllString(title, -1)
	if len(results) == 0 {
		return ""
	}
	return strings.ToLower(results[0])
}
func getGroup(title string) string {
	results := group.FindStringSubmatch(title)
	if len(results) < 2 {
		return ""
	}
	return strings.ToLower(results[1])
}
func getWebsite(title string) string {
	results := regexList(website, title)
	if len(results) < 1 {
		return ""
	}
	if getResolution(results[0]) != "" {
		return ""
	}
	return strings.ToLower(results[0])
}

func parseTitle(title string, catType string) (int, map[string]string) {
	list := regexes
	switch catType {
	case "movies":
		list = regexesMovies
	case "tv":
		list = regexesTV
	case "anime":
		list = regexesAnime
	}

	for i, r := range list {
		params := regexParams(r, title)
		if params != nil {
			return i, params
		}
	}
	return -1, nil
}

func regexList(r *regexp.Regexp, title string) []string {
	list := []string{}
	results := r.FindStringSubmatch(title)
	if len(results) == 0 {
		return nil
	}
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			list = append(list, results[i])
		}
	}
	return lo.Compact(list)
}

func regexParams(r *regexp.Regexp, title string) map[string]string {
	results := r.FindStringSubmatch(title)
	if len(results) == 0 {
		return nil
	}
	params := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			params[name] = results[i]
		}
	}
	return params
}

func Parse(title, catType string) (*TorrentInfo, error) {
	info := &TorrentInfo{
		// Title:      ,
		Resolution: getResolution(title),
		Quality:    getQuality(title),
		Encodings:  getEncodings(title),
		Bluray:     isBluray(title),
		Uncensored: isUncensored(title),
		Group:      getGroup(title),
		Website:    getWebsite(title),
	}
	i, params := parseTitle(title, catType)
	if i >= 0 {
		info.Title = CleanTitle(params["title"])
		info.setSeason(params["season"])
		info.setEpisode(params["episode"])
		info.setYear(params["year"])
		// info.Volume = params["volume"]
	}
	return info, nil
}
