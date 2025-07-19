package dbforge

type Option struct {
	Engine  string
	Charset string
	Collate string
}

type Table struct {
	Name        string
	Columns     []TableColumn
	Indexes     []TableIndex
	ForeignKeys []TableForeignKey
}

type TableColumn struct {
	Name            string
	Type            string
	Length          *int    // optional
	Default         *string // optional, but if supplied then should use "`" or backticks on string because NO ESCAPING HERE!
	IsNullable      *bool   // optional
	IsUnsigned      *bool   // optional
	IsAutoIncrement *bool   // optional
	IsPrimaryIndex  *bool   // optional
	IsUnique        *bool   // optional
}

type TableIndex struct {
	Name string
}

type TableForeignKey struct {
	Name      string
	Column    string
	RefTable  string
	RefColumn string
	OnDelete  *string // optional, default would be restrict
	OnUpdate  *string // optional, default would be restrict
}
