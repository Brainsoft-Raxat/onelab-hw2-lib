package testlibrary

func FlipString(str string) string {
	var newString string
	for _, ch := range str {
		if ch >= 65 && ch <= 90 {
			newString += string(ch + 32)
		} else if ch >= 97 && ch <= 122 {
			newString += string(ch - 32)
		} else {
			newString += string(ch)
		}
	}

	return newString
}