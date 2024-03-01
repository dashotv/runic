package parser

import (
	"regexp"
	"strings"

	"github.com/samber/lo"
)

var group *regexp.Regexp
var _group = `[\[\(]\s*(?P<group>[^\]\)]+?)\s*[\)\]]`
var website *regexp.Regexp
var _website = sp + `*[_-]+` + sp + `*(?P<website>[\w_-]+?)`
var uncensored *regexp.Regexp
var encodings *regexp.Regexp
var _encodings = []string{
	"10bit",
	"120fps",
	"144fps",
	"264",
	"30fps",
	"50fps",
	"60fps",
	"aac",
	"aac2",
	"ac3",
	"avc",
	"dd5.1",
	"dd5",
	"ddp2",
	"ddp5.1",
	"ddp5",
	"dts",
	"h264",
	"h265",
	"hdr",
	"hevc",
	"w4f",
	"x264",
	"x265",
	"xvid",
	"aac.2.0",
	"aac2.0",
	"h.265",
	"h.264",
}
var resolutions *regexp.Regexp
var _resolutions = []string{
	"1080",
	"2160",
	"480",
	"720",
}
var qualities *regexp.Regexp
var _qualities = []string{
	"bd",
	"bdrip",
	"bluray",
	"blurayrip",
	"brrip",
	"cam",
	"camrip",
	"dvdrip",
	"hd",
	"hdrip",
	"hdtv",
	"rip",
	"truehd",
	"ts",
	"uhd",
	"web-dl",
	"web",
	"webrip",
}
var bluray *regexp.Regexp
var _bluray = []string{
	"bd",
	"bdrip",
	"bluray",
	"blurayrip",
	"brrip",
}

func init() {
	encodings = regexp.MustCompile(`(?i)\b(` + strings.Join(_encodings, "|") + `)\b`)
	resolutions = regexp.MustCompile(`(?i)\b(` + strings.Join(_resolutions, "|") + `)[p]*\b`)
	qualities = regexp.MustCompile(`(?i)\b(` + strings.Join(_qualities, "|") + `)\b`)
	bluray = regexp.MustCompile(`(?i)\b(` + strings.Join(_bluray, "|") + `)\b`)
	uncensored = regexp.MustCompile(`(?i)\b(unc(en)*(sored)*)\b`)
	group = regexp.MustCompile(`(?i)^` + _group)
	website = regexp.MustCompile(`(?i)` + _website + `$`)
}

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
	results := website.FindStringSubmatch(title)
	if len(results) < 2 {
		return ""
	}
	return strings.ToLower(results[1])
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
