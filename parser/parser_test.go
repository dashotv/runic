package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEncodings(t *testing.T) {
	testdata := []struct {
		subject string
		want    []string
	}{
		{"The.Jungle.Book.2016.3D.1080p.BRRip.SBS.x264.AAC-ETRG", []string{"x264", "aac"}},
		{"Hercules (2014) 1080p BrRip H264 - YIFY", []string{"h264"}},
		{"Dawn.of.the.Planet.of.the.Apes.2014.HDRip.XViD-EVO", []string{"xvid"}},
		{"The Big Bang Theory S08E06 HDTV XviD-LOL [eztv]", []string{"xvid"}},
		{"22 Jump Street (2014) 720p BrRip x264 - YIFY", []string{"x264"}},
		{"Hercules.2014.EXTENDED.1080p.WEB-DL.DD5.1.H264-RARBG", []string{"dd5.1", "h264"}},
		{"Hercules.2014.Extended.Cut.HDRip.XViD-juggs[ETRG]", []string{"xvid"}},
		{"Hercules (2014) WEBDL DVDRip XviD-MAX", []string{"xvid"}},
		{"WWE Hell in a Cell 2014 PPV WEB-DL x264-WD -={SPARROW}=-", []string{"x264"}},
		{"UFC.179.PPV.HDTV.x264-Ebi[rartv]", []string{"x264"}},
		{"Marvels Agents of S H I E L D S02E05 HDTV x264-KILLERS [eztv]", []string{"x264"}},
		{"X-Men.Days.of.Future.Past.2014.1080p.WEB-DL.DD5.1.H264-RARBG", []string{"dd5.1", "h264"}},
		{"Guardians Of The Galaxy 2014 R6 720p HDCAM x264-JYK", []string{"x264"}},
		{"Marvel's.Agents.of.S.H.I.E.L.D.S02E01.Shadows.1080p.WEB-DL.DD5.1", []string{"dd5.1"}},
	}
	for _, tt := range testdata {
		t.Run(tt.subject, func(t *testing.T) {
			assert.Equal(t, tt.want, getEncodings(tt.subject))
		})
	}
}

func TestNewResolutions(t *testing.T) {
	testdata := []struct {
		subject string
		want    string
	}{
		{"The.Jungle.Book.2016.3D.1080p.BRRip.SBS.x264.AAC-ETRG", "1080"},
		{"Hercules (2014) 1080p BrRip H264 - YIFY", "1080"},
		{"Dawn.of.the.Planet.of.the.Apes.2014.HDRip.XViD-EVO", ""},
		{"The Big Bang Theory S08E06 HDTV XviD-LOL [eztv]", ""},
		{"22 Jump Street (2014) 720p BrRip x264 - YIFY", "720"},
		{"Hercules.2014.EXTENDED.1080p.WEB-DL.DD5.1.H264-RARBG", "1080"},
		{"Hercules.2014.Extended.Cut.HDRip.XViD-juggs[ETRG]", ""},
		{"Hercules (2014) WEBDL DVDRip XviD-MAX", ""},
		{"WWE Hell in a Cell 2014 PPV WEB-DL x264-WD -={SPARROW}=-", ""},
		{"UFC.179.PPV.HDTV.x264-Ebi[rartv]", ""},
		{"Marvels Agents of S H I E L D S02E05 HDTV x264-KILLERS [eztv]", ""},
		{"X-Men.Days.of.Future.Past.2014.1080p.WEB-DL.DD5.1.H264-RARBG", "1080"},
		{"Guardians Of The Galaxy 2014 R6 720p HDCAM x264-JYK", "720"},
		{"Marvel's.Agents.of.S.H.I.E.L.D.S02E01.Shadows.1080p.WEB-DL.DD5.1", "1080"},
	}
	for _, tt := range testdata {
		t.Run(tt.subject, func(t *testing.T) {
			assert.Equal(t, tt.want, getResolution(tt.subject))
		})
	}
}

func TestNewQualities(t *testing.T) {
	testdata := []struct {
		subject string
		want    string
	}{
		{"The.Jungle.Book.2016.3D.1080p.BRRip.SBS.x264.AAC-ETRG", "brrip"},
		{"Hercules (2014) 1080p BrRip H264 - YIFY", "brrip"},
		{"Dawn.of.the.Planet.of.the.Apes.2014.HDRip.XViD-EVO", "hdrip"},
		{"The Big Bang Theory S08E06 HDTV XviD-LOL [eztv]", "hdtv"},
		{"22 Jump Street (2014) 720p BrRip x264 - YIFY", "brrip"},
		{"Hercules.2014.EXTENDED.1080p.WEB-DL.DD5.1.H264-RARBG", "web-dl"},
		{"Hercules.2014.Extended.Cut.HDRip.XViD-juggs[ETRG]", "hdrip"},
		{"Hercules (2014) WEBDL DVDRip XviD-MAX", "dvdrip"},
		{"WWE Hell in a Cell 2014 PPV WEB-DL x264-WD -={SPARROW}=-", "web-dl"},
		{"UFC.179.PPV.HDTV.x264-Ebi[rartv]", "hdtv"},
		{"Marvels Agents of S H I E L D S02E05 HDTV x264-KILLERS [eztv]", "hdtv"},
		{"X-Men.Days.of.Future.Past.2014.1080p.WEB-DL.DD5.1.H264-RARBG", "web-dl"},
		{"Guardians Of The Galaxy 2014 R6 720p HDCAM x264-JYK", ""},
		{"Marvel's.Agents.of.S.H.I.E.L.D.S02E01.Shadows.1080p.WEB-DL.DD5.1", "web-dl"},
	}
	for _, tt := range testdata {
		t.Run(tt.subject, func(t *testing.T) {
			assert.Equal(t, tt.want, getQuality(tt.subject))
		})
	}
}

func TestNewUncensored(t *testing.T) {
	testdata := []struct {
		subject string
		want    bool
	}{
		{"The.Jungle.Book.2016.3D.1080p.BRRip.SBS.x264.AAC-ETRG", false},
		{"Hercules (2014) 1080p BrRip H264 - YIFY", false},
		{"[AniSuki] Ayakashi Triangle Volume 5 (BD) (x265 HEVC OPUS) (Uncensored)", true},
		{"[AE] Tokyo Ghoul - [Batch] [UNCEN] [720p]", true},
	}
	for _, tt := range testdata {
		t.Run(tt.subject, func(t *testing.T) {
			assert.Equal(t, tt.want, isUncensored(tt.subject))
		})
	}
}

func TestNewBluray(t *testing.T) {
	testdata := []struct {
		subject string
		want    bool
	}{
		{"The.Jungle.Book.2016.3D.1080p.BRRip.SBS.x264.AAC-ETRG", true},
		{"Hercules (2014) 1080p BrRip H264 - YIFY", true},
		{"[AniSuki] Ayakashi Triangle Volume 5 (BD) (x265 HEVC OPUS) (Uncensored)", true},
		{"[AE] Tokyo Ghoul - [Batch] [UNCEN] [720p]", false},
	}
	for _, tt := range testdata {
		t.Run(tt.subject, func(t *testing.T) {
			assert.Equal(t, tt.want, isBluray(tt.subject))
		})
	}
}
