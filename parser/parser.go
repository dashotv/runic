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
func getResolution(title string) (string, string) {
	results := resolutions.FindStringSubmatch(title)
	if len(results) == 0 {
		return title, ""
	}
	title = strings.Replace(title, results[0], "", 1)
	return title, results[1]
}
func getEncodings(title string) (string, []string) {
	results := encodings.FindAllString(title, -1)
	matches := lo.Map(results, func(s string, i int) string {
		return strings.ToLower(s)
	})
	for _, match := range results {
		title = strings.Replace(title, match, "", 1)
	}
	return title, matches
}
func getQualities(title string) (string, string) {
	results := qualities.FindAllString(title, -1)
	matches := lo.Map(results, func(s string, i int) string {
		return strings.ToLower(s)
	})
	for _, match := range results {
		title = strings.Replace(title, match, "", 1)
	}
	return title, strings.Join(matches, ", ")
}
func getGroup(title string) (string, string) {
	results := group.FindStringSubmatch(title)
	if len(results) < 2 {
		return title, ""
	}
	title = strings.Replace(title, results[0], "", 1)
	return title, strings.ToLower(results[1])
}
func getWebsite(title string) (string, string) {
	results := regexList(website, title)
	if len(results) < 1 {
		return title, ""
	}
	// if getResolution(results[0]) != "" {
	// 	return "", ""
	// }
	title = strings.Replace(title, results[0], "", 1)
	return title, strings.ToLower(results[0])
}
func getChecksum(title string) (string, string) {
	results := checksum.FindStringSubmatch(title)
	if len(results) < 1 {
		return title, ""
	}
	title = strings.Replace(title, results[0], "", 1)
	return title, strings.ToLower(results[1])
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
			if params["episode"] != "" {
				params["episode"] = CleanEpisode(params["episode"])
			}
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
	params["raw"] = results[0]
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			params[name] = results[i]
		}
	}
	return params
}

func Parse(title, catType string) (*TorrentInfo, error) {
	title = strings.ToLower(title)
	bluray := isBluray(title)
	uncensored := isUncensored(title)
	title, group := getGroup(title)
	title, website := getWebsite(title)
	title, qualities := getQualities(title)
	title, encodings := getEncodings(title)
	title, res := getResolution(title)
	title, checksum := getChecksum(title)

	title = strings.TrimLeft(title, " _-")
	title = strings.TrimRight(title, " _-")
	title = strings.TrimSpace(title)
	info := &TorrentInfo{
		Resolution: res,
		Quality:    qualities,
		Encodings:  encodings,
		Bluray:     bluray,
		Uncensored: uncensored,
		Group:      group,
		Website:    website,
		Checksum:   checksum,
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

func ParseTitle(title, catType string) (*TorrentInfo, error) {
	info := &TorrentInfo{}
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
