package newznab

// fork of https://github.com/mrobinsn/go-newznab/tree/master/newznab

import (
	"crypto/tls"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

const (
	apiPath = "/api"
	rssPath = "/rss"
)

var sizeRegex = regexp.MustCompile(`(\d+\.?\d*)\s([KMG]iB)`)

// Client is a type for interacting with a newznab or torznab api
type Client struct {
	apikey     string
	apiBaseURL string
	apiUserID  int
	client     *http.Client
}

// New returns a new instance of Client
func New(baseURL string, apikey string, userID int, insecure bool) *Client {
	ret := &Client{
		apikey:     apikey,
		apiBaseURL: baseURL,
		apiUserID:  userID,
	}
	if insecure {
		ret.client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}}
	} else {
		ret.client = &http.Client{}
	}
	return ret
}

// SearchWithTVRage returns NZBs for the given parameters
func (c *Client) SearchWithTVRage(categories []int, tvRageID int, season int, episode int) ([]*NZB, error) {
	return c.search(url.Values{
		"rid":     []string{strconv.Itoa(tvRageID)},
		"cat":     c.splitCats(categories),
		"season":  []string{strconv.Itoa(season)},
		"episode": []string{strconv.Itoa(episode)},
		"t":       []string{"tvsearch"},
	})
}

// SearchWithTVDB returns NZBs for the given parameters
func (c *Client) SearchWithTVDB(categories []int, tvDBID int, season int, episode int) ([]*NZB, error) {
	return c.search(url.Values{
		"tvdbid":  []string{strconv.Itoa(tvDBID)},
		"cat":     c.splitCats(categories),
		"season":  []string{strconv.Itoa(season)},
		"episode": []string{strconv.Itoa(episode)},
		"t":       []string{"tvsearch"},
	})
}

// SearchWithTVMaze returns NZBs for the given parameters
func (c *Client) SearchWithTVMaze(categories []int, tvMazeID int, season int, episode int) ([]*NZB, error) {
	return c.search(url.Values{
		"tvmazeid": []string{strconv.Itoa(tvMazeID)},
		"cat":      c.splitCats(categories),
		"season":   []string{strconv.Itoa(season)},
		"episode":  []string{strconv.Itoa(episode)},
		"t":        []string{"tvsearch"},
	})
}

// SearchWithIMDB returns NZBs for the given parameters
func (c *Client) SearchWithIMDB(categories []int, imdbID string) ([]*NZB, error) {
	return c.search(url.Values{
		"imdbid": []string{imdbID},
		"cat":    c.splitCats(categories),
		"t":      []string{"movie"},
	})
}

// SearchWithQuery returns NZBs for the given parameters
func (c *Client) SearchWithQuery(categories []int, query string, searchType string) ([]*NZB, error) {
	return c.search(url.Values{
		"q":   []string{query},
		"cat": c.splitCats(categories),
		"t":   []string{searchType},
	})
}

// LoadRSSFeed returns up to <num> of the most recent NZBs of the given categories.
func (c *Client) LoadRSSFeed(categories []int, num int) ([]*NZB, error) {
	return c.rss(url.Values{
		"num": []string{strconv.Itoa(num)},
		"t":   c.splitCats(categories),
		"dl":  []string{"1"},
	})
}

// Capabilities returns the capabilities of this tracker
func (c *Client) Capabilities() (Capabilities, error) {
	return c.caps(url.Values{
		"t": []string{"caps"},
	})
}

// LoadRSSFeedUntilNZBID fetches NZBs until a given NZB id is reached.
func (c *Client) LoadRSSFeedUntilNZBID(categories []int, num int, id string, maxRequests int) ([]*NZB, error) {
	count := 0
	var nzbs []*NZB
	for {
		partition, err := c.rss(url.Values{
			"num":    []string{strconv.Itoa(num)},
			"t":      c.splitCats(categories),
			"dl":     []string{"1"},
			"offset": []string{strconv.Itoa(num * count)},
		})
		count++
		if err != nil {
			return nil, err
		}
		for k, nzb := range partition {
			if nzb.ID == id {
				return append(nzbs, partition[:k]...), nil
			}
		}
		nzbs = append(nzbs, partition...)
		if maxRequests != 0 && count == maxRequests {
			break
		}
	}
	return nzbs, nil
}

// Details get the details of a particular nzb
func (c *Client) Details(guid string) (Details, error) {
	return c.details(url.Values{
		"t":    []string{"details"},
		"guid": []string{guid},
	})
}

func (c *Client) splitCats(cats []int) []string {
	var categories, catsOut []string
	for _, v := range cats {
		categories = append(categories, strconv.Itoa(v))
	}
	catsOut = append(catsOut, strings.Join(categories, ","))
	return catsOut
}

func (c *Client) rss(vals url.Values) ([]*NZB, error) {
	vals.Set("r", c.apikey)
	vals.Set("i", strconv.Itoa(c.apiUserID))
	return c.process(vals, rssPath)
}

func (c *Client) search(vals url.Values) ([]*NZB, error) {
	vals.Set("apikey", c.apikey)
	return c.process(vals, apiPath)
}

func (c *Client) caps(vals url.Values) (Capabilities, error) {
	vals.Set("apikey", c.apikey)
	resp, err := c.getURL(c.buildURL(vals, apiPath))
	if err != nil {
		return Capabilities{}, errors.Wrap(err, "failed to get capabilities")
	}
	var cResp Capabilities
	if err = xml.Unmarshal(resp, &cResp); err != nil {
		return cResp, errors.Wrap(err, "failed to unmarshal xml")
	}
	return cResp, nil
}

func (c *Client) details(vals url.Values) (Details, error) {
	vals.Set("apikey", c.apikey)
	resp, err := c.getURL(c.buildURL(vals, apiPath))
	if err != nil {
		return Details{}, errors.Wrap(err, "failed to get details")
	}
	var dResp Details
	if err = xml.Unmarshal(resp, &dResp); err != nil {
		return dResp, errors.Wrap(err, "failed to unmarshal xml")
	}
	return dResp, nil
}

func (c *Client) process(vals url.Values, path string) ([]*NZB, error) {
	var nzbs []*NZB

	resp, err := c.getURL(c.buildURL(vals, path))
	if err != nil {
		return nzbs, err
	}

	var feed SearchResponse
	err = xml.Unmarshal(resp, &feed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal xml feed")
	}
	if feed.ErrorCode != 0 {
		return nil, errors.Errorf("newznab api error %d: %s", feed.ErrorCode, feed.ErrorDesc)
	}
	for _, gotNZB := range feed.Channel.NZBs {
		nzb := &NZB{
			Title:          gotNZB.Title,
			Description:    gotNZB.Description,
			PubDate:        gotNZB.Date.Add(0),
			DownloadURL:    gotNZB.Enclosure.URL,
			SourceEndpoint: c.apiBaseURL,
			SourceAPIKey:   c.apikey,
		}
		// see: https://inhies.github.io/Newznab-API/attributes/
		for _, attr := range gotNZB.Attributes {
			switch attr.Name {
			case "size":
				v, _ := parseSize(attr.Value)
				nzb.Size = v
			case "category":
				nzb.Category = append(nzb.Category, attr.Value)
			case "guid":
				nzb.ID = attr.Value
			case "poster":
				nzb.Poster = attr.Value
			case "group":
				nzb.Group = attr.Value
			case "team":
				nzb.Team = attr.Value
			case "grabs":
				parsedInt, _ := strconv.ParseInt(attr.Value, 0, 32)
				nzb.NumGrabs = int(parsedInt)
			case "password":
				nzb.Password = attr.Value
			case "comments":
				parsedInt, _ := strconv.ParseInt(attr.Value, 0, 32)
				nzb.NumComments = int(parsedInt)
			case "usenetdate":
				if parsedUsetnetDate, err := parseDate(attr.Value); err != nil {
					log.WithError(err).WithField("usenetdate", attr.Value).Debug("failed to parse usenetdate")
				} else {
					nzb.UsenetDate = parsedUsetnetDate
				}
			case "info":
				nzb.Info = attr.Value
			case "year":
				parsedInt, _ := strconv.ParseInt(attr.Value, 0, 32)
				nzb.Year = int(parsedInt)
			case "season":
				nzb.Season = attr.Value
			case "episode":
				nzb.Episode = attr.Value
			case "tvairdate":
				if parsedAirDate, err := parseDate(attr.Value); err != nil {
					log.WithError(err).WithField("tvairdate", attr.Value).Debug("newznab:Client:Search: failed to parse tvairdate")
				} else {
					nzb.AirDate = parsedAirDate
				}
			case "seeders":
				parsedInt, _ := strconv.ParseInt(attr.Value, 0, 32)
				nzb.Seeders = int(parsedInt)
				nzb.IsTorrent = true
			case "peers":
				parsedInt, _ := strconv.ParseInt(attr.Value, 0, 32)
				nzb.Peers = int(parsedInt)
				nzb.IsTorrent = true
			case "infohash":
				nzb.InfoHash = attr.Value
				nzb.IsTorrent = true
			case "magneturl":
				nzb.DownloadURL = attr.Value
			case "genre":
				nzb.Genre = attr.Value
			case "tvdbid":
				nzb.TVDBID = attr.Value
			case "rageid":
				nzb.TVRageID = attr.Value
			case "tvmazeid":
				nzb.TVMazeID = attr.Value
			case "tvtitle":
				nzb.TVTitle = attr.Value
			case "rating":
				parsedInt, _ := strconv.ParseInt(attr.Value, 0, 32)
				nzb.Rating = int(parsedInt)
			case "coverurl":
				nzb.CoverURL = attr.Value
			case "resolution":
				nzb.Resolution = attr.Value
			case "video":
				nzb.Video = attr.Value
			case "audio":
				nzb.Audio = attr.Value
			case "framerate":
				nzb.Framerate = attr.Value
			case "language":
				nzb.Language = attr.Value
			case "subs":
				nzb.Subs = strings.Split(attr.Value, ",")
			case "imdb":
				nzb.IMDBID = attr.Value
			case "imdbtitle":
				nzb.IMDBTitle = attr.Value
			case "imdbyear":
				parsedInt, _ := strconv.ParseInt(attr.Value, 0, 32)
				nzb.IMDBYear = int(parsedInt)
			case "imdbscore":
				parsedFloat, _ := strconv.ParseFloat(attr.Value, 32)
				nzb.IMDBScore = float32(parsedFloat)
			case "imdbtagline":
				nzb.IMDBTagline = attr.Value
			case "imdbplot":
				nzb.IMDBPlot = attr.Value
			case "imdbdirector":
				nzb.IMDBDirector = attr.Value
			case "imdbactors":
				nzb.IMDBActors = strings.Split(attr.Value, ",")
			case "artist":
				nzb.Artist = attr.Value
			case "album":
				nzb.Album = attr.Value
			case "publisher":
				nzb.Publisher = attr.Value
			case "tracks":
				nzb.Tracks = attr.Value
			case "backdropcoverurl":
				nzb.BackdropCoverURL = attr.Value
			case "review":
				nzb.Review = attr.Value
			case "booktitle":
				nzb.BookTitle = attr.Value
			case "publishdate":
				if d, err := parseDate(attr.Value); err != nil {
					log.WithError(err).WithField("tvairdate", attr.Value).Debug("newznab:Client:Search: failed to parse tvairdate")
				} else {
					nzb.PublishDate = d
				}
			case "author":
				nzb.Author = attr.Value
			case "pages":
				parsedInt, _ := strconv.ParseInt(attr.Value, 0, 32)
				nzb.Pages = int(parsedInt)
			default:
				log.WithFields(log.Fields{
					"name":  attr.Name,
					"value": attr.Value,
				}).Debug("encontered unknown attribute")
			}
		}
		if nzb.Size == 0 {
			parsedInt, _ := parseSize(gotNZB.Size)
			nzb.Size = parsedInt
		}
		if len(nzb.Category) > 0 {
			nzb.Category = lo.Uniq(nzb.Category)
		}
		nzbs = append(nzbs, nzb)
	}
	return nzbs, nil
}

// PopulateComments fills in the Comments for the given NZB
func (c *Client) PopulateComments(nzb *NZB) error {
	data, err := c.getURL(c.buildURL(url.Values{
		"t":      []string{"comments"},
		"id":     []string{nzb.ID},
		"apikey": []string{c.apikey},
	}, apiPath))
	if err != nil {
		return err
	}
	var resp commentResponse
	err = xml.Unmarshal(data, &resp)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal comments xml data")
	}

	for _, rawComment := range resp.Channel.Comments {
		comment := Comment{
			Title:   rawComment.Title,
			Content: rawComment.Description,
		}
		if parsedPubDate, err := time.Parse(time.RFC1123Z, rawComment.PubDate); err != nil {
			log.WithError(err).WithField("pubdate", rawComment.PubDate).Debug("failed to parse comment date")
		} else {
			comment.PubDate = parsedPubDate
		}
		nzb.Comments = append(nzb.Comments, comment)
	}
	return nil
}

// NZBDownloadURL returns a URL to download the NZB from
func (c *Client) NZBDownloadURL(nzb *NZB) (string, error) {
	return c.buildURL(url.Values{
		"t":      []string{"get"},
		"id":     []string{nzb.ID},
		"apikey": []string{c.apikey},
	}, apiPath)
}

// DownloadNZB returns the bytes of the actual NZB file for the given NZB
func (c *Client) DownloadNZB(nzb *NZB) ([]byte, error) {
	return c.getURL(c.NZBDownloadURL(nzb))
}

func (c *Client) getURL(url string, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	res, err := c.client.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "http request failed: %s", url)
	}

	var data []byte
	data, err = io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}
	return data, nil
}

func (c *Client) buildURL(vals url.Values, path string) (string, error) {
	parsedURL, err := url.Parse(c.apiBaseURL + path)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse base API url")
	}

	parsedURL.RawQuery = vals.Encode()
	return parsedURL.String(), nil
}

func parseDate(date string) (time.Time, error) {
	formats := []string{time.RFC3339, time.RFC1123Z}
	var parsedTime time.Time
	var err error
	for _, format := range formats {
		if parsedTime, err = time.Parse(format, date); err == nil {
			return parsedTime, nil
		}
	}
	return parsedTime, errors.Errorf("failed to parse date %s as one of %s", date, strings.Join(formats, ", "))
}

func parseSize(size string) (int64, error) {
	if matches := sizeRegex.FindStringSubmatch(size); len(matches) > 0 {
		size, _ := strconv.ParseFloat(matches[1], 64)
		switch matches[2] {
		case "KiB":
			size *= 1024
		case "MiB":
			size *= 1024 * 1024
		case "GiB":
			size *= 1024 * 1024 * 1024
		}
		return int64(size), nil
	}

	return strconv.ParseInt(size, 0, 64)
}

type commentResponse struct {
	Channel struct {
		Comments []rssComment `xml:"item"`
	} `xml:"channel"`
}

type rssComment struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
