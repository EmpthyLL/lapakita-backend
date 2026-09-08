package locale

var UserMessages = map[string]map[string]string{
	"en": {
		// General Profile
		"user.profile_get_success":      "Successfully retrieved user profile",
		"user.profile_update_success":   "Successfully updated general profile",
		"user.profile_failed_to_get":    "Failed to retrieve user profile",
		"user.profile_failed_to_update": "Failed to update general profile",

		// Phone Numbers
		"user.phone_get_success":           "Successfully retrieved phone numbers",
		"user.phone_add_success":           "Successfully added phone number",
		"user.phone_update_success":        "Successfully updated phone number",
		"user.phone_delete_success":        "Successfully deleted phone number",
		"user.phone_duplicate":             "This phone number is already registered to your account",
		"user.phone_index_invalid":         "Invalid phone number index",
		"user.phone_cannot_delete_primary": "Primary phone number cannot be deleted",
		"user.phone_failed_to_add":         "Failed to add phone number",
		"user.phone_failed_to_update":      "Failed to update phone number",
		"user.phone_failed_to_delete":      "Failed to delete phone number",

		// Password
		"user.password_change_success":   "Successfully changed password",
		"user.password_incorrect":        "Current password is incorrect",
		"user.password_failed_to_change": "Failed to change password",

		// Persona Profile
		"user.persona_get_success":      "Successfully retrieved persona profile",
		"user.persona_update_success":   "Successfully updated persona profile",
		"user.persona_not_found":        "Persona profile not found for this role",
		"user.persona_role_invalid":     "Invalid persona role",
		"user.persona_failed_to_get":    "Failed to retrieve persona profile",
		"user.persona_failed_to_update": "Failed to update persona profile",

		// Documents & Watermark
		"user.document_upload_success":   "Successfully uploaded identity document",
		"user.document_delete_success":   "Successfully deleted identity document",
		"user.document_nik_exists":       "This NIK is already registered to another account",
		"user.document_file_required":    "Identity document photo is required",
		"user.document_file_invalid":     "Invalid identity document file format",
		"user.document_watermark_failed": "Failed to apply watermark to document",
		"user.document_failed_to_upload": "Failed to upload identity document",
		"user.document_failed_to_delete": "Failed to delete identity document",
	},
	"id": {
		// General Profile
		"user.profile_get_success":      "Berhasil mengambil profil pengguna",
		"user.profile_update_success":   "Berhasil memperbarui profil umum",
		"user.profile_failed_to_get":    "Gagal mengambil profil pengguna",
		"user.profile_failed_to_update": "Gagal memperbarui profil umum",

		// Phone Numbers
		"user.phone_get_success":           "Berhasil mengambil daftar nomor telepon",
		"user.phone_add_success":           "Berhasil menambahkan nomor telepon",
		"user.phone_update_success":        "Berhasil memperbarui nomor telepon",
		"user.phone_delete_success":        "Berhasil menghapus nomor telepon",
		"user.phone_duplicate":             "Nomor telepon ini sudah terdaftar di akun Anda",
		"user.phone_index_invalid":         "Index nomor telepon tidak valid",
		"user.phone_cannot_delete_primary": "Nomor telepon utama tidak dapat dihapus",
		"user.phone_failed_to_add":         "Gagal menambahkan nomor telepon",
		"user.phone_failed_to_update":      "Gagal memperbarui nomor telepon",
		"user.phone_failed_to_delete":      "Gagal menghapus nomor telepon",

		// Password
		"user.password_change_success":   "Berhasil mengubah kata sandi",
		"user.password_incorrect":        "Kata sandi saat ini tidak cocok",
		"user.password_failed_to_change": "Gagal mengubah kata sandi",

		// Persona Profile
		"user.persona_get_success":      "Berhasil mengambil profil persona",
		"user.persona_update_success":   "Berhasil memperbarui profil persona",
		"user.persona_not_found":        "Profil persona tidak ditemukan untuk role ini",
		"user.persona_role_invalid":     "Role persona tidak valid",
		"user.persona_failed_to_get":    "Gagal mengambil profil persona",
		"user.persona_failed_to_update": "Gagal memperbarui profil persona",

		// Documents & Watermark
		"user.document_upload_success":   "Berhasil mengunggah dokumen identitas",
		"user.document_delete_success":   "Berhasil menghapus dokumen identitas",
		"user.document_nik_exists":       "NIK KTP ini telah terdaftar pada akun lain",
		"user.document_file_required":    "Foto KTP wajib diunggah",
		"user.document_file_invalid":     "Format file dokumen identitas tidak valid",
		"user.document_watermark_failed": "Gagal menambahkan watermark pada dokumen",
		"user.document_failed_to_upload": "Gagal mengunggah dokumen identitas",
		"user.document_failed_to_delete": "Gagal menghapus dokumen identitas",
	},
}
