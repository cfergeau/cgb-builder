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
	Code        string      `json:"code"`                   // "01011"
	FactionCode FactionCode `json:"faction_code,omitempty"` // "neutral"
	PackCode    PackCode    `json:"pack_code,omitempty"`    // "core"
	TypeCode    TypeCode    `json:"type_code,omitempty"`    // "treachery"
	SubtypeCode SubTypeCode `json:"subtype_code,omitempty"`

	// Card Text
        // Structs keys are marshalled in the order defined in the struct
	Flavor     string `json:"flavor,omitempty"`
	Name       string `json:"name,omitempty"`
	SubName    string `json:"subname,omitempty"`
	Text       string `json:"text,omitempty"`
	BackName   string `json:"backname,omitempty"`
	Traits     string `json:"traits,omitempty"`
	BackFlavor string `json:"back_flavor,omitempty"`
	BackText   string `json:"back_text,omitempty"`
	Slot       string `json:"slot,omitempty"`

	// Card Info
	DoubleSided bool   `json:"double_sided,omitempty"`
	IsUnique    bool   `json:"is_unique,omitempty"`
	Position    int    `json:"position,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Illustrator string `json:"illustrator,omitempty"`

	// Encounters
	EncounterCode     string `json:"encounter_code,omitempty"`
	EncounterPosition int    `json:"encounter_position,omitempty"`

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
