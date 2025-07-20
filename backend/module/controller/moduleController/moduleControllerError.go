package moduleController

import "backend/library/common"

var ModuleCreateError map[string]common.ErrorMessageInterface = map[string]common.ErrorMessageInterface{
	"Name": {
		Field: "name",
		Messages: map[string]string{
			"required":  "Anda belum mengisi nama.",
			"max":       "Nama maksimal 200 karakter.",
			"is_unique": "Nama modul sudah dipakai",
		},
	},
	"Alias": {
		Field: "alias",
		Messages: map[string]string{
			"required": "Anda belum mengisi nama alias.",
			"max":      "Username maksimal 200 karakter.",
		},
	},
}
