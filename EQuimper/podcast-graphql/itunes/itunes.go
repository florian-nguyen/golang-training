package itunes

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

// ItunesApiServices struct definition
type ItunesApiServices struct {
}

// ItunesApiServices constructor
func NewItunesApiServices() *ItunesApiServices {
	return &ItunesApiServices{}
}

// Source : https://itunes.apple.com/search?entity=podcast&term=syntax
// Source : https://mholt.github.io/json-to-go/
type SearchResponse struct {
	Resultcount int `json:"resultCount"`
	Results     []struct {
		Wrappertype            string    `json:"wrapperType"`
		Kind                   string    `json:"kind"`
		Collectionid           int       `json:"collectionId"`
		Trackid                int       `json:"trackId"`
		Artistname             string    `json:"artistName"`
		Collectionname         string    `json:"collectionName"`
		Trackname              string    `json:"trackName"`
		Collectioncensoredname string    `json:"collectionCensoredName"`
		Trackcensoredname      string    `json:"trackCensoredName"`
		Collectionviewurl      string    `json:"collectionViewUrl"`
		Feedurl                string    `json:"feedUrl,omitempty"`
		Trackviewurl           string    `json:"trackViewUrl"`
		Artworkurl30           string    `json:"artworkUrl30"`
		Artworkurl60           string    `json:"artworkUrl60"`
		Artworkurl100          string    `json:"artworkUrl100"`
		Collectionprice        float64   `json:"collectionPrice"`
		Trackprice             float64   `json:"trackPrice"`
		Trackrentalprice       int       `json:"trackRentalPrice"`
		Collectionhdprice      int       `json:"collectionHdPrice"`
		Trackhdprice           int       `json:"trackHdPrice"`
		Trackhdrentalprice     int       `json:"trackHdRentalPrice"`
		Releasedate            time.Time `json:"releaseDate"`
		Collectionexplicitness string    `json:"collectionExplicitness"`
		Trackexplicitness      string    `json:"trackExplicitness"`
		Trackcount             int       `json:"trackCount"`
		Country                string    `json:"country"`
		Currency               string    `json:"currency"`
		Primarygenrename       string    `json:"primaryGenreName"`
		Contentadvisoryrating  string    `json:"contentAdvisoryRating,omitempty"`
		Artworkurl600          string    `json:"artworkUrl600"`
		Genreids               []string  `json:"genreIds"`
		Genres                 []string  `json:"genres"`
		Artistid               int       `json:"artistId,omitempty"`
		Artistviewurl          string    `json:"artistViewUrl,omitempty"`
	} `json:"results"`
}

// Warning: URL encoding needed to deal with spaces that might be included in the search term.
// Response is then obtained based on the encoded URL.
func (ias *ItunesApiServices) Search(term string) (SearchResponse, error) {
	searchUrl := url.URL{
		Scheme: "https",
		Host:   "itunes.apple.com",
		Path:   "search",
	}

	q := url.Values{}
	q.Set("entity", "podcast")
	q.Set("term", term)

	searchUrl.RawQuery = q.Encode()

	res, err := http.Get(searchUrl.String())
	if err != nil {
		return SearchResponse{}, err
	}

	defer res.Body.Close()

	var searchResponse SearchResponse
	err = json.NewDecoder(res.Body).Decode(&searchResponse)

	return searchResponse, err
}
