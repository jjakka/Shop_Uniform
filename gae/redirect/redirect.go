package redirect

import (
    "net/http"
    "net/url"
)

func init() {
    http.HandleFunc("/", redirect)
}

const dest string = "http://shopuniform.com"

func redirect(w http.ResponseWriter, r *http.Request) {
	respURL, _ := url.Parse(dest)
    http.Redirect(w, r, respURL.String(), 301)
}