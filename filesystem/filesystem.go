package filesystem

import (
	"log"
	"os"

	customerrors "github.com/joaosp7/passwordCLI/errors"
	"github.com/joaosp7/passwordCLI/helpers"
)

func CreateDirectory(path string) error{
	
	if (len(path) <= 0) {
		return customerrors.ErrDirectoryInvalidString
	}
	_, err := os.Stat(path)

	if os.IsNotExist(err){
		errDir := os.MkdirAll(path, 0755) // 0755 creates with permission to Read/Write/Execute
		if errDir != nil{
			log.Fatal(errDir)
			return errDir
		}
		return nil

	}else if err!=nil {
		log.Fatal("Error creating directory: ",err)
		return err
	}
	return customerrors.ErrDirectoryExists
}


func CreateFile(path string, passwords string) error {
	if (len(path) == 0) {
		return customerrors.ErrDirectoryInvalidString
	}
	pathFile :=  path + "/passwords.txt"
	err := os.WriteFile(pathFile, []byte(passwords), 0755)
	if err != nil {
		log.Fatal("Unable to write file: ", err)
		return err
	}
	return nil 
}

func GenerateTxt(path string, passwords []string) error {
	err := CreateDirectory(path)
	if err!= nil {
		log.Fatal("Unable to create directory.")
		return err
	}
	passString := helpers.PassowrdsToString(passwords)
	err = CreateFile(path, passString )
	if (err != nil){
		log.Fatal("Unable to create file.")
		return err
	}
	return nil
}