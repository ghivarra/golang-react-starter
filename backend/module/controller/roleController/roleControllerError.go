package roleController

import "backend/library/common"

var RoleCreateError map[string]common.ErrorMessageInterface = map[string]common.ErrorMessageInterface{
	"Name": {
		Field: "name",
		Messages: map[string]string{
			"required":  "Nama role wajib diisi",
			"max":       "Nama maksimal 60 karakter",
			"is_unique": "Nama role sudah dipakai sebelumnya",
		},
	},
	"IsSuperadmin": {
		Field: "is_superadmin",
		Messages: map[string]string{
			"in_list": "Hanya boleh menggunakan value 'Super Admin' atau 'Bukan Super Admin'",
		},
	},
}

var RoleUpdateError map[string]common.ErrorMessageInterface = map[string]common.ErrorMessageInterface{
	"ID": {
		Field: "id",
		Messages: map[string]string{
			"required":      "Identitas role wajib diisi",
			"is_not_unique": "Identitas role tidak ditemukan",
		},
	},
	"Name": {
		Field: "name",
		Messages: map[string]string{
			"required":  "Nama role wajib diisi",
			"max":       "Nama maksimal 60 karakter",
			"is_unique": "Nama role sudah dipakai sebelumnya",
		},
	},
	"IsSuperadmin": {
		Field: "is_superadmin",
		Messages: map[string]string{
			"in_list": "Hanya boleh menggunakan value 'Super Admin' atau 'Bukan Super Admin'",
		},
	},
}
