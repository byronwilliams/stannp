package stannp

import "net/http"

type StannpAPI struct {
	client *http.Client
	apiKey string
}

type MergedLetterOpts struct {

}