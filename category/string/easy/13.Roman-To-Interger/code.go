package string

func romanToInt(s string) int {
	var normalValue = make(map[string]int, 7)
	normalValue["I"] = 1
	normalValue["V"] = 5
	normalValue["X"] = 10
	normalValue["L"] = 50
	normalValue["C"] = 100
	normalValue["D"] = 500
	normalValue["M"] = 1000

	var result int
	var lastKeyValue int
	for _, runeValue := range s {
		if value, ok := normalValue[string(runeValue)]; ok {
			result += value
			if (value-lastKeyValue) > 0 && ((value-lastKeyValue)%4 == 0 || (value-lastKeyValue)%9 == 0) {
				result -= lastKeyValue * 2
			}
			lastKeyValue = value
		}
	}

	return result
}
