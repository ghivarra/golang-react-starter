package maker

import (
	"backend/config/variable"
	"backend/library/common/filesystem"
	"fmt"
	"path"
	"strings"
)

func BuildMigration(tableName string) {

	// set error
	var err error

	// check if file already exist
	targetDir := path.Clean(fmt.Sprintf("%s/migration/tables/%sTable", variable.LibraryPath, tableName))
	targetFile := path.Clean(fmt.Sprintf("%s/%sTable.go", targetDir, tableName))

	if filesystem.PathExist(targetDir) {
		fmt.Println("The migration table has been created already")
		return
	}

	// check master
	masterPath := path.Clean(fmt.Sprintf("%s/cli/maker/master/migration.txt", variable.LibraryPath))
	masterData, err := filesystem.FileGetContent(masterPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	// create package name
	packageName := fmt.Sprintf("%sTable", tableName)

	// mutate master data
	masterData = strings.ReplaceAll(masterData, "#packageName", packageName)
	masterData = strings.ReplaceAll(masterData, "#tableName", tableName)

	// create file
	err = filesystem.FilePutContent(targetFile, masterData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print
	fmt.Println("The migration table has been created in path: " + targetFile)
}

func BuildController(controllerName string) {
	// set error
	var err error

	// check if file already exist
	targetDir := path.Clean(fmt.Sprintf("%s/%sController", variable.ControllerPath, controllerName))
	targetFile := path.Clean(fmt.Sprintf("%s/%sController.go", targetDir, controllerName))

	if filesystem.PathExist(targetDir) {
		fmt.Println("The controller file has been created already")
		return
	}

	// check master
	masterPath := path.Clean(fmt.Sprintf("%s/cli/maker/master/controller.txt", variable.LibraryPath))
	masterData, err := filesystem.FileGetContent(masterPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	// create package name
	packageName := fmt.Sprintf("%sController", controllerName)

	// mutate master data
	masterData = strings.ReplaceAll(masterData, "#packageName", packageName)

	// create file
	err = filesystem.FilePutContent(targetFile, masterData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print
	fmt.Println("The controller file has been created in path: " + targetFile)
}
