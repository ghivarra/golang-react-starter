package variable

// base path from current working directory (main.go)
var BasePath string

// controller path
var ControllerPath string

// library path
var LibraryPath string

// set Request Type: between CLI or APP
var RequestType string

// Allowed arguments in Command Line Interface
var CLIAllowedArguments []string = []string{"db:migrate", "db:reset", "db:seed", "db:refresh", "make:migration", "make:controller"}
