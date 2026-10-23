package locale

var MasterMessages = map[string]map[string]string{
	"en": {
		"area.search_success":              "Area search successful",
		"area.search_failed":               "Failed to search area",
		"area.history_get_success":         "Search history retrieved successfully",
		"area.history_save_success":        "Search history saved successfully",
		"area.history_delete_success":      "Search history deleted successfully",
		"area.device_id_required":          "Device ID header (X-Device-ID) is required for guest history",
		"area.history_delete_item_success": "Search history item deleted successfully",

		"public.faq.get_success":   "Successfully retrieved FAQs",
		"public.faq.not_found":     "FAQs not found",
		"public.faq.failed_to_get": "Failed to retrieve FAQs data",

		"public.legal.get_success":   "Successfully retrieved legal document",
		"public.legal.not_found":     "Legal document not found",
		"public.legal.failed_to_get": "Failed to retrieve legal document",

		"public.contact_inquiry.submit_success": "Thank you! Your inquiry has been submitted successfully.",
		"public.contact_inquiry.submit_failed":  "Failed to submit your inquiry. Please try again later.",

		"business_type.get_success":   "Successfully retrieved business types",
		"business_type.failed_to_get": "Failed to retrieve business types data",
	},
	"id": {
		"area.search_success":              "Pencarian area berhasil",
		"area.search_failed":               "Gagal mencari area",
		"area.history_get_success":         "Riwayat pencarian berhasil diambil",
		"area.history_save_success":        "Riwayat pencarian berhasil disimpan",
		"area.history_delete_success":      "Riwayat pencarian berhasil dihapus",
		"area.device_id_required":          "Header Device ID (X-Device-ID) diperlukan untuk riwayat guest",
		"area.history_delete_item_success": "Item riwayat pencarian berhasil dihapus",

		"public.faq.get_success":   "Berhasil mengambil data FAQ",
		"public.faq.not_found":     "Pertanyaan umum (FAQ) tidak ditemukan",
		"public.faq.failed_to_get": "Gagal mengambil data FAQ",

		"public.legal.get_success":   "Berhasil mengambil dokumen hukum",
		"public.legal.not_found":     "Dokumen hukum tidak ditemukan",
		"public.legal.failed_to_get": "Gagal mengambil dokumen hukum",

		"public.contact_inquiry.submit_success": "Terima kasih! Pesan Anda telah berhasil terkirim.",
		"public.contact_inquiry.submit_failed":  "Gagal mengirim pesan Anda. Silakan coba lagi nanti.",

		"business_type.get_success":   "Berhasil mengambil data jenis bisnis",
		"business_type.failed_to_get": "Gagal mengambil data jenis bisnis",
	},
}
