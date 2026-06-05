package raindrops

import "strconv"

func Convert(number int) string {
	var str string
	if number%3 == 0 || number%5 == 0 || number%7 == 0 {
		if number%3 == 0 {
			str = "Pling"
		}
		if number%5 == 0 {
			str += "Plang"
		}
		if number%7 == 0 {
			str += "Plong"
		}
	} else {
		str = strconv.Itoa(number)
	}
	return str
}
