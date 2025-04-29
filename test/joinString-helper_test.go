package test

import (
	"testing"

	"github.com/joaosp7/passwordCLI/helpers"
	"github.com/stretchr/testify/assert"
)

func TestJoinString(t *testing.T){
	tests := []struct{
		name string
		passwords []string
		want string
		wantErr bool   
	}{
		{name: "Regular Test",
		 passwords: []string{"ABC", "BCA", "CAB"},
		 want: "ABC\nBCA\nCAB\n",
		 wantErr: false},
	}

	
	for _,tt := range tests{
	
		t.Run(tt.name, func(t *testing.T) {
		results := helpers.PassowrdsToString(tt.passwords)
		assert.Equal(t, results, tt.want)
	
	})
	
	}
	}