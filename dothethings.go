package main

import "fmt"

var allTheMessages [5000]string
var messageInfo struct {
	from    string
	subject string
}

func getMessages() {
	allTheMessages := gmail.users().messages().list(q == "from:user@domain.com")
}

func printTheMessages() {
	// TODO: for loop ?
	fmt.Println(allTheMessages)
}
