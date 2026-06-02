package ascii

import "strings"

func ValidateInput(text string) bool {

	for _, ch := range text {

		if ch < ' ' || ch > '~' {
			return false
		}
	}
	return true
}

// func AddBorder(lines []string) []string {
// 	if len(lines) == 0 {
// 		return lines
// 	}

// 	Width := 0

// 	for _, line := range lines {
// 		if len(line) > Width {
// 			Width = len(line)
// 		}
// 	}

// 	borderLength := Width + 2
// 	dashes := strings.Repeat("-", borderLength)
// 	horizontalBorder := "+" + dashes + "+"

// 	var styledOutput []string

// 	styledOutput = append(styledOutput, horizontalBorder)

// 	// Frame and pad each text row
// 	for _, line := range lines {
// 		currentLength := len(line)
// 		missingSpaces := Width - currentLength

// 		padding := strings.Repeat(" ", missingSpaces)

// 		framedLine := "| " + line + padding + " |"
// 		styledOutput = append(styledOutput, framedLine)
// 	}

// 	styledOutput = append(styledOutput, horizontalBorder)

// 	return styledOutput
// }


func AddBorder(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	width := 0
	for _, line := range lines{
		if len(line) > width{
			width = len(line) 
		}
	}


	border := "+" + strings.Repeat("_", width) + "+"
	result := []string{border}
	
	for _, line := range lines {
		result = append(result, "|"+line+"|")
	}

	result = append(result, border)
	return result
}

