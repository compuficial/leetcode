package main

// Symbol   Value
// I        1
// V        5
// X        10
// L        50
// C        100
// D        500
// M        1000
//
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
//
//

var romans = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToInt(s string) int {
	// fmt.Println("-------------------------")
	// fmt.Println("input string:", s)
	runes := []rune(s)
	result := 0
	i := 0

	for i < len(runes) {
		if i+1 < len(runes) {
			next := string(runes[i]) + string(runes[i+1])
			if value, exists := romans[next]; exists {
				result += value
				// fmt.Printf("character: %s, number: %d\n", next, value)
				i += 2
				continue
			}
		}

		romanNum := romans[string(runes[i])]
		result += romanNum
		// fmt.Printf("character: %c, number: %d\n", runes[i], romanNum)
		i++
	}

	// fmt.Printf("accumulated result: %d\n", result)
	// fmt.Println("-------------------------")
	return result
}

func main() {
	romanToInt("III")
	romanToInt("LVIII")
	romanToInt("MCMXCIV")
	romanToInt("IVIXCIV")
}
