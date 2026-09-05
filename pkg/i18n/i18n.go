package i18n

import (
	"sync"

	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n/locale"

	"github.com/gin-gonic/gin"
)

var (
	dictionary = make(map[string]map[MessageKey]string)
	once       sync.Once
)

func RegisterMessages(moduleMessages map[string]map[string]string) {
	for lang, keys := range moduleMessages {
		if _, exists := dictionary[lang]; !exists {
			dictionary[lang] = make(map[MessageKey]string)
		}
		for key, value := range keys {
			dictionary[lang][MessageKey(key)] = value
		}
	}
}

func TByLang(lang string, key MessageKey) string {
	if lang != "id" && lang != "en" {
		lang = "en"
	}
	if val, exists := dictionary[lang][key]; exists {
		return val
	}
	if val, exists := dictionary["en"][key]; exists {
		return val
	}
	return string(key)
}

func T(c *gin.Context, key MessageKey) string {
	lang := api.GetLanguageFromHeader(c)
	return TByLang(lang, key)
}

func Init() {
	RegisterMessages(locale.GeneralMessages)
	RegisterMessages(locale.AreaMessages)
	RegisterMessages(locale.PublicMessages)
	RegisterMessages(locale.AuthMessages)
	RegisterMessages(locale.StallMessages)
}
