package usecase

// ResolveLanguage chooses a valid request language, falling back to the
// language supplied by the caller and then to English.
func ResolveLanguage(requested, fallback string) string {
	if requested != "" {
		return requested
	}
	if fallback != "" {
		return fallback
	}
	return "en"
}
