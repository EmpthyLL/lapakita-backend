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

	// CMS Feature Keys
	KeyCMSFAQGetSuccess    MessageKey = "cms.faq.get_success"
	KeyCMSFAQNotFound      MessageKey = "cms.faq.not_found"
	KeyCMSFAQFailedToGet   MessageKey = "cms.faq.failed_to_get"
	KeyCMSLegalGetSuccess  MessageKey = "cms.legal.get_success"
	KeyCMSLegalNotFound    MessageKey = "cms.legal.not_found"
	KeyCMSLegalFailedToGet MessageKey = "cms.legal.failed_to_get"

	// Business Type Feature Keys
	KeyBusinessTypeGetSuccess  MessageKey = "business_type.get_success"
	KeyBusinessTypeFailedToGet MessageKey = "business_type.failed_to_get"
)
