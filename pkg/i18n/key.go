package i18n

type MessageKey string

const (
	// General / System Keys
	KeyQueryInvalid          MessageKey = "general.query_invalid"
	KeyAuthHeaderNotFound    MessageKey = "general.auth_header_not_found"
	KeyTokenFormatInvalid    MessageKey = "general.token_format_invalid"
	KeyTokenInvalidOrExpired MessageKey = "general.token_invalid_or_expired"
	KeyTokenClaimsFailed     MessageKey = "general.token_claims_failed"
	KeyRateLimitExceeded     MessageKey = "general.rate_limit_exceeded"
	KeyInternalServerError   MessageKey = "general.internal_server_error"

	// Area Feature Keys
	KeyAreaSearchSuccess MessageKey = "area.search_success"
	KeyAreaSearchFailed  MessageKey = "area.search_failed"

	// User Feature Keys (Contoh untuk masa depan)
	KeyUserNotFound MessageKey = "user.not_found"
)
