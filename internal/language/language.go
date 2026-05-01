package language

import (
	"encoding/json"
	"os"
)

type Translations map[string]string

var languages = map[string]Translations{}

func LoadLanguages() error {
	files := map[string]string{
		"en": "internal/locales/en.json",
		"de": "internal/locales/de.json",
	}

	for lang, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var t Translations
		err = json.Unmarshal(data, &t)
		if err != nil {
			return err
		}

		languages[lang] = t
	}

	return nil
}

func T(lang, key string) string {
	if l, ok := languages[lang]; ok {
		if v, ok := l[key]; ok {
			return v
		}
	}

	return key
}
