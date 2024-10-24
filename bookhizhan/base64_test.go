package bookhizhan

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	message :="Away from keyboard.https://golang.org"
	encodeMessage:=base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encodeMessage)

	data,err:=base64.StdEncoding.DecodeString(encodeMessage)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(string(data))
	}

	
}