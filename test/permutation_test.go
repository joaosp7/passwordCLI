package test

import (
	"testing"

	customerrors "github.com/joaosp7/passwordCLI/errors"
	"github.com/joaosp7/passwordCLI/permutation"
	"github.com/stretchr/testify/assert"
)

func TestPermutation(t *testing.T){
tests := []struct{
	name string
	password string
	want string
	wantErr bool
	errMsg error   
}{
	{name: "Regular Test", password: "ABC",want: "BCA", wantErr: false},
	{name: "Empty String", password: "", wantErr: true, errMsg: customerrors.ErrEmptyInput},
	{name: "Non Asci Code", password: "@!&*#ç", wantErr: false, want: "@ç!*&#"},
	{name: "Remove WhiteSpaces from Password", password: " test", want: "stet", wantErr: false},
	{name: "Remove ALL (Regex) WhiteSpaces from Password", password: " t e s t ", want: "stet", wantErr: false},

}

for _,tt := range tests{

	t.Run(tt.name, func(t *testing.T) {
	perm := permutation.NewPermutator()
	results, err := perm.FindPermutations(tt.password)

	if (err!=nil){
		assert.Error(t, err, customerrors.ErrEmptyInput)
		assert.Nil(t, results, "Result permutation should be empty")
		return
	}

	assert.NoError(t, err)
	assert.Contains(t, results, tt.want)

})

}
}