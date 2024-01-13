package runic

import "time"

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`

	Categories []string `json:"categories"`
	Size       int64    `json:"size"`
	Guid       string   `json:"guid"`

	View     string `json:"view"`
	Download string `json:"download"`

	Published time.Time `json:"published"`
}

/*
   &gofeed.Item{
     Title:       "Self.Reliance.2023.720p.WEBRip-LAMA",
     Description: "<div><table style=\"width:100%;\"><tr><td width=\"100%\"><ul><li>ID: fb3b3a15b1e68d66819d923cf01218e9</li><li>Name: <a href=\"https://nzbgeek.info/geekseek.php?guid=fb3b3a15b1e68d66819d923cf01218e9\">Self.Reliance.2023.720p.WEBRip-LAMA</a></li><li>Size: 921.14 MB </li><li>Attributes: Category: <a href=\"https://nzbgeek.info/geekseek.php?c=2040\">Movies > HD</a></li><li>PostDate: Fri, 12 Jan 2024 06:19:51 +0000</li><li>Imdb Info<ul><li>IMDB Link: <a href=\"http://www.imdb.com/title/tt26084002\">Self Reliance</a></li><li>Rating: <span style=\"font-weight:bold; color:#2AFF00;\">7.4</span></li><li>Plot: Given the opportunity to participate in a life or death reality game show, one man discovers theres a lot to live for</li><li>Year: 2023</li><li>Genre: Comedy, Thriller</li><li>Director: Jake Johnson</li><li>Actors: Jake Johnson, Anna Kendrick, Andy Samberg</li><li>Runtime: 85 min</li></ul></li><br><font size=\"2\" face=\"Verdana\" color=\"#999999\">1</font>&nbsp;<img height=\"15\" src=\"https://api.nzbgeek.info/covers/grabs.png\">&nbsp;&nbsp;<font size=\"2\" face=\"Verdana\" color=\"#0000FF\">0</font>&nbsp;<img height=\"15\" src=\"https://api.nzbgeek.info/covers/comments.png\">&nbsp;&nbsp;<font size=\"2\" face=\"Verdana\" color=\"#008000\">0</font>&nbsp;<img height=\"15\" src=\"https://api.nzbgeek.info/covers/thumbup.png\">&nbsp;&nbsp;<font size=\"2\" face=\"Verdana\" color=\"#FF0000\">0</font>&nbsp;<img height=\"15\" src=\"https://api.nzbgeek.info/covers/thumbdown.png\">&nbsp;&nbsp;</ul></td><td width=\"120px\" align=\"right\" valign=\"top\"><img style=\"margin-left:10px;margin-bottom:10px;float:right;\" src=\"https://api.nzbgeek.info/covers/movies/26084002-cover.jpg\" width=\"120\" border=\"0\" alt=\"Self Reliance\" /></td></tr></table></div><div style=\"clear:both;\">",
     Content:     "",
     Link:        "https://nzbgeek.info/geekseek.php?guid=fb3b3a15b1e68d66819d923cf01218e9",
     Links:       []string{
       "https://nzbgeek.info/geekseek.php?guid=fb3b3a15b1e68d66819d923cf01218e9",
     },
     Updated:         "",
     UpdatedParsed:   (*time.Time)(nil),
     Published:       "Fri, 12 Jan 2024 06:19:51 +0000",
     PublishedParsed: &2024-01-12 06:19:51 UTC,
     Author:          (*gofeed.Person)(nil),
     Authors:         []*gofeed.Person{},
     GUID:            "https://nzbgeek.info/geekseek.php?guid=fb3b3a15b1e68d66819d923cf01218e9",
     Image:           (*gofeed.Image)(nil),
     Categories:      []string{
       "Movies > HD",
     },
     Enclosures: []*gofeed.Enclosure{
       &gofeed.Enclosure{
         URL:    "https://api.nzbgeek.info/api?t=get&id=fb3b3a15b1e68d66819d923cf01218e9&apikey=eISG7JzxXnmWskK632mjY3CHRylfVuiX",
         Length: "965890000",
         Type:   "application/x-nzb",
       },
     },
     DublinCoreExt: (*ext.DublinCoreExtension)(nil),
     ITunesExt:     (*ext.ITunesItemExtension)(nil),
     Extensions:    ext.Extensions{
       "newznab": map[string][]ext.Extension{
         "attr": []ext.Extension{
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "category",
               "value": "2000",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "category",
               "value": "2040",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "size",
               "value": "965890000",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "guid",
               "value": "fb3b3a15b1e68d66819d923cf01218e9",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "imdbtitle",
               "value": "Self Reliance",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "imdb",
               "value": "26084002",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "imdbtagline",
               "value": "Surviving is all about the company you keep",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "imdbplot",
               "value": "Given the opportunity to participate in a life or death reality game show, one man discovers theres a lot to live for",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "imdbscore",
               "value": "7.4",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "genre",
               "value": "Comedy, Thriller",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "imdbyear",
               "value": "2023",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "imdbdirector",
               "value": "Jake Johnson",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "imdbactors",
               "value": "Jake Johnson, Anna Kendrick, Andy Samberg",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "coverurl",
               "value": "https://api.nzbgeek.info/covers/movies/26084002-cover.jpg",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "runtime",
               "value": "85 min",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "grabs",
               "value": "1",
             },
             Children: map[string][]ext.Extension{},
           },
           ext.Extension{
             Name:  "attr",
             Value: "",
             Attrs: map[string]string{
               "name":  "usenetdate",
               "value": "Fri, 12 Jan 2024 06:19:51 +0000",
             },
             Children: map[string][]ext.Extension{},
           },
         },
       },
     },
     Custom: map[string]string{},
   },
*/
