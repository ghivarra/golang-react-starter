package bootstrap

import "backend/config/validation"

// add custom configuration in here
func Run() {
	// register validation
	validation.Register()
}
