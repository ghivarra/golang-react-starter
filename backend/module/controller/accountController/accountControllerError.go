package accountController

import "backend/library/common"

var AccountUpdateError map[string]common.ErrorMessageInterface = map[string]common.ErrorMessageInterface{
	"ID": {
		Field: "id",
		Messages: map[string]string{
			"required":      "ID pengguna harus dikirim",
			"is_not_unique": "User tidak ditemukan",
		},
	},
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
	"RoleID": {
		Field: "role_id",
		Messages: map[string]string{
			"required":      "Anda belum mengisi role user.",
			"is_not_unique": "Role tidak ditemukan.",
		},
	},
}

var AccountStatusError map[string]common.ErrorMessageInterface = map[string]common.ErrorMessageInterface{
	"ID": {
		Field: "id",
		Messages: map[string]string{
			"required":      "Identitas akun wajib diisi",
			"is_not_unique": "Akun tidak ditemukan",
		},
	},
	"Status": {
		Field: "status",
		Messages: map[string]string{
			"required": "Target status akun wajib diisi",
			"in_list":  "status hanya boleh salah satu antara activate dan deactivate",
		},
	},
}
