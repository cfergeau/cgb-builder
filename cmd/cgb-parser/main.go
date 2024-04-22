package main

import (
	"fmt"
	"strings"

	"github.com/cfergeau/cgb-parser/pkg/arkhamdb"
	"github.com/cfergeau/cgb-parser/pkg/html"
	"github.com/cfergeau/cgb-parser/pkg/text"
	log "github.com/sirupsen/logrus"
	gohtml "golang.org/x/net/html"
)

const tcuId = 26

var tcuURL = fmt.Sprintf("https://haa.cgbuilder.fr/liste_carte/%d/", tcuId)

func parseCardText(node *gohtml.Node) error {
	replacer := text.NewReplacer()
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != gohtml.ElementNode {
			return fmt.Errorf("unexpected HTML node type: %d", c.Type)
		}
		if c.Data != "p" && c.Data != "ul" && c.Data != "br" {
			return fmt.Errorf("unexpected HTML node: %s", c.Data)
		}
		var str string
		var err error
		if c.Data == "p" && html.HasClass(c, "zone_texte") {
			//str, err = dumpNode(c.FirstChild)
			str, err = html.DumpNode(c)
		} else {
			str, err = html.DumpNode(c)
		}
		if err != nil {
			return fmt.Errorf("failed to dump node: %w", err)
		}
		if str == "<p></p>" {
			continue
		}
		if c.Data == "br" {
			continue
		}
		if strings.HasPrefix(str, fmt.Sprintf(`<p><i><a href="https://haa.cgbuilder.fr/liste_carte/%d/">`, 26)) {
			// cycle/card metadata: `<p><i><a href="https://haa.cgbuilder.fr/liste_carte/26/">Pour le Bien Commun</a>, 203</i></p>`
			continue
		}
		if strings.HasPrefix(str, `<p style="margin-top: 16px;"><span class="face_quete">`) {
			// next item is 'card.Text', card type is Act or Agenda
		}
		if strings.HasPrefix(str, `<p><span class="face_quete">`) {
			// next item is 'card.BackText', card type is Act or Agenda
		}

		/*
			<p><span>4<img src="https://haa-src.cgbuilder.fr/images/carte_occulte.png"/></span><span>1<img src="https://haa-src.cgbuilder.fr/images/carte_indice.png"/></span></p>
			<p><span>3<img src="https://haa-src.cgbuilder.fr/images/carte_lutte.png"/></span><span>3<img src="https://haa-src.cgbuilder.fr/images/carte_point_vie.png"/></span><span>4<img src="https://haa-src.cgbuilder.fr/images/carte_agilite.png"/></span><b>|</b> <span><img src="https://haa-src.cgbuilder.fr/images/carte_point_vie.png"/></span><span><img src="https://haa-src.cgbuilder.fr/images/carte_sante_mentale.png"/></span></p>
			<p><span>1<img src="https://haa-src.cgbuilder.fr/images/carte_cout.png"/></span><span>3 <i class="fas fa-layer-group"></i></span><span><img src="https://haa-src.cgbuilder.fr/images/carte_intelligence.png"/><img src="https://haa-src.cgbuilder.fr/images/carte_intelligence.png"/></span></p>
		*/
		if strings.HasPrefix(str, `<p style="font-style: italic; font-weight: bold;">`) {
			log.Infof("traits: %s", c.FirstChild.Data)
			continue
		}
		log.Infof("card text: %s", gohtml.UnescapeString(replacer.Replace(str)))
	}

	return nil
}

var infoBulleCount = 0

func parseCardTitle(infoBulle *gohtml.Node) (string, error) {
	replacer := text.NewReplacer()
	titleMatcher := func(node *gohtml.Node) bool {
		hasBackgroundClass := false
		hasBorderClass := false
		classes := html.GetClasses(node)
		for _, cl := range classes {
			if strings.HasPrefix(cl, "background_") {
				hasBackgroundClass = true
			}
			if strings.HasPrefix(cl, "border_") {
				hasBorderClass = true
			}
		}
		return hasBackgroundClass && hasBorderClass
	}

	titleNode := html.FindNode(infoBulle, titleMatcher)
	for c := titleNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "b" && c.FirstChild == c.LastChild {
			titleNode = c.FirstChild
		}
	}

	str, err := html.DumpNode(titleNode)
	if err != nil {
		return "", err
	}
	title := gohtml.UnescapeString(replacer.Replace(str))

	return title, nil
}

func parseInfoBulle(node *gohtml.Node) error {
	infoBulleCount += 1
	var cycleId, cardId int
	infoBulleId := html.GetId(node)
	if _, err := fmt.Sscanf(infoBulleId, "info_bulle_%d_%d", &cycleId, &cardId); err != nil {
		log.Infof("Failed to parse string '%s': %v", infoBulleId, err)
		return err
	}
	log.Infof("parsing card %d for cycle %d", cardId, cycleId)
	// we are parsing cycle 26, which is part of tcu in arkhamdb, which has ID 5
	//const tcuCycleId = 5
	card := arkhamdb.Card{
		Code:     fmt.Sprintf("%02d%03d", arkhamdb.TheCircleUndone, cardId),
		PackCode: arkhamdb.TheCircleUndone,
	}
	log.Infof("arkhamdb url: %s", card.URL())

	name, err := parseCardTitle(node)
	if err != nil {
		return err
	}
	card.Name = name
	log.Infof("Title: %s", card.Name)

	cardTextMatcher := func(node *gohtml.Node) bool {
		return html.HasClass(node, "texte") && html.HasClass(node, "texte_carte")
	}
	cardTextNode := html.FindNode(node, cardTextMatcher)
	return parseCardText(cardTextNode)
}

func parse(doc *gohtml.Node) error {
	infoBulleMatcher := func(n *gohtml.Node) bool {
		return strings.HasPrefix(html.GetId(n), "info_bulle_")
	}
	infoBulles := html.FindNodes(doc, infoBulleMatcher)
	for _, infoBulle := range infoBulles {
		if err := parseInfoBulle(infoBulle); err != nil {
			return err
		}
		log.Infof("====")
	}
	log.Infof("Parsed %d info bulles", infoBulleCount)

	return nil
}

func main() {
	log.Info("CGB-Parser")
	log.Infof("Fetching %s", tcuURL)
	htmlBody, err := html.FetchURL(tcuURL)
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

	doc, err := gohtml.Parse(htmlBody)
	if err != nil {
		log.Fatalf("Failed to parse HTML document: %v", err)
	}
	if err := parse(doc); err != nil {
		log.Fatalf("Failed to parse HTML data: %v", err)
	}
}
