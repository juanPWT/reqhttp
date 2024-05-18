package main

import (
	"os"
)

func main() {
	if len(os.Args) < 3 {
		PrintError("USAGE: reqhttp <METHOD> <URL> <DATA>")
		os.Exit(1)
	}

	method := os.Args[1]
	url := os.Args[2]

	var data string
	if len(os.Args) > 3 {
		data = os.Args[3]
	}

	switch method {
	case "GET":
		HandleGET(url)
	case "POST":
		HandlePOST(url, data)
	default:
		PrintError("Method http not found or not support, make issue in\nhttps://github.com/juanPWT/reqhttp")
	}
}
