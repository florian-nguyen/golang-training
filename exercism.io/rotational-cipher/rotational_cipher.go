package rotationalcipher

// RotationalCipher rotates all letters from a given value n between 0-25
func RotationalCipher(input string, n int) (s string) {
	runes := []rune(input)
	for _, value := range runes {
		switch {
		case value >= 'a' && value <= 'z':
			s += string('a' + (value-'a'+rune(n))%26)
		case value >= 'A' && value <= 'Z':
			s += string('A' + (value-'A'+rune(n))%26)
		default:
			s += string(value)
		}
	}
	return s
}
