package process

import "errors"

// invalid input
func IsValidInput(str string) (rune, error) {
	for _, chara := range str {
		if chara < 32 || chara > 126 {
			return chara, errors.New("Invalid input")
		}
	}
	return 0, nil
}
