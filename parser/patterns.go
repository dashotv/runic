package parser

import (
	"regexp"
	"strings"
)

func init() {
	encodings = regexp.MustCompile(`(?i)\b(` + strings.Join(_encodings, "|") + `)\b`)
	resolutions = regexp.MustCompile(`(?i)\b(` + strings.Join(_resolutions, "|") + `)[p]*\b`)
	qualities = regexp.MustCompile(`(?i)\b(` + strings.Join(_qualities, "|") + `)\b`)
	bluray = regexp.MustCompile(`(?i)\b(` + strings.Join(_bluray, "|") + `)\b`)
	uncensored = regexp.MustCompile(`(?i)\b(unc(en)*(sored)*)\b`)
	group = regexp.MustCompile(`(?i)^` + _group)
	website = regexp.MustCompile(`(?i)` + _website + `$`)
}

var sp = `[\s_.]`
var episode = `(e(p)*(isode)*` + sp + `+)*(?P<episode>\d{1,3})\b`
var se = `s(eason)*` + sp + `*(?P<season>\d+)` + sp + `*e(pisode)*` + sp + `*(?P<episode>\d{1,3})\b`
var sx = `\b(?P<season>\d+)x(?P<episode>\d{1,3})\b`
var volume = `(?P<volume>(vol(ume)*|part)` + sp + `+(?P<volnum>\d+))+`
var title = `(?P<title>.*?)`
var titleGreedy = `(?P<title>.*)`
var year = `[\(\[]*(?P<year>\d{4})[\(\[]*`
var date = `(?P<date>\d{4}-\d{2}-\d{2})`

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

var regexes = []*regexp.Regexp{
	regexesAnime[13],
	regexesAnime[14],
	regexesAnime[15],
	regexesAnime[16],
	regexesMovies[0],
	regexesTV[0],
	regexesTV[1],
	regexesTV[2],
}

var regexesMovies = []*regexp.Regexp{
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+` + year + `.*` + _website + `$`),
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+` + year + ``),
	regexp.MustCompile(`(?i)^` + titleGreedy + _website + `$`),
	regexp.MustCompile(`(?i)^` + titleGreedy),
}

var regexesTV = []*regexp.Regexp{
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+` + date),
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+` + se),
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+` + sx),
}

var regexesAnime = []*regexp.Regexp{
	regexp.MustCompile(`(?i)^` + _group + `` + sp + `*` + title + `` + sp + `+-` + sp + `+` + volume + ``),
	regexp.MustCompile(`(?i)^` + _group + `` + sp + `*` + title + `` + sp + `+-` + sp + `+` + se + ``),
	regexp.MustCompile(`(?i)^` + _group + `` + sp + `*` + title + `` + sp + `+-` + sp + `+` + episode + ``),
	regexp.MustCompile(`(?i)^` + _group + `` + sp + `*` + title + `-` + se + ``),
	regexp.MustCompile(`(?i)^` + _group + `` + sp + `*` + title + `-` + episode + ``),
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+-` + sp + `+` + volume + ``),
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+-` + sp + `+` + se + ``),
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+-` + sp + `+` + episode + ``),
	regexp.MustCompile(`(?i)^` + title + `-` + episode + ``),
	regexp.MustCompile(`(?i)^` + _group + `` + sp + `*` + title + `` + sp + `+` + volume + ``),
	regexp.MustCompile(`(?i)^` + _group + `` + sp + `*` + title + `` + sp + `+` + se + ``),
	regexp.MustCompile(`(?i)^` + _group + `` + sp + `*` + title + `` + sp + `+` + episode + ``),
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+` + volume + ``),
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+` + se + ``),
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+` + episode + ``),
	regexp.MustCompile(`(?i)^` + _group + `` + sp + `*` + titleGreedy + ``),
	regexp.MustCompile(`(?i)^` + titleGreedy + ``),
}
