package common

// common index error
var IndexError map[string]ErrorMessageInterface = map[string]ErrorMessageInterface{
	"Limit": {
		Field: "limit",
		Messages: map[string]string{
			"required": "Limit wajib diisi",
			"numeric":  "Harus diisi dengan nomor",
		},
	},
	"Offset": {
		Field: "offset",
		Messages: map[string]string{
			"required": "Offset wajib diisi",
			"numeric":  "Harus diisi dengan nomor",
		},
	},
	"ExcludeID": {
		Field: "excludeID",
		Messages: map[string]string{
			"required": "ID item yang tidak akan ditarik wajib diisi atau isi dengan array kosong",
		},
	},
	"Order": {
		Field: "order",
		Messages: map[string]string{
			"required": "Kolom urutan/order wajib diisi",
		},
	},
	"AccountIndexForm.Order.Name": {
		Field: "order.name",
		Messages: map[string]string{
			"required": "Nama kolom urutan/order wajib diisi",
		},
	},
	"AccountIndexForm.Order.Dir": {
		Field: "order.dir",
		Messages: map[string]string{
			"required": "Arah atau target urutan/order wajib diisi",
			"in_list":  "Arah urutan hanya bisa diisi dengan 'asc' untuk ascending dan 'desc' untuk descending",
		},
	},
	"AccountIndexForm.Query.Name": {
		Field: "order.name",
		Messages: map[string]string{
			"required": "Nama kolom urutan/order wajib diisi",
		},
	},
	"QueryColumn": {
		Field: "query.column",
		Messages: map[string]string{
			"required": "Nama kolom yang akan dijadikan parameter query wajib diisi",
		},
	},
	"QueryCommand": {
		Field: "query.command",
		Messages: map[string]string{
			"required": "Command kolom yang akan query wajib diisi",
			"in_list":  "Command hanya bisa diisi oleh is,is_not,contain,not_contain",
		},
	},
	"QueryValue": {
		Field: "query.value",
		Messages: map[string]string{
			"required": "Value parameter query tidak boleh kosong bila dikirim",
		},
	},
}
