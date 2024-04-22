package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"slices"
	"strings"

	"github.com/cfergeau/cgb-parser/pkg/text"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

const tcuId = 26
const userAgent = "cgb-parser/0.0.1"

var tcuURL = fmt.Sprintf("https://haa.cgbuilder.fr/liste_carte/%d/", tcuId)

func fetchURL(url string) (io.ReadCloser, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if false {
		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatalf("Error dumping HAABuilder response: %v", err)
		}
		log.Infof("HAABuilder data: %s", string(b))
	}

	return resp.Body, nil
}

func findAttr(node *html.Node, attrName string) string {
	for _, a := range node.Attr {
		if a.Key == attrName {
			return a.Val
		}
	}

	return ""
}

func getClasses(node *html.Node) []string {
	return strings.Split(findAttr(node, "class"), " ")
}

func hasClass(node *html.Node, class string) bool {
	classes := getClasses(node)
	return slices.Contains(classes, class)
}

func getId(node *html.Node) string {
	return findAttr(node, "id")
}

func findNodes(root *html.Node, match func(*html.Node) bool) []*html.Node {
	matches := []*html.Node{}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		if match(c) {
			matches = append(matches, c)
		}
		childMatches := findNodes(c, match)
		if childMatches != nil {
			matches = append(matches, childMatches...)
		}
	}

	if len(matches) == 0 {
		return nil
	}
	return matches
}

func findNode(root *html.Node, match func(*html.Node) bool) *html.Node {
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		if match(c) {
			return c
		}
		if n := findNode(c, match); n != nil {
			return n
		}
	}

	return nil
}

func dumpNode(node *html.Node) (string, error) {
	strBuilder := &strings.Builder{}
	if err := html.Render(strBuilder, node); err != nil {
		return "", err
	}
	return strBuilder.String(), nil
}

func parseCardText(node *html.Node) error {
	replacer := text.NewReplacer()
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			return fmt.Errorf("unexpected HTML node type: %d", c.Type)
		}
		if c.Data == "br" {
			continue
		}
		if c.Data != "p" && c.Data != "ul" {
			return fmt.Errorf("unexpected HTML node: %s", c.Data)
		}
		var str string
		var err error
		if c.Data == "p" && hasClass(c, "zone_texte") {
			//str, err = dumpNode(c.FirstChild)
			str, err = dumpNode(c)
		} else {
			str, err = dumpNode(c)
		}
		if err != nil {
			return fmt.Errorf("failed to dump node: %w", err)
		}
		if str == "<p></p>" {
			continue
		}
		if strings.HasPrefix(str, fmt.Sprintf(`<p><i><a href="https://haa.cgbuilder.fr/liste_carte/%d/">`, 26)) {
			continue
		}
		log.Infof("card text: %s", html.UnescapeString(replacer.Replace(str)))
	}

	return nil
}

var infoBulleCount = 0

func parseInfoBulle(node *html.Node) error {
	infoBulleCount += 1
	var cycleId, cardId int
	infoBulleId := getId(node)
	if _, err := fmt.Sscanf(infoBulleId, "info_bulle_%d_%d", &cycleId, &cardId); err != nil {
		log.Infof("Failed to parse string '%s': %v", infoBulleId, err)
		return err
	}
	log.Infof("parsing card %d for cycle %d", cardId, cycleId)
	// we are parsing cycle 26, which is part of tcu in arkhamdb, which has ID 5
	const tcuCycleId = 5
	log.Infof("arkhamdb url: https://arkhamdb.com/card/%02d%03d", tcuCycleId, cardId)
	nodeMatcher := func(node *html.Node) bool {
		return hasClass(node, "texte") && hasClass(node, "texte_carte")
	}
	cardTextNode := findNode(node, nodeMatcher)
	return parseCardText(cardTextNode)
}

func parse(doc *html.Node) error {
	infoBulleMatcher := func(n *html.Node) bool {
		return strings.HasPrefix(getId(n), "info_bulle_")
	}
	infoBulles := findNodes(doc, infoBulleMatcher)
	for _, infoBulle := range infoBulles {
		if err := parseInfoBulle(infoBulle); err != nil {
			return err
		}
	}
	log.Infof("Parsed %d info bulles", infoBulleCount)

	return nil
}

func main() {
	log.Info("CGB-Parser")
	log.Infof("Fetching %s", tcuURL)
	htmlBody, err := fetchURL(tcuURL)
	if err != nil {
		log.Fatalf("Failed to fetch HAABuilder data from %s: %v", tcuURL, err)
	}
	defer htmlBody.Close()
	/*
	   htmlData, err := io.ReadAll(htmlBody)
	   if err != nil {
	           log.Fatalf("Failed to read HTML body from %s: %v", tcuURL, err)
	   }
	   log.Infof("HAABuilder data: %s", htmlData)
	*/

	doc, err := html.Parse(htmlBody)
	if err != nil {
		log.Fatalf("Failed to parse HTML document: %v", err)
	}
	if err := parse(doc); err != nil {
		log.Fatalf("Failed to parse HTML data: %v", err)
	}
}
