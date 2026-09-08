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
	KeyUnauthorized          MessageKey = "general.unauthorized"

	// Area Feature Keys
	KeyAreaSearchSuccess            MessageKey = "area.search_success"
	KeyAreaSearchFailed             MessageKey = "area.search_failed"
	KeyAreaHistoryGetSuccess        MessageKey = "area.history_get_success"
	KeyAreaHistorySaveSuccess       MessageKey = "area.history_save_success"
	KeyAreaHistoryDeleteSuccess     MessageKey = "area.history_delete_success"
	KeyAreaDeviceIDRequired         MessageKey = "area.device_id_required"
	KeyAreaHistoryDeleteItemSuccess MessageKey = "area.history_delete_item_success"

	// Public Feature Keys
	KeyPublicFAQGetSuccess               MessageKey = "public.faq.get_success"
	KeyPublicFAQNotFound                 MessageKey = "public.faq.not_found"
	KeyPublicFAQFailedToGet              MessageKey = "public.faq.failed_to_get"
	KeyPublicLegalGetSuccess             MessageKey = "public.legal.get_success"
	KeyPublicLegalNotFound               MessageKey = "public.legal.not_found"
	KeyPublicLegalFailedToGet            MessageKey = "public.legal.failed_to_get"
	KeyPublicContactInquirySubmitSuccess MessageKey = "public.contact_inquiry.submit_success"
	KeyPublicContactInquirySubmitFailed  MessageKey = "public.contact_inquiry.submit_failed"

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

	// Stall Keys
	KeyStallCreateSuccess     MessageKey = "stall.create_success"
	KeyStallUpdateSuccess     MessageKey = "stall.update_success"
	KeyStallDeleteSuccess     MessageKey = "stall.delete_success"
	KeyStallGetSuccess        MessageKey = "stall.get_success"
	KeyStallSearchSuccess     MessageKey = "stall.search_success"
	KeyStallNotFound          MessageKey = "stall.not_found"
	KeyStallUnauthorized      MessageKey = "stall.unauthorized"
	KeyStallFailedToCreate    MessageKey = "stall.failed_to_create"
	KeyStallFailedToUpdate    MessageKey = "stall.failed_to_update"
	KeyStallFailedToDelete    MessageKey = "stall.failed_to_delete"
	KeyStallFailedToGet       MessageKey = "stall.failed_to_get"
	KeyStallGetSimilarSuccess MessageKey = "stall.get_similar_success"

	// User Keys
	KeyUserProfileGetSuccess     MessageKey = "user.profile_get_success"
	KeyUserProfileUpdateSuccess  MessageKey = "user.profile_update_success"
	KeyUserProfileFailedToGet    MessageKey = "user.profile_failed_to_get"
	KeyUserProfileFailedToUpdate MessageKey = "user.profile_failed_to_update"

	KeyUserPhoneGetSuccess          MessageKey = "user.phone_get_success"
	KeyUserPhoneAddSuccess          MessageKey = "user.phone_add_success"
	KeyUserPhoneUpdateSuccess       MessageKey = "user.phone_update_success"
	KeyUserPhoneDeleteSuccess       MessageKey = "user.phone_delete_success"
	KeyUserPhoneDuplicate           MessageKey = "user.phone_duplicate"
	KeyUserPhoneIndexInvalid        MessageKey = "user.phone_index_invalid"
	KeyUserPhoneCannotDeletePrimary MessageKey = "user.phone_cannot_delete_primary"
	KeyUserPhoneFailedToAdd         MessageKey = "user.phone_failed_to_add"
	KeyUserPhoneFailedToUpdate      MessageKey = "user.phone_failed_to_update"
	KeyUserPhoneFailedToDelete      MessageKey = "user.phone_failed_to_delete"

	KeyUserPasswordChangeSuccess  MessageKey = "user.password_change_success"
	KeyUserPasswordIncorrect      MessageKey = "user.password_incorrect"
	KeyUserPasswordFailedToChange MessageKey = "user.password_failed_to_change"

	KeyUserPersonaGetSuccess     MessageKey = "user.persona_get_success"
	KeyUserPersonaUpdateSuccess  MessageKey = "user.persona_update_success"
	KeyUserPersonaNotFound       MessageKey = "user.persona_not_found"
	KeyUserPersonaRoleInvalid    MessageKey = "user.persona_role_invalid"
	KeyUserPersonaFailedToGet    MessageKey = "user.persona_failed_to_get"
	KeyUserPersonaFailedToUpdate MessageKey = "user.persona_failed_to_update"

	KeyUserDocumentUploadSuccess   MessageKey = "user.document_upload_success"
	KeyUserDocumentDeleteSuccess   MessageKey = "user.document_delete_success"
	KeyUserDocumentNIKExists       MessageKey = "user.document_nik_exists"
	KeyUserDocumentFileRequired    MessageKey = "user.document_file_required"
	KeyUserDocumentFileInvalid     MessageKey = "user.document_file_invalid"
	KeyUserDocumentWatermarkFailed MessageKey = "user.document_watermark_failed"
	KeyUserDocumentFailedToUpload  MessageKey = "user.document_failed_to_upload"
	KeyUserDocumentFailedToDelete  MessageKey = "user.document_failed_to_delete"
)
