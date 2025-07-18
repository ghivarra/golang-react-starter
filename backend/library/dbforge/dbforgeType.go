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
	Constraint      int
	Default         string // should use ` or backticks on string because NO ESCAPING HERE!
	IsNullable      bool
	IsUnsigned      bool
	IsAutoIncrement bool
	IsPrimaryIndex  bool
	IsUnique        bool
}

type TableIndex struct {
	Name string
}

type TableForeignKey struct {
	Name      string
	Column    string
	RefTable  string
	RefColumn string
	OnDelete  string // default would be restrict
	OnUpdate  string // default would be restrict
}
