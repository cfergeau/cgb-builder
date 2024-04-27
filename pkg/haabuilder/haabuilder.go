package haabuilder

import (
        "fmt"
        "path/filepath"
)

const baseURL = "https://haa.cgbuilder.fr/liste_carte/"

func (pack *Pack) URL() string {
        return baseURL+pack.HaaBuilderCode+"/"
}

func (pack *Pack) Path() string {
        return filepath.Join("pack", pack.CycleCode, fmt.Sprintf("%s.json", pack.Code))
}

func (pack *Pack) I18nPath(lang string) string {
        return filepath.Join("translations", lang, pack.Path())
}

func (pack *Pack) EncountersPath() string {
        return filepath.Join("pack", pack.CycleCode, fmt.Sprintf("%s_encounter.json", pack.Code))
}

func (pack *Pack) I18nEncountersPath(lang string) string {
        return filepath.Join("translations", lang, pack.EncountersPath())
}
