package parser

import "regexp"

var sp = `[\s_.]`
var episode = `(e(p)*(isode)*` + sp + `+)*(?P<episode>\d{1,3})\b`
var se = `s(eason)*` + sp + `*(?P<season>\d+)` + sp + `*e(pisode)*` + sp + `*(?P<episode>\d{1,3})\b`
var sx = `\b(?P<season>\d+)x(?P<episode>\d{1,3})\b`
var volume = `(?P<volume>(vol(ume)*|part)` + sp + `+(?P<volnum>\d+))+`
var title = `(?P<title>.*?)`
var titleGreedy = `(?P<title>.*)`
var year = `[\(\[]*(?P<year>\d{4})[\(\[]*`
var date = `(?P<date>\d{4}-\d{2}-\d{2})`

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
