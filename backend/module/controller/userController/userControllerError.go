package userController

import "backend/library/common"

var UserChangePasswordError map[string]common.ErrorMessageInterface = map[string]common.ErrorMessageInterface{
	"Password": {
		Field: "password",
		Messages: map[string]string{
			"required": "Password sebelumnya wajib diisi.",
		},
	},
	"PasswordConfirmation": {
		Field: "password_confirmation",
		Messages: map[string]string{
			"required": "Konfirmasi Password Baru wajib diisi.",
		},
	},
	"PasswordNew": {
		Field: "password_new",
		Messages: map[string]string{
			"required":      "Anda belum mengisi password baru.",
			"min":           "Password baru minimal 10 karakter.",
			"confirmed":     "Password baru dan Konfirmasi Password Baru tidak sama.",
			"has_uppercase": "Password baru harus mengandung minimal 1 huruf kapital.",
			"has_lowercase": "Password baru harus mengandung minimal 1 huruf non-kapital.",
			"has_number":    "Password baru harus mengandung minimal 1 nomor.",
			"has_symbol":    "Password baru harus mengandung minimal 1 simbol selain angka dan huruf.",
		},
	},
}

var UserUpdateError map[string]common.ErrorMessageInterface = map[string]common.ErrorMessageInterface{
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
			"required":          "Anda belum mengisi username.",
			"max":               "Username maksimal 100 karakter.",
			"is_unique":         "Username sudah dipakai",
			"alphanumeric_dash": "Username hanya boleh menggunakan alfabet, nomor, strip, dan underscore",
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
}
