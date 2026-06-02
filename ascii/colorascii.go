package ascii

import "strings"

func AsciiColor(lines []string, colorName string) []string {

	

	Colors := map[string]string{
		"Black"   : "\033[30m",
		"Red"     : "\033[31m",
		"Green"   : "\033[32m",
		"Yellow"  : "\033[33m",
		"Blue"    : "\033[34m",
		"Magenta" : "\033[35m",
		"Cyan"    : "\033[36m",
		"White"   : "\033[37m",
	}
	const Reset = "\033[0m"
	
	lowerkey := map[string]string{}

	for k, v := range Colors{
		lowerkey[strings.ToLower(k)] = v
	}

	result := make([]string, len(lines))

	color, ok := lowerkey[colorName]
	if !ok {
		return lines
	}

	for i, line := range lines{
		result[i] = color + line + Reset
	}
	
	return result
}
