package menuController

import "backend/library/common"

var MenuCreateError map[string]common.ErrorMessageInterface = map[string]common.ErrorMessageInterface{
	"Alias": {
		Field: "alias",
		Messages: map[string]string{
			"required": "Alias atau nama wajib diisi",
			"max":      "Alias atau nama maksimal 200 karakter",
		},
	},
	"RouteName": {
		Field: "route_name",
		Messages: map[string]string{
			"required": "Nama router wajib diisi",
			"max":      "Nama router maksimal 200 karakter",
		},
	},
	"SortNumber": {
		Field: "sort_number",
		Messages: map[string]string{
			"required": "Urutan wajib diisi",
			"numeric":  "Urutan wajib berbentuk angka",
		},
	},
	"Icon": {
		Field: "icon",
		Messages: map[string]string{
			"max": "Icon maksimal 100 karakter",
		},
	},
}
