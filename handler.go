package urlshort

import (
	"log"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if redirectUrl, ok := pathsToUrls[path]; ok {
			log.Println("Redirecting to", redirectUrl)

			// Let's set a cache control header
			// Normally 301 (Moved Permanently) are cached by browsers
			// Could use 302 (Found) as alternative to setting cache control
			w.Header().Set("Cache-Control", "private, max-age=30")
			http.Redirect(w, r, redirectUrl, 301)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func RedirectHandler(filename string, fallback http.Handler) (http.HandlerFunc, error) {
	mapper := Mapper{filename}

	// Decode
	urlMapping, err := mapper.decode()
	if err != nil {
		return nil, err
	}

	return MapHandler(urlMapping, fallback), nil
}

