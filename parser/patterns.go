package parser

import (
	"regexp"
	"strings"
)

func init() {
	encodings = regexp.MustCompile(`(?i)\b(` + strings.Join(_encodings, "|") + `)\b`)
	resolutions = regexp.MustCompile(`(?i)[\[]*\b(?:\d{3,4}x)*(` + strings.Join(_resolutions, "|") + `)[p\]]*`)
	qualities = regexp.MustCompile(`(?i)\b(` + strings.Join(_qualities, "|") + `)\b`)
	bluray = regexp.MustCompile(`(?i)\b(` + strings.Join(_bluray, "|") + `)\b`)
	uncensored = regexp.MustCompile(`(?i)\b(unc(en)*(sored)*)\b`)
	group = regexp.MustCompile(`(?i)^` + _group)
	website = regexp.MustCompile(`(?i)` + _website + `$`)
	checksum = regexp.MustCompile(`(?i)\[([0-9a-f]{8})\]`)
}

var sp = `[\s\-_.]`
var episode = `(?:(e(p)*(isode)*` + sp + `+)*(?P<episode>\d{1,3})\b)` // |(?P<episode>\d{4}.\d{2}.\d{2}))
var episodeDate = `((?:(e(p)*(isode)*` + sp + `+)*(?P<episode>\d{1,3})\b)|(?P<episode>\d{4}.\d{2}.\d{2}))`
var se = `s(eason)*` + sp + `*(?P<season>\d+)` + sp + `*e(pisode)*` + sp + `*(?P<episode>\d{1,3})\b`
var sx = `\b(?P<season>\d+)x(?P<episode>\d{1,3})\b`
var volume = `(?P<volume>(vol(ume)*|part)` + sp + `+(?P<volnum>\d+))+`
var title = `(?P<title>.*?)`
var titleGreedy = `(?P<title>.*)`
var year = `[\(\[]*(?P<year>\d{4})[\(\[]*`
var date = `(?P<date>\d{4}-\d{2}-\d{2})`

// -[\w]+\[(\w+)\]
// \[([\w]+)\]
// sp+`*\[([\w]+)\]`
// \[([\w]+)\]\.\w{3,4}
// ((-*\[*(\w+)\]*)|(-[\w]+\[(\w+)\])|(\[([\w]+)\])|(\[([\w]+)\]\.\w{3,4}))$
// (?P<group>-*\s*[\[]*([\w]+))[\]]*$
// var _website = `(?:(?:-*\[*(?P<website>\w+)\]*)|(?:-[\w]+\[(?P<website>\w+)\])|(?:\[(?P<website>[\w]+)\])|(?:\[(?P<website>[\w]+)\]\.\w{3,4}))`
var group *regexp.Regexp
var _group = `[\[\(]\s*(?P<group>[^\]\)]+?)\s*[\)\]]`
var website *regexp.Regexp
var _website = sp + `*[_-]+` + sp + `*(?P<website>[\w_-]+?)`
var uncensored *regexp.Regexp
var checksum *regexp.Regexp
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
	"dtshd",
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
	"sbs",
	"3d",
	"vhs",
	"flac",
	"cr",
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
	"web dl",
	"webdl",
	"web",
	"webrip",
	"web-rip",
	"web rip",
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
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+` + se),
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+` + sx),
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+` + year + sp + `+` + episodeDate),
	regexp.MustCompile(`(?i)^` + title + `` + sp + `+` + episodeDate),
	regexp.MustCompile(`(?i)^` + se + sp + `+-` + sp + `+` + titleGreedy),
	regexp.MustCompile(`(?i)^` + se + sp + `*` + titleGreedy),
	regexp.MustCompile(`(?i)^` + se),
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
	regexp.MustCompile(`(?i)^` + se + sp + `+-` + sp + `+` + titleGreedy),
	regexp.MustCompile(`(?i)^` + se + sp + `*` + titleGreedy),
	regexp.MustCompile(`(?i)^` + se),
	regexp.MustCompile(`(?i)^` + _group + `` + sp + `*` + titleGreedy + ``),
	regexp.MustCompile(`(?i)^` + titleGreedy + ``),
}
