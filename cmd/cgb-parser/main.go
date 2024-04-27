package main

import (
        "encoding/json"
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

func dump(node *gohtml.Node, indent string) {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
                log.Infof("%s%s(%d) %v", indent, c.Data, c.Type, c.Attr)
                dump(c, indent + "  ")
        }
}

func isAsset(cardText *gohtml.Node) bool {
	assetMatcher := func(node *gohtml.Node) bool {
                return node.Data == "img" && html.HasAttrWithValue(node, "src", "https://haa-src.cgbuilder.fr/images/carte_cout.png")
	}
        return html.FindNode(cardText, assetMatcher) != nil
}

func isTraits(cardText *gohtml.Node) bool {
        // `<p style="font-style: italic; font-weight: bold;">`
	traitsMatcher := func(node *gohtml.Node) bool {
                return node.Data == "p" && html.HasAttrWithValue(node, "style", "font-style: italic; font-weight: bold;")
	}
        return html.FindNode(cardText, traitsMatcher) != nil
}

func isMetadataFooter(cardText *gohtml.Node) bool {
        // cycle/card metadata: `<p><i><a href="https://haa.cgbuilder.fr/liste_carte/26/">Pour le Bien Commun</a>, 203</i></p>`
	footerMatcher := func(node *gohtml.Node) bool {
                if node.Data != "a" {
                        return false
                }
                href := html.FindAttr(node, "href")
                return strings.HasPrefix(href, "https://haa.cgbuilder.fr/liste_carte/26/")
	}
        return html.FindNode(cardText, footerMatcher) != nil
}

func isLocation(cardText *gohtml.Node) bool {
        // <p><span>4<img src="https://haa-src.cgbuilder.fr/images/carte_occulte.png"/></span><span>1<img src="https://haa-src.cgbuilder.fr/images/carte_indice.png"/></span></p>
        shroudMatcher  := func(node *gohtml.Node) bool {
                return node.Data == "img" && html.HasAttrWithValue(node, "src", "https://haa-src.cgbuilder.fr/images/carte_occulte.png")
	}
        clueMatcher  := func(node *gohtml.Node) bool {
                return node.Data == "img" && html.HasAttrWithValue(node, "src", "https://haa-src.cgbuilder.fr/images/carte_indice.png")
	}
        return html.FindNode(cardText, shroudMatcher) != nil || html.FindNode(cardText, clueMatcher) != nil
}

func isQuest(cardText *gohtml.Node) bool {
        // no difference between act and agenda :-/
        // front:
        // `<p style="margin-top: 16px;"><span class="face_quete">`
        // back:
        // `<p><span class="face_quete">`
        questMatcher := func(node *gohtml.Node) bool {
                return node.Data == "span" && html.HasClass(node, "face_quete")
        }
        return html.FindNode(cardText, questMatcher) != nil
}



func isQuestFront(cardText *gohtml.Node) bool {
        if !isQuest(cardText) {
                return false
        }

        return cardText.Data == "p" && html.HasAttrWithValue(cardText, "style", "margin-top: 16px;")
}

func isQuestBack(cardText *gohtml.Node) bool {
        if !isQuest(cardText) {
                return false
        }

        return cardText.Data == "p" && cardText.Attr == nil
}

func isText(cardText *gohtml.Node) bool {
	return cardText.Data == "p" && html.HasClass(cardText, "zone_texte")
}

func parseCardText(node *gohtml.Node, card *arkhamdb.Card) error {
	replacer := text.NewReplacer()
        if isAsset(node) {
                card.TypeCode = arkhamdb.Asset
        }
        var str string
        var isBack bool
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != gohtml.ElementNode {
			return fmt.Errorf("unexpected HTML node type: %d", c.Type)
		}
		if c.Data != "p" && c.Data != "ul" && c.Data != "br" {
			return fmt.Errorf("unexpected HTML node: %s", c.Data)
		}
		if isText(c) {
                        for n := c.FirstChild; n != nil; n = n.NextSibling {
                                txt, err := html.DumpNode(n)
                                if err != nil {
			                return fmt.Errorf("failed to dump node: %w", err)
                                }
                                str += txt
                        }
		}
		if str == "<p></p>" {
			continue
		}
		if c.Data == "br" {
                        str += "\n"
			continue
		}


                if isMetadataFooter(c) {
                        if isBack {
                                card.BackText = gohtml.UnescapeString(replacer.Replace(str))
                        } else {
                                card.Text = gohtml.UnescapeString(replacer.Replace(str))
                        }
			continue
		}
                if isLocation(c) {
                        card.TypeCode = arkhamdb.Location
                        continue
                }
                if isQuestFront(c) {
                        card.TypeCode = arkhamdb.Act
                        continue
                }
                if isQuestBack(c) {
                        isBack = true
                        card.Text = gohtml.UnescapeString(replacer.Replace(str))
                        str = ""
                        continue
                }
                /*
		if strings.HasPrefix(str, `<p style="margin-top: 16px;"><span class="face_quete">`) {
			// next item is 'card.Text', card type is Act or Agenda
		}
		if strings.HasPrefix(str, `<p><span class="face_quete">`) {
			// next item is 'card.BackText', card type is Act or Agenda
		}
                */
                /* Asset:
                 * - hasTag <img src="https://haa-src.cgbuilder.fr/images/carte_cout.png"/>
                 * - traits
                 * - text
                 * - card metadata

		/*
			<p><span>4<img src="https://haa-src.cgbuilder.fr/images/carte_occulte.png"/></span><span>1<img src="https://haa-src.cgbuilder.fr/images/carte_indice.png"/></span></p>
			<p><span>3<img src="https://haa-src.cgbuilder.fr/images/carte_lutte.png"/></span><span>3<img src="https://haa-src.cgbuilder.fr/images/carte_point_vie.png"/></span><span>4<img src="https://haa-src.cgbuilder.fr/images/carte_agilite.png"/></span><b>|</b> <span><img src="https://haa-src.cgbuilder.fr/images/carte_point_vie.png"/></span><span><img src="https://haa-src.cgbuilder.fr/images/carte_sante_mentale.png"/></span></p>
			<p><span>1<img src="https://haa-src.cgbuilder.fr/images/carte_cout.png"/></span><span>3 <i class="fas fa-layer-group"></i></span><span><img src="https://haa-src.cgbuilder.fr/images/carte_intelligence.png"/><img src="https://haa-src.cgbuilder.fr/images/carte_intelligence.png"/></span></p>
		*/
                 if isTraits(c) { 
                         card.Traits = c.FirstChild.Data
			//log.Infof("traits: %s", c.FirstChild.Data)
			continue
		}
		//log.Infof("card text: %s", gohtml.UnescapeString(replacer.Replace(str)))
	}
        //dump(node, "")

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

func parseInfoBulle(node *gohtml.Node) (*arkhamdb.Card, error) {
	infoBulleCount += 1
	var cycleId, cardId int
	infoBulleId := html.GetId(node)
	if _, err := fmt.Sscanf(infoBulleId, "info_bulle_%d_%d", &cycleId, &cardId); err != nil {
		log.Infof("Failed to parse string '%s': %v", infoBulleId, err)
		return nil, err
	}
	log.Infof("parsing card %d for cycle %d", cardId, cycleId)
	// we are parsing cycle 26, which is part of tcu in arkhamdb, which has ID 5
	//const tcuCycleId = 5
	card := &arkhamdb.Card{
		Code:     fmt.Sprintf("%02d%03d", arkhamdb.TheCircleUndone, cardId),
		PackCode: arkhamdb.TheCircleUndone,
	}
	log.Infof("arkhamdb url: %s", card.URL())

	name, err := parseCardTitle(node)
	if err != nil {
		return nil, err
	}
	card.Name = name

	cardTextMatcher := func(node *gohtml.Node) bool {
		return html.HasClass(node, "texte") && html.HasClass(node, "texte_carte")
	}
	cardTextNode := html.FindNode(node, cardTextMatcher)
        if err := parseCardText(cardTextNode, card); err != nil {
                return nil, err
        }

        return card, nil
}

func parse(doc *gohtml.Node) error {
        var cards []*arkhamdb.Card
	infoBulleMatcher := func(n *gohtml.Node) bool {
		return strings.HasPrefix(html.GetId(n), "info_bulle_")
	}
	infoBulles := html.FindNodes(doc, infoBulleMatcher)
	for _, infoBulle := range infoBulles {
		card, err := parseInfoBulle(infoBulle)
                if err != nil {
			return err
		}
		log.Infof("====")
                cards = append(cards, card)
	}
	log.Infof("Parsed %d info bulles", infoBulleCount)

        strBuilder := strings.Builder{}
        enc := json.NewEncoder(&strBuilder)
        enc.SetEscapeHTML(false)
        enc.SetIndent("", "    ")
        if err := enc.Encode(cards); err != nil {
                return err
        }

        log.Info(strBuilder.String())
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
