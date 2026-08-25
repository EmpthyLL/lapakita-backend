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

	// Generic & Common
	KeyInvalidPayload MessageKey = "error.invalid_payload"
	KeyUserNotFound   MessageKey = "error.user_not_found"

	// Registration & Login
	KeyEmailAlreadyRegistered MessageKey = "error.email_already_registered"
	KeyWrongPassword          MessageKey = "error.wrong_password"
	KeyRegisterSuccess        MessageKey = "success.register"
	KeyLoginSuccess           MessageKey = "success.login"

	// Google Auth & Complete Profile
	KeyGoogleAuthFailed            MessageKey = "error.google_auth_failed"
	KeySetupTokenExpired           MessageKey = "error.setup_token_expired"
	KeyGoogleAuthProfileIncomplete MessageKey = "success.google_auth_profile_incomplete"
	KeyProfileCompleteSuccess      MessageKey = "success.profile_completed"

	// OTP System
	KeyOTPSendFailed    MessageKey = "error.otp_send_failed"
	KeyOTPSendSuccess   MessageKey = "success.otp_sent"
	KeyOTPVerifySuccess MessageKey = "success.otp_verified"
	KeyOTPExpired       MessageKey = "error.otp_expired"

	// Password Reset & Tokens
	KeyResetTokenExpired    MessageKey = "error.reset_token_expired"
	KeyResetPasswordSuccess MessageKey = "success.password_reset"
	KeyRefreshTokenExpired  MessageKey = "error.refresh_token_expired"
	KeyTokenRefreshSuccess  MessageKey = "success.token_refreshed"
)
