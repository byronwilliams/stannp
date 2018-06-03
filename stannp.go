package stannp

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NewClient(apiKey string, httpClient *http.Client) *StannpAPI {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &StannpAPI{
		apiKey: apiKey,
		client: httpClient,
	}
}

func (st *StannpAPI) PostMergedLetterURL(country, pdfURL string, isTest bool) (*http.Response, error) {
	const endpoint = "https://dash.stannp.com/api/v1/letters/post"
	u, _ := url.Parse(endpoint)

	var vals = url.Values{}
	vals.Set("test", strconv.FormatBool(isTest))
	vals.Set("country", country)
	vals.Set("pdf", pdfURL)

	var bf = &bytes.Buffer{}
	var mw = multipart.NewWriter(bf)
	mw.WriteField("test", strconv.FormatBool(isTest))
	mw.WriteField("country", country)
	mw.WriteField("pdf", pdfURL)
	mw.Close()
	// log.Println(bf.String())
	u.RawQuery = vals.Encode()
	log.Println(u.String())
	req, _ := http.NewRequest("POST", u.String(), strings.NewReader(vals.Encode()))
	req.SetBasicAuth(st.apiKey, "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))

	return st.client.Do(req)
	// return st.PostMergedLetter(country, strings.NewReader(vals.Encode()), isTest)
}

// func (st *StannpAPI) PostMergedLetter(country, pdf io.Reader, isTest bool) {
// 	const endpoint = "https://dash.stannp.com/api/v1/letters/post"
// 	u, _ := url.Parse(endpoint)

// 	var vals = url.Values{}
// 	vals.Set("test", string(isTest))
// 	vals.Set("country", country)
// 	vals.Set("pdf", pdf)

// 	var bf = &bytes.Buffer{}
// 	var mw = multipart.NewWriter(bf)
// 	mw.

// 	return st.PostMergedLetter(country, strings.NewReader(vals.Encode()), isTest)
// }

// return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
