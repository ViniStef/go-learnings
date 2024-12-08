package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	formattedString := fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)

	return formattedString
}