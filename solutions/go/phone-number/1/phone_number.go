package phonenumber

import (
	"fmt"
)

func Number(phoneNumber string) (string, error) {
	var digits []byte
	for i := 0; i < len(phoneNumber); i++ {
		c := phoneNumber[i]
		if c == '+' || c == '-' || c == '.' || c == '(' || c == ')' || c == ' ' {
			continue
		}
		if c < '0' || c > '9' {
			return "", fmt.Errorf("Invalid phone number")
		}
		digits = append(digits, c)
	}

	if len(digits) == 11 && digits[0] == '1' {
		digits = digits[1:]
	}

	if len(digits) != 10 {
		return "", fmt.Errorf("Invalid phone number")
	}

	if digits[0] < '2' {
		return "", fmt.Errorf("Invalid phone number")
	}

	if digits[3] < '2' {
		return "", fmt.Errorf("Invalid phone number")
	}

	return string(digits), nil
}

func AreaCode(phoneNumber string) (string, error) {
	var err error
	phoneNumber, err = Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return phoneNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
	var err error
	phoneNumber, err = Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", phoneNumber[:3], phoneNumber[3:6], phoneNumber[6:]), nil
}
