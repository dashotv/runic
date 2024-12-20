package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var err error
var updateGolden bool
var titlesAnime []string
var titlesMovies []string
var titlesTv []string

func init() {
	updateGolden = os.Getenv("UPDATE_GOLDEN") == "true"
	titlesAnime, err = loadTitles("anime")
	if err != nil {
		panic(err)
	}
	titlesMovies, err = loadTitles("movies")
	if err != nil {
		panic(err)
	}
	titlesTv, err = loadTitles("tv")
	if err != nil {
		panic(err)
	}
}

func TestEncodings(t *testing.T) {
	testdata := []struct {
		subject   string
		encodings []string
		title     string
	}{
		{"The.Jungle.Book.2016.3D.1080p.BRRip.SBS.x264.AAC-ETRG", []string{"3d", "sbs", "x264", "aac"}, "The.Jungle.Book.2016..1080p.BRRip...-ETRG"},
		{"Hercules (2014) 1080p BrRip H264 - YIFY", []string{"h264"}, "Hercules (2014) 1080p BrRip  - YIFY"},
		{"Dawn.of.the.Planet.of.the.Apes.2014.HDRip.XViD-EVO", []string{"xvid"}, "Dawn.of.the.Planet.of.the.Apes.2014.HDRip.-EVO"},
		{"The Big Bang Theory S08E06 HDTV XviD-LOL [eztv]", []string{"xvid"}, "The Big Bang Theory S08E06 HDTV -LOL [eztv]"},
		{"22 Jump Street (2014) 720p BrRip x264 - YIFY", []string{"x264"}, "22 Jump Street (2014) 720p BrRip  - YIFY"},
		{"Hercules.2014.EXTENDED.1080p.WEB-DL.DD5.1.H264-RARBG", []string{"dd5.1", "h264"}, "Hercules.2014.EXTENDED.1080p.WEB-DL..-RARBG"},
		{"Hercules.2014.Extended.Cut.HDRip.XViD-juggs[ETRG]", []string{"xvid"}, "Hercules.2014.Extended.Cut.HDRip.-juggs[ETRG]"},
		{"Hercules (2014) WEBDL DVDRip XviD-MAX", []string{"xvid"}, "Hercules (2014) WEBDL DVDRip -MAX"},
		{"WWE Hell in a Cell 2014 PPV WEB-DL x264-WD -={SPARROW}=-", []string{"x264"}, "WWE Hell in a Cell 2014 PPV WEB-DL -WD -={SPARROW}=-"},
		{"UFC.179.PPV.HDTV.x264-Ebi[rartv]", []string{"x264"}, "UFC.179.PPV.HDTV.-Ebi[rartv]"},
		{"Marvels Agents of S H I E L D S02E05 HDTV x264-KILLERS [eztv]", []string{"x264"}, "Marvels Agents of S H I E L D S02E05 HDTV -KILLERS [eztv]"},
		{"X-Men.Days.of.Future.Past.2014.1080p.WEB-DL.DD5.1.H264-RARBG", []string{"dd5.1", "h264"}, "X-Men.Days.of.Future.Past.2014.1080p.WEB-DL..-RARBG"},
		{"Guardians Of The Galaxy 2014 R6 720p HDCAM x264-JYK", []string{"x264"}, "Guardians Of The Galaxy 2014 R6 720p HDCAM -JYK"},
		{"Marvel's.Agents.of.S.H.I.E.L.D.S02E01.Shadows.1080p.WEB-DL.DD5.1", []string{"dd5.1"}, "Marvel's.Agents.of.S.H.I.E.L.D.S02E01.Shadows.1080p.WEB-DL."},
	}
	for _, tt := range testdata {
		t.Run(tt.subject, func(t *testing.T) {
			title, encodings := getEncodings(tt.subject)
			assert.Equal(t, tt.encodings, encodings)
			assert.Equal(t, tt.title, title)
		})
	}
}

func TestResolutions(t *testing.T) {
	testdata := []struct {
		subject string
		res     string
		title   string
	}{
		{"The.Jungle.Book.2016.3D.1080p.BRRip.SBS.x264.AAC-ETRG", "1080", "The.Jungle.Book.2016.3D..BRRip.SBS.x264.AAC-ETRG"},
		{"Hercules (2014) [1080p] BrRip H264 - YIFY", "1080", "Hercules (2014)  BrRip H264 - YIFY"},
		{"Dawn.of.the.Planet.of.the.Apes.2014.HDRip.XViD-EVO", "", "Dawn.of.the.Planet.of.the.Apes.2014.HDRip.XViD-EVO"},
		{"The Big Bang Theory S08E06 HDTV XviD-LOL [eztv]", "", "The Big Bang Theory S08E06 HDTV XviD-LOL [eztv]"},
		{"22 Jump Street (2014) 720p BrRip x264 - YIFY", "720", "22 Jump Street (2014)  BrRip x264 - YIFY"},
		{"Hercules.2014.EXTENDED.1080p.WEB-DL.DD5.1.H264-RARBG", "1080", "Hercules.2014.EXTENDED..WEB-DL.DD5.1.H264-RARBG"},
		{"Hercules.2014.Extended.Cut.HDRip.XViD-juggs[ETRG]", "", "Hercules.2014.Extended.Cut.HDRip.XViD-juggs[ETRG]"},
		{"Hercules (2014) WEBDL DVDRip XviD-MAX", "", "Hercules (2014) WEBDL DVDRip XviD-MAX"},
		{"X-Men.Days.of.Future.Past.2014.1080p.WEB-DL.DD5.1.H264-RARBG", "1080", "X-Men.Days.of.Future.Past.2014..WEB-DL.DD5.1.H264-RARBG"},
		{"Guardians Of The Galaxy 2014 R6 720p HDCAM x264-JYK", "720", "Guardians Of The Galaxy 2014 R6  HDCAM x264-JYK"},
		{"Marvel's.Agents.of.S.H.I.E.L.D.S02E01.Shadows.1080p.WEB-DL.DD5.1", "1080", "Marvel's.Agents.of.S.H.I.E.L.D.S02E01.Shadows..WEB-DL.DD5.1"},
		{"[Erai-raws] Shangri-La Frontier - Kusogee Hunter, Kamige ni Idoman to Su - 20 [720p][Multiple Subtitle] [ENG][POR-BR][SPA-LA][SPA][ARA][FRE][GER][ITA][RUS]", "720", "[Erai-raws] Shangri-La Frontier - Kusogee Hunter, Kamige ni Idoman to Su - 20 [Multiple Subtitle] [ENG][POR-BR][SPA-LA][SPA][ARA][FRE][GER][ITA][RUS]"},
		{"[FSP DN] 牧神记 Mushen Ji (Tales of Herding Gods) - 02 [4K]", "2160", "[FSP DN] 牧神记 Mushen Ji (Tales of Herding Gods) - 02 "},
	}
	for _, tt := range testdata {
		t.Run(tt.subject, func(t *testing.T) {
			title, res := getResolution(tt.subject)
			assert.Equal(t, tt.res, res)
			assert.Equal(t, tt.title, title)
		})
	}
}

func TestQualities(t *testing.T) {
	testdata := []struct {
		subject string
		quality string
		title   string
	}{
		{"The.Jungle.Book.2016.3D.1080p.BRRip.SBS.x264.AAC-ETRG", "brrip", "The.Jungle.Book.2016.3D.1080p..SBS.x264.AAC-ETRG"},
		{"Hercules (2014) 1080p BrRip H264 - YIFY", "brrip", "Hercules (2014) 1080p  H264 - YIFY"},
		{"Dawn.of.the.Planet.of.the.Apes.2014.HDRip.XViD-EVO", "hdrip", "Dawn.of.the.Planet.of.the.Apes.2014..XViD-EVO"},
		{"The Big Bang Theory S08E06 HDTV XviD-LOL [eztv]", "hdtv", "The Big Bang Theory S08E06  XviD-LOL [eztv]"},
		{"22 Jump Street (2014) 720p BrRip x264 - YIFY", "brrip", "22 Jump Street (2014) 720p  x264 - YIFY"},
		{"Hercules.2014.EXTENDED.1080p.WEB-DL.DD5.1.H264-RARBG", "web-dl", "Hercules.2014.EXTENDED.1080p..DD5.1.H264-RARBG"},
		{"Hercules.2014.Extended.Cut.HDRip.XViD-juggs[ETRG]", "hdrip", "Hercules.2014.Extended.Cut..XViD-juggs[ETRG]"},
		{"Hercules (2014) WEBDL DVDRip XviD-MAX", "webdl, dvdrip", "Hercules (2014)   XviD-MAX"},
		{"WWE Hell in a Cell 2014 PPV WEB-DL x264-WD -={SPARROW}=-", "web-dl", "WWE Hell in a Cell 2014 PPV  x264-WD -={SPARROW}=-"},
		{"UFC.179.PPV.HDTV.x264-Ebi[rartv]", "hdtv", "UFC.179.PPV..x264-Ebi[rartv]"},
		{"Marvels Agents of S H I E L D S02E05 HDTV x264-KILLERS [eztv]", "hdtv", "Marvels Agents of S H I E L D S02E05  x264-KILLERS [eztv]"},
		{"X-Men.Days.of.Future.Past.2014.1080p.WEB-DL.DD5.1.H264-RARBG", "web-dl", "X-Men.Days.of.Future.Past.2014.1080p..DD5.1.H264-RARBG"},
		{"Guardians Of The Galaxy 2014 R6 720p HDCAM x264-JYK", "", "Guardians Of The Galaxy 2014 R6 720p HDCAM x264-JYK"},
		{"Marvel's.Agents.of.S.H.I.E.L.D.S02E01.Shadows.1080p.WEB-DL.DD5.1", "web-dl", "Marvel's.Agents.of.S.H.I.E.L.D.S02E01.Shadows.1080p..DD5.1"},
	}
	for _, tt := range testdata {
		t.Run(tt.subject, func(t *testing.T) {
			title, quality := getQualities(tt.subject)
			assert.Equal(t, tt.quality, quality)
			assert.Equal(t, tt.title, title)
		})
	}
}

func TestUncensored(t *testing.T) {
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

func TestBluray(t *testing.T) {
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

func TestGroup(t *testing.T) {
	testdata := []struct {
		subject string
		group   string
		title   string
	}{
		{"The.Jungle.Book.2016.3D.1080p.BRRip.SBS.x264.AAC-ETRG", "", "The.Jungle.Book.2016.3D.1080p.BRRip.SBS.x264.AAC-ETRG"},
		{"Hercules (2014) 1080p BrRip H264 - YIFY", "", "Hercules (2014) 1080p BrRip H264 - YIFY"},
		{"[AniSuki] Ayakashi Triangle Volume 5 (BD) (x265 HEVC OPUS) (Uncensored)", "anisuki", " Ayakashi Triangle Volume 5 (BD) (x265 HEVC OPUS) (Uncensored)"},
		{"[AE] Tokyo Ghoul - [Batch] [UNCEN] [720p]", "ae", " Tokyo Ghoul - [Batch] [UNCEN] [720p]"},
	}
	for _, tt := range testdata {
		t.Run(tt.subject, func(t *testing.T) {
			title, group := getGroup(tt.subject)
			assert.Equal(t, tt.group, group)
			assert.Equal(t, tt.title, title)
		})
	}
}

func TestWebsite(t *testing.T) {
	testdata := []struct {
		subject string
		website string
		title   string
	}{
		{"The.Jungle.Book.2016.3D.1080p.BRRip.SBS.x264.AAC-ETRG", "etrg", "The.Jungle.Book.2016.3D.1080p.BRRip.SBS.x264.AAC-"},
		{"Hercules (2014) 1080p BrRip H264 - YIFY", "yify", "Hercules (2014) 1080p BrRip H264 - "},
		{"[AniSuki] Ayakashi Triangle Volume 5 (BD) (x265 HEVC OPUS) (Uncensored)", "", "[AniSuki] Ayakashi Triangle Volume 5 (BD) (x265 HEVC OPUS) (Uncensored)"},
		{"[AE] Tokyo Ghoul - [Batch] [UNCEN] [720p]", "", "[AE] Tokyo Ghoul - [Batch] [UNCEN] [720p]"},
		// TODO: fix this
		{"thomas.and.friends.s19e09_s20e14.convert.hdtv.x264-w4f[eztv].mkv", "", "thomas.and.friends.s19e09_s20e14.convert.hdtv.x264-w4f[eztv].mkv"},
		{"Doctor.Who.2005.8x11.Dark.Water.720p.HDTV.x264-FoV[rartv]", "", "Doctor.Who.2005.8x11.Dark.Water.720p.HDTV.x264-FoV[rartv]"},
		{"The Simpsons S26E05 HDTV x264 PROPER-LOL [eztv]", "", "The Simpsons S26E05 HDTV x264 PROPER-LOL [eztv]"},
		{"The Flash 2014 S01E01 HDTV x264-LOL[ettv]", "", "The Flash 2014 S01E01 HDTV x264-LOL[ettv]"},
	}
	for _, tt := range testdata {
		t.Run(tt.subject, func(t *testing.T) {
			title, website := getWebsite(tt.subject)
			assert.Equal(t, tt.website, website)
			assert.Equal(t, tt.title, title)
		})
	}
}

func TestParser_anime(t *testing.T) {
	for i, tt := range titlesAnime {
		t.Run(fmt.Sprintf("%03d %s", i, tt), func(t *testing.T) {
			info, err := Parse(tt, "anime")
			assert.NoError(t, err)
			assert.NotEmpty(t, info.Title)

			err = saveGolden("anime", i, info)
			assert.NoError(t, err)

			gold, err := loadGolden("anime", i)
			require.NoError(t, err)
			assert.Equal(t, gold, info)
		})
	}
}
func TestParser_movies(t *testing.T) {
	for i, tt := range titlesMovies {
		t.Run(fmt.Sprintf("%03d %s", i, tt), func(t *testing.T) {
			info, err := Parse(tt, "movies")
			assert.NoError(t, err)
			assert.NotEmpty(t, info.Title)

			err = saveGolden("movies", i, info)
			assert.NoError(t, err)

			gold, err := loadGolden("movies", i)
			assert.NoError(t, err)
			assert.Equal(t, gold, info)
		})
	}
}
func TestParser_tv(t *testing.T) {
	for i, tt := range titlesTv {
		t.Run(fmt.Sprintf("%03d %s", i, tt), func(t *testing.T) {
			info, err := Parse(tt, "tv")
			assert.NoError(t, err)
			assert.NotEmpty(t, info.Title)

			err = saveGolden("tv", i, info)
			assert.NoError(t, err)

			gold, err := loadGolden("tv", i)
			assert.NoError(t, err)
			assert.Equal(t, gold, info)
		})
	}
}

func saveGolden(cat string, i int, info *TorrentInfo) error {
	if !updateGolden {
		return nil
	}
	fmt.Printf("SAVING GOLDEN %s %d\n", cat, i)
	f, err := os.Create(fmt.Sprintf("testdata/%s_%03d.json", cat, i))
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(info)
}

func loadGolden(cat string, i int) (*TorrentInfo, error) {
	f, err := os.Open(fmt.Sprintf("testdata/%s_%03d.json", cat, i))
	if err != nil {
		return nil, err
	}

	defer f.Close()
	dec := json.NewDecoder(f)
	golden := &TorrentInfo{}
	err = dec.Decode(golden)
	if err != nil {
		return nil, err
	}
	return golden, err
}

func loadTitles(cat string) ([]string, error) {
	var titles []string
	file, err := os.Open(fmt.Sprintf("testdata/titles_%s.txt", cat))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		titles = append(titles, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return titles, nil
}
