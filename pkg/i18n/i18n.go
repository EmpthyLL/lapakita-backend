package i18n

import (
	"lapakita-backend/pkg/i18n/locale"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	dictionary = make(map[string]map[MessageKey]string)
	once       sync.Once
)

// RegisterMessages menggabungkan pesan terjemahan dari modul ke dictionary utama
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

// GetLang mengambil bahasa dari header
func GetLang(c *gin.Context) string {
	lang := c.GetHeader("lang")
	if lang == "" {
		lang = c.GetHeader("Accept-Language")
	}
	if lang != "id" && lang != "en" {
		lang = "en"
	}
	return lang
}

// T mengambil pesan terjemahan secara Type-Safe berdasarkan MessageKey
func T(c *gin.Context, key MessageKey) string {
	lang := GetLang(c)
	if val, exists := dictionary[lang][key]; exists {
		return val
	}
	// Fallback ke bahasa Inggris jika key tidak ada di bahasa target
	if val, exists := dictionary["en"][key]; exists {
		return val
	}
	return string(key)
}

// Init mendaftarkan seluruh dictionary bahasa dari setiap modul/kategori
func Init() {
	RegisterMessages(locale.GeneralMessages)
	RegisterMessages(locale.AreaMessages)
}
