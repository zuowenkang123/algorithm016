package Week_01

import "fmt"

func getHint(secret string, guess string) string {
	x, y := 0, 0
	m := make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			x++
		} else {
			m[secret[i]]++
			m[guess[i]]--
		}
	}
	y = len(secret) - x
	for _, v := range m {
		if v > 0 {
			y -= v
		}
	}
	return fmt.Sprintf("%dA%dB", x, y)
}
