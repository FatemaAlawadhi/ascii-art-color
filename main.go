package main

import (
	"ascii-art-color/pkg"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the input
	input := os.Args
	if len(input) > 5 {
		fmt.Print("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"\n")
	} else if len(input) < 2 {
		fmt.Print("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"\n")
	} else {
		var Text string
		var TextArray []string
		var Banner []string
		var FinalText []string
		var FileName string
		var TextString string
		var Color string

		for i := 0; i < len(input); i++ {
			if strings.HasPrefix(input[i], "--output=") || strings.HasPrefix(input[i], "--color=") {
				if i != 1 || (i == 1 && len(input) == 2) {
					fmt.Print("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"\n")
					return
				}
			}
			if input[i] == "shadow" || input[i] == "standard" || input[i] == "thinkertoy" {
				if (len(input) == 3 && i != 2) || (len(input) == 4 && i != 3) || len(input) == 2 || (len(input) == 5 && i != 4) {
					fmt.Print("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"\n")
					return
				}
			}
		}
		if input[1] == "--color" || input[1] == "--color="{
			fmt.Print("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"\n")
			return
		} 
		if strings.HasPrefix(input[1], "--color=") {
			Color = strings.TrimPrefix(input[1], "--color=")
			Color = pkg.CheckColor(Color)
		}

		if len(input) == 5 {
			Text = string(input[3])
			TextArray = strings.Split(Text, "\\n")
			Banner = pkg.Banner(input[4])
			if strings.HasPrefix(input[1], "--color=") {
				FinalText = pkg.PrintString(Banner, TextArray, Color, input[2])
				TextString = strings.Join(FinalText, "")
				fmt.Print(TextString)
			}
		} else if len(input) == 4 {
			if (input[3] != "shadow" && input[3] != "standard" && input[3] != "thinkertoy") && strings.HasPrefix(input[1], "--color=") {
				Text = string(input[3])
				TextArray = strings.Split(Text, "\\n")
				Banner = pkg.Banner("No Banner")
				FinalText = pkg.PrintString(Banner, TextArray, Color, input[2])
				TextString = strings.Join(FinalText, "")
				fmt.Print(TextString)
			} else {
				Text = string(input[2])
				TextArray = strings.Split(Text, "\\n")
				Banner = pkg.Banner(input[3])
				if strings.HasPrefix(input[1], "--output=") {
					FinalText = pkg.PrintString(Banner, TextArray, "No Color", "empty")
					FileName = strings.TrimPrefix(input[1], "--output=")
					pkg.CreateFile(FileName, FinalText)
				}
				if strings.HasPrefix(input[1], "--color=") {
					FinalText = pkg.PrintString(Banner, TextArray, Color, "empty")
					TextString = strings.Join(FinalText, "")
					fmt.Print(TextString)
				}
			}
		} else if len(input) == 3 {
			if strings.HasPrefix(input[1], "--output=") || strings.HasPrefix(input[1], "--color=") {
				Text = string(input[2])
				TextArray = strings.Split(Text, "\\n")
				Banner = pkg.Banner("No Banner")
				if strings.HasPrefix(input[1], "--output=") {
					FinalText = pkg.PrintString(Banner, TextArray, "No Color", "empty")
					FileName = strings.TrimPrefix(input[1], "--output=")
					pkg.CreateFile(FileName, FinalText)
				} else if strings.HasPrefix(input[1], "--color=") {
					FinalText = pkg.PrintString(Banner, TextArray, Color, "empty")
					TextString = strings.Join(FinalText, "")
					fmt.Print(TextString)
				}

			} else {
				Text = string(input[1])
				TextArray = strings.Split(Text, "\\n")
				Banner = pkg.Banner(input[2])
				FinalText = pkg.PrintString(Banner, TextArray, "No Color", "empty")
				TextString = strings.Join(FinalText, "")
				fmt.Print(TextString)
			}
		} else if len(input) == 2 {
			Text = string(input[1])
			TextArray = strings.Split(Text, "\\n")
			Banner = pkg.Banner("No Banner")
			FinalText = pkg.PrintString(Banner, TextArray, "No Color", "empty")
			TextString = strings.Join(FinalText, "")
			fmt.Print(TextString)
		}

		if Color == "" {
			fmt.Println("The Specified Color is not available")
		}

	}
}