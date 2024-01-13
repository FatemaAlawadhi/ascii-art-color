package pkg

import ("fmt"
        "strings")

func PrintString(Banner []string, TextArray []string, color string, ColoredLetters string) []string{
	var word []rune
	var LineBanner int
	var Line string
	var Text []string
	NewLine := 0 
	for str := 0; str < len(TextArray); str++ {
		word = []rune(TextArray[str])
		if len(word) > 0 {
			for Linenum := 1; Linenum <= 8; Linenum++ {
				for chr := 0; chr < len(word); chr++ {
					LineBanner = ((int(word[chr]) - 32) * 9) + (1 * Linenum)
					if LineBanner < 0 || LineBanner > 856 {
						fmt.Println("Please write in English")
						return Text
					}
					if color != "" && color != "No Color" && ColoredLetters != "empty" && strings.TrimSpace(ColoredLetters) != "" {
						Banner[LineBanner] = Color(ColoredLetters, color, Banner[LineBanner], word[chr])
					}
					Line = Line + Banner[LineBanner]
				}

				if color != "No Color" && ColoredLetters == "empty" {
					Line = Color("",color , Line , ' ')
				}
				Text = append(Text, Line)
				Text = append(Text, "\n")
				Line = ""
			}
		} else if len(word) == 0 {
			NewLine++
			if str == (len(TextArray)-1) && NewLine == len(TextArray){
				return Text
			}
			Text = append(Text, "\n")
		}
	}
	return Text
}