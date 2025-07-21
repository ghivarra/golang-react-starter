package seeder

import (
	"backend/database"
	"backend/module/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// all of the seeder will run here
func Run() error {

	// connect db
	database.Connect(false)

	// seed role
	seedRole()

	// seed user
	seedUser()

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

// seed user
func seedUser() {
	// generate password hash
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("User*12345"), bcrypt.DefaultCost)

	var user model.User
	user.Name = "Ghivarra Senandika"
	user.Username = "ghivarra"
	user.Email = "gsenandika@gmail.com"
	user.Password = string(passwordHash)
	user.RoleID = 1

	if result := database.CONN.Create(&user); result.Error != nil {
		fmt.Printf("Failed to seed users. Reason: %v", result.Error.Error())
	}
}
