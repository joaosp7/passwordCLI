package test

import (
	"os"
	"testing"

	customerrors "github.com/joaosp7/passwordCLI/errors"
	"github.com/joaosp7/passwordCLI/filesystem"
	"github.com/stretchr/testify/assert"
)

func TestCreateDirectory(t *testing.T){

	tests:= []struct{
		name string
		path string
		wantErr bool
		errMsg error
	}{
		{name: "Default Case",
			path: "./LocalTest", wantErr: false,
		},
		{
			name: "Directory Already Exists",
			path: "./LocalTest", wantErr: true, errMsg: customerrors.ErrDirectoryExists,
		},
		{
			name: "Invalid Directory String",
			path: "", wantErr: true, errMsg: customerrors.ErrDirectoryInvalidString,
		},
	}
	defer os.Remove(tests[0].path)	// maybe it does the trick
	for _,tt := range tests{

		t.Run(tt.name, func (t *testing.T)  {
			err := filesystem.CreateDirectory(tt.path)

	if (err!=nil){
		assert.True(t, tt.wantErr)
		assert.Error(t, err, tt.errMsg)
		return
	}
	assert.False(t, tt.wantErr)
	assert.NoError(t, err)


		})
	}



}