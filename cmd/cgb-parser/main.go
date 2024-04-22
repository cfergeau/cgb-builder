package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	log "github.com/sirupsen/logrus"
)

const tcuId = 23
const userAgent = "cgb-parser/0.0.1"

var tcuURL = fmt.Sprintf("https://haa.cgbuilder.fr/liste_carte/%d/", tcuId)

func fetchURL(url string) (string, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if false {
		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatalf("Error dumping HAABuilder response: %v", err)
		}
		log.Infof("HAABuilder data: %s", string(b))
	}

	return string(body), nil
}

func main() {
	log.Info("CGB-Parser")
	log.Infof("Fetching %s", tcuURL)
	htmlData, err := fetchURL(tcuURL)
	if err != nil {
		log.Fatalf("Failed to fetch HAABuilder data from %s: %v", tcuURL, err)
	}
	log.Infof("HAABuilder data: %s", htmlData)

}
