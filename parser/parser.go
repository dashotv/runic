package parser

import (
	"regexp"
	"strings"

	"github.com/samber/lo"
)

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
	uncensored = regexp.MustCompile(`(?i)\b(unc(en)*(sored)*)\b`)
}

func isUncensored(title string) bool {
	return uncensored.MatchString(title)
}
func isBluray(title string) bool {
	return qualities.MatchString(title)
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
