package main

import (
	"fmt"
)

type authenticationInfo struct {
	userName string
	password string
}

// Criando um metodo
func (authInfo authenticationInfo) geBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: basic USERNAME: '%s' PASSWORD '%s'",
		authInfo.userName,
		authInfo.password)

}

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.geBasicAuth())
	fmt.Println("=======================")
}

func main() {

	test(authenticationInfo{
		userName: "Gabriel V",
		password: "12345678",
	})
}
