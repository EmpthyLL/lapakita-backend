package locale

var GeneralMessages = map[string]map[string]string{
	"en": {
		"general.query_invalid":            "Query parameter invalid",
		"general.auth_header_not_found":    "Authorization header not found",
		"general.token_format_invalid":     "Token format must be 'Bearer <token>'",
		"general.token_invalid_or_expired": "Invalid or expired token",
		"general.token_claims_failed":      "Failed to process token claim data",
		"general.rate_limit_exceeded":      "Too many requests. Please slow down.",
		"general.internal_server_error":    "Internal server error",
		"general.unauthorized":             "Unauthorized access",
	},
	"id": {
		"general.query_invalid":            "Parameter query tidak valid",
		"general.auth_header_not_found":    "Header otorisasi tidak ditemukan",
		"general.token_format_invalid":     "Format token harus 'Bearer <token>'",
		"general.token_invalid_or_expired": "Token tidak valid atau sudah kadaluwarsa",
		"general.token_claims_failed":      "Gagal memproses data klaim token",
		"general.rate_limit_exceeded":      "Terlalu banyak permintaan. Silakan pelan-pelan.",
		"general.internal_server_error":    "Terjadi kesalahan pada server",
		"general.unauthorized":             "Akses tidak diizinkan",
	},
}
