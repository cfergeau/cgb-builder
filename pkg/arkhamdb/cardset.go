package arkhamdb

import (
	"encoding/json"
	"os"
	"strings"
)

type CardSet struct {
	cards   []*Card
	cardMap map[string]*Card
}

func NewEmpty() *CardSet {
	cs := CardSet{}
	cs.cardMap = map[string]*Card{}

	return &cs
}

func NewFromFile(filename string) (*CardSet, error) {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cards []*Card
	if err := json.Unmarshal(jsonData, &cards); err != nil {
		return nil, err
	}

	cardSet := NewEmpty()
	for _, c := range cards {
		cardSet.AddCard(c)
	}

	return cardSet, nil
}

func (cs *CardSet) AddCard(card *Card) {
	cs.cards = append(cs.cards, card)
	cs.cardMap[card.Code] = card
}

func (cs *CardSet) MarshalIndent() (string, error) {
	// sort cs.cards by `Code`

	strBuilder := strings.Builder{}
	enc := json.NewEncoder(&strBuilder)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "    ")
	if err := enc.Encode(cs.cards); err != nil {
		return "", err
	}

	return strBuilder.String(), nil
}

func (cs *CardSet) WriteFile(filename string, perm os.FileMode) error {
        jsonStr, err := cs.MarshalIndent()
        if err != nil {
                return err
        }
        if err := os.WriteFile(filename, []byte(jsonStr), perm); err != nil {
                return err
        }

        return nil
}

func (cs *CardSet) MergeCardSetText(newCardSet *CardSet) {
	for _, newCard := range newCardSet.cards {
		card, hasCard := cs.cardMap[newCard.Code]
		if !hasCard {
			continue
		}
		if newCard.BackFlavor != "" {
			card.BackFlavor = newCard.BackFlavor
		}
		if newCard.BackName != "" {
			card.BackName = newCard.BackName
		}
		if newCard.BackText != "" {
			card.BackText = newCard.BackText
		}
		if newCard.Flavor != "" {
			card.Flavor = newCard.Flavor
		}
		if newCard.Name != "" {
			card.Name = newCard.Name
		}
		if newCard.Slot != "" {
			card.Slot = newCard.Slot
		}
		if newCard.SubName != "" {
			card.SubName = newCard.SubName
		}
		if newCard.Text != "" {
			card.Text = newCard.Text
		}
		if newCard.Traits != "" {
			card.Traits = newCard.Traits
		}
	}
}
