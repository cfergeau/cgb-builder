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
