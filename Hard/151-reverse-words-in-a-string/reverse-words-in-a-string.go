func reverseWords(s string) string {
	reversed := ""
	word := ""

	for _, char := range s {

		if unicode.IsSpace(char) {
			if word != "" {
				if reversed == "" {
					reversed = word
				} else {
					reversed = fmt.Sprintf("%s %s", word, reversed)
				}
				word = ""
			}
		} else {
			word = fmt.Sprintf("%v%c", word, char)
		}
	}
	if word != "" {
		if reversed == "" {
			reversed = word
		} else {
			reversed = fmt.Sprintf("%s %s", word, reversed)
		}
	}

	return reversed
}

