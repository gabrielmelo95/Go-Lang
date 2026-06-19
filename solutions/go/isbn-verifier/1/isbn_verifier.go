package isbnverifier

func IsValidISBN(isbn string) bool {
	var digits [10]int
	count := 0

	for i := 0; i < len(isbn); i++ {
		c := isbn[i]
		if c == '-' {
			continue
		}

		if count == 9 {
			if c == 'X' {
				digits[count] = 10
				count++
				continue
			}
		}

		if c < '0' || c > '9' {
			return false
		}

		if count >= len(digits) {
			return false
		}

		digits[count] = int(c - '0')
		count++
	}

	if count != len(digits) {
		return false
	}

	sum := 0
	for i, value := range digits {
		weight := 10 - i
		sum += value * weight
	}

	return sum%11 == 0
}
