package text

import "strings"

type Substitution struct {
	haaBuilderText string
	arkhamDBText   string
}

var substitutions []Substitution = []Substitution{
	{`<br/>`, `\n`},
	{`<img src="https://haa-src.cgbuilder.fr/images/carte_pion_cultiste.png" alt="pion_cultiste"/>`, `[cultist]`},
	{`<img src="https://haa-src.cgbuilder.fr/images/carte_pion_tentacule.png" alt="pion_tentacule"/>`, `[elder_thing]`},
	{`<img src="https://haa-src.cgbuilder.fr/images/carte_pion_crane.png" alt="pion_crane"/>`, `[skull]`},
	{`<img src="https://haa-src.cgbuilder.fr/images/carte_pion_vitrail.png" alt="pion_vitrail"/>`, `[tablet]`},

	{`<img src="https://haa-src.cgbuilder.fr/images/carte_volonte.png" alt="volonte"/>`, `[willpower]`},
	{`<img src="https://haa-src.cgbuilder.fr/images/carte_intelligence.png" alt="intelligence"/>`, `[intellect]`},
	{`<img src="https://haa-src.cgbuilder.fr/images/carte_agilite.png" alt="agilite"/>`, `[agility]`},
	{`<img src="https://haa-src.cgbuilder.fr/images/carte_lutte.png" alt="lutte"/>`, `[combat]`},

	{`<img src="https://haa-src.cgbuilder.fr/images/carte_action.png" alt="action"/>`, `[action]`},
	{`<img src="https://haa-src.cgbuilder.fr/images/carte_reaction.png" alt="reaction"/>`, `[reaction]`},
	{`<img src="https://haa-src.cgbuilder.fr/images/carte_rapide.png" alt="rapide"/>`, `[free]`},
	{`<img src="https://haa-src.cgbuilder.fr/images/carte_indice.png" alt="indice"/>`, `[per_investigator]`},
	/* `<b><i>Cultiste</i></b>`, `[[Cultiste]]` */
}

func NewReplacer() *strings.Replacer {
	replacements := []string{}
	for _, s := range substitutions {
		replacements = append(replacements, s.haaBuilderText, s.arkhamDBText)
	}

	return strings.NewReplacer(replacements...)
}
