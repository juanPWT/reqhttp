package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var errorStuff = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5733")).Bold(true).Margin(2).Padding(2)
var successStuff = lipgloss.NewStyle().Foreground(lipgloss.Color("#75FF3A")).Bold(true).Margin(1).Padding(2)
var resultStuf = lipgloss.NewStyle().Margin(1)

func PrintError(s string) {
	fmt.Println(errorStuff.Render(s))
}

func PrettierJSONFortmater(body []byte) string {
	var formatedJson bytes.Buffer

	// formated
	err := json.Indent(&formatedJson, body, "", "  ")
	if err != nil {
		PrintError("cannot format json")
		return ""
	}

	return formatedJson.String()
}
