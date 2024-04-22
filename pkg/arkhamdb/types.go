package arkhamdb

import "fmt"

type PackCode int

const (
	Core PackCode = iota + 1
	TheDunwichLegacy
	ThePathToCarcosa
	TheForgottenAge
	TheCircleUndone
	TheDreamEaters
	TheInssmouthConspiracy
	EdgeOfEarth
	TheScarletKeys
	TheFeastOfHemlockVale
	ReturnTo                 = 50
	InvestigatorStarterDecks = 60
	SideStories              = 70
	Parallel                 = 90
)

type FactionCode string

const (
	Guardian FactionCode = "guardian"
	Mystic               = "mystic"
	Mythos               = "mythos"
	Neutral              = "neutral"
	Rogue                = "rogue"
	Seeker               = "seeker"
	Survivor             = "survivor"
)

type TypeCode string

const (
	Act          TypeCode = "act"
	Agenda                = "agenda"
	Asset                 = "asset"
	Enemy                 = "enemy"
	Event                 = "event"
	Investigator          = "investigator"
	Key                   = "key"
	Location              = "location"
	Scenario              = "scenario"
	Skill                 = "skill"
	Story                 = "story"
	Treachery             = "treachery"
)

type SubTypeCode string

const (
	BasicWeakness SubTypeCode = "basicweakness"
	Weakness                  = "weakness"
)

type Card struct {
	Code        string      `json:"code"`         // "01011"
	FactionCode FactionCode `json:"faction_code"` // "neutral"
	PackCode    PackCode    `json:"pack_code"`    // "core"
	TypeCode    TypeCode    `json:"type_code"`    // "treachery"
	SubtypeCode SubTypeCode `json:"subtype_code"`

	// Card Text
	Name       string `json:"name"`
	Subname    string `json:"subname"`
	BackFlavor string `json:"back_flavor"`
	BackText   string `json:"back_text"`
	Flavor     string `json:"flavor"`
	Text       string `json:"text"`
	Traits     string `json:"traits"`

	// Card Info
	DoubleSided bool   `json:"double_sided"`
	IsUnique    bool   `json:"is_unique"`
	Position    int    `json:"position"`
	Quantity    int    `json:"quantity"`
	Illustrator string `json:"illustrator"`

	// Encounters
	EncounterCode     string `json:"encounter_code"`
	EncounterPosition int    `json:"encounter_position"`

	// Deck Options
	// DeckLimit int
	// DeckOptions
	// DeckRequirements

	// Investigator characteristics
	// Health         int
	// Sanity         int
	// SkillAgility   int
	// SkillIntellect int
	// SkillCombat    int
	// SkillWillPower int
}

var cycleShortStr = map[PackCode]string{
	Core:                     "core",
	TheDunwichLegacy:         "dwl",
	ThePathToCarcosa:         "ptc",
	TheForgottenAge:          "tfa",
	TheCircleUndone:          "tcu",
	TheDreamEaters:           "tde",
	TheInssmouthConspiracy:   "tic",
	EdgeOfEarth:              "eoe",
	TheScarletKeys:           "tsk",
	TheFeastOfHemlockVale:    "fhv",
	ReturnTo:                 "return",
	InvestigatorStarterDecks: "investigator",
	SideStories:              "side_stories",
	Parallel:                 "parallel",
}

func (id PackCode) String() string {
	shortStr, ok := cycleShortStr[id]
	if !ok {
		shortStr = fmt.Sprintf("unknown(%d)", id)
	}

	return shortStr
}

func (card *Card) URL() string {
	return fmt.Sprintf("https://arkhamdb.com/card/%s", card.Code)
}
