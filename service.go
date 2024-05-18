package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func HandleGET(url string) {
	if url == "" {
		PrintError("url not found")
		return
	}

	// fetch
	response, err := http.Get(url)
	if err != nil {
		PrintError("cannot fatch url method GET, check your api")
		return
	}

	if response.StatusCode == http.StatusNotFound {
		PrintError("url not found")
		return
	}

	defer response.Body.Close()
	// io read all response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		PrintError("cannot read all response")
		return
	} else {
		fmt.Printf(successStuff.Render("success GET %s"), url)
		fmt.Printf(resultStuf.Render("Response body:\n\n%s"), PrettierJSONFortmater(body))
		return
	}
}

func HandlePOST(url, data string) {
	if url == "" && data == "" {
		PrintError("url & data not found")
		return
	}

	// post fetch
	response, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		PrintError("cannot POST data")
		return
	}

	if response.StatusCode == http.StatusNotFound {
		PrintError("url not found")
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		PrintError("cannot read response")
		return
	} else {
		fmt.Printf(successStuff.Render("success POST %s"), url)
		fmt.Printf(resultStuf.Render("Response body:\n\n%s"), PrettierJSONFortmater(body))
		return
	}

}
