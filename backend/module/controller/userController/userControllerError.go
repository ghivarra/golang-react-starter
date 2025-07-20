package userController

import "backend/library/common"

var RegisterError map[string]common.ErrorMessageInterface = map[string]common.ErrorMessageInterface{
	"Name": {
		Field: "name",
		Messages: map[string]string{
			"required": "Anda belum mengisi nama.",
			"max":      "Nama maksimal 100 karakter.",
		},
	},
	"Username": {
		Field: "username",
		Messages: map[string]string{
			"required":  "Anda belum mengisi username.",
			"max":       "Username maksimal 100 karakter.",
			"is_unique": "Username sudah dipakai",
		},
	},
	"Email": {
		Field: "email",
		Messages: map[string]string{
			"required":  "Anda belum mengisi email.",
			"max":       "Email maksimal 100 karakter.",
			"email":     "Email tidak valid.",
			"is_unique": "Email sudah dipakai",
		},
	},
	"RoleID": {
		Field: "role_id",
		Messages: map[string]string{
			"required":      "Anda belum mengisi role user.",
			"is_not_unique": "Role tidak ditemukan.",
		},
	},
	"Password": {
		Field: "password",
		Messages: map[string]string{
			"required":      "Anda belum mengisi password.",
			"min":           "Password minimal 10 karakter.",
			"confirmed":     "Password dan Konfirmasi Password tidak sama.",
			"has_uppercase": "Password harus mengandung minimal 1 huruf kapital.",
			"has_lowercase": "Password harus mengandung minimal 1 huruf non-kapital.",
			"has_number":    "Password harus mengandung minimal 1 nomor.",
			"has_symbol":    "Password harus mengandung minimal 1 simbol selain angka dan huruf.",
		},
	},
	"PasswordConfirmation": {
		Field: "password_confirmation",
		Messages: map[string]string{
			"required": "Anda belum mengisi konfirmasi password.",
		},
	},
}
