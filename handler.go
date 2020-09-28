package urlshort

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
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

func YAMLHandler(filename string, fallback http.Handler) (http.HandlerFunc, error) {
	// Read the YAML mapping
	mapPath, err := readYAML(filename)
	if err != nil {
		return nil, err
	}

	// Return handler
	return MapHandler(mapPath, fallback), nil
}

func readYAML(filename string) (map[string]string, error) {
	// Open yaml file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error in opening file: ", err)
	}
	defer file.Close()

	// Create a yaml decoder from *File (File is a Reader since it implements Read)
	decoder := yaml.NewDecoder(file)

	// Decode
	var decoded []map[string]string
	err = decoder.Decode(&decoded)
	if err != nil {
		return nil, err
	}

	// Process to be a key-value pair / map
	result := make(map[string]string)
	for _, mapping := range decoded {
		result[mapping["path"]] = mapping["url"]
	}

	return result, nil
}
