package seeder

import (
	"backend/database"
	"backend/module/model"
	"fmt"
)

// all of the seeder will run here
func Run() error {

	// connect db
	database.Connect(false)

	// seed role
	seedRole()

	// return
	return nil
}

// seed roles
func seedRole() {
	var role model.Role
	role.Name = "Super Admin"
	role.IsSuperadmin = 1

	// add
	if result := database.CONN.Create(&role); result.Error != nil {
		fmt.Printf("Failed to seed roles. Reason: %v", result.Error.Error())
	}
}
