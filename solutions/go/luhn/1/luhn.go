package luhn

func Valid(id string) bool {
	var sum int
	var double bool
	var num_len int

	for i := len(id) - 1; i >= 0; i-- {
		c := id[i]

		if c == ' ' {
			continue
		}

		if c < '0' || c > '9' {
			return false
		}

		num_len++

		digit := int(c - '0')

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return num_len > 1 && sum%10 == 0
}
