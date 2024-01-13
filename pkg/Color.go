package pkg

func Color(ColoredLetters string, color string, LetterLine string, char rune) string {
	Reset := "\u001b[0m"

	if ColoredLetters == "" && char == ' ' {
		LetterLine = color + LetterLine + Reset
	}else {
		ColoredLetter := []rune(ColoredLetters)
		for i:=0; i<len(ColoredLetter); i++ {
			if char == ColoredLetter[i] {
				LetterLine = color + LetterLine + Reset
			}
		}
	}
	
	return LetterLine
}