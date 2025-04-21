package filesystem

import (
	"log"
	"os"

	customerrors "github.com/joaosp7/passwordCLI/errors"
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

