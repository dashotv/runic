package parser

import (
	"regexp"
	"strings"
)

var cleanEpisodeRegex = regexp.MustCompile(`[\W\.]+`)

var replacements []*Replacement
var words = []string{
	"aac",
	"amzn",
	"aoz",
	"av1",
	"axxo",
	"bluray",
	"brrip",
	"ddp5",
	"divx",
	"dsnp",
	"dvdrip",
	"eac",
	"eng",
	"eztv",
	"fxg",
	"fxm",
	"h264",
	"h265",
	"hbomax",
	"hdrip",
	"hdtv",
	"hevc",
	"klaxxon",
	"kyr",
	"mkv",
	"mp4",
	"multisub",
	"mxmg",
	"netflix",
	"notv",
	"opus",
	"pdtv",
	"proper",
	"r5",
	"sub",
	"tvrip",
	"ultrahd",
	"uncen",
	"uncensored",
	"web",
	"webdl",
	"webrip",
	"ws",
	"x264",
	"x265",
	"xvid",
	"480",
	"480p",
	"720",
	"720p",
	"1080",
	"1080p",
	"2160",
	"2160p",
	"4k",
	"aac.*2.*0",
	"aac2.*0",
	"h.*265",
	"h.*264",
	"dvd.*rip",
	"web.*dl",
	"ddp5.*1",
	"dual.*audio",
	"multi.*audio",
	"multi.*sub[s]*",
	"blu.*ray",
	"audios",
	"eng",
	"english",
	"jpn",
	"jap",
	"japanese",
	"rus",
	"russian",
	"spa",
	"spanish",
	"hindi",
}

func init() {
	replacements = []*Replacement{
		{Regex: regexp.MustCompile(`\'(\w{1,2})\b`), Sub: "$1"},
		{Regex: regexp.MustCompile(`[\W\.\_]+`), Sub: " "},
	}

	for _, w := range words {
		replacements = append(replacements, &Replacement{
			Regex: regexp.MustCompile(`\b` + w + `\b`),
			Sub:   "",
		})
	}

	replacements = append(replacements, []*Replacement{
		{Regex: regexp.MustCompile(`\s+`), Sub: " "},
		{Regex: regexp.MustCompile(`^\s+`), Sub: ""},
		{Regex: regexp.MustCompile(`\s+$`), Sub: ""},
	}...)
}

type Replacement struct {
	Regex *regexp.Regexp
	Sub   string
}

func CleanTitle(title string) string {
	for _, r := range replacements {
		title = r.Regex.ReplaceAllString(title, r.Sub)
	}
	return strings.ToLower(title)
}

func CleanEpisode(episode string) string {
	return cleanEpisodeRegex.ReplaceAllString(episode, "")
}
