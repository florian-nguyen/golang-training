package cipher

import (
	"regexp"
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

/*
STEP 1 :
Giving "iamapandabear" as input to the encode function returns the
cipher "ldpdsdqgdehdu". Obscure enough to keep our message secret in
transit.

When "ldpdsdqgdehdu" is put into the decode function it would return
the original "iamapandabear" letting your friend read your original
message.
*/

type Caesar struct {
}

func NewCaesar() *Caesar {
	return &Caesar{}
}

func (c Caesar) Encode(s string) string {
	s = StringFormat(s)
	return ApplyShift(s, 3)
}

func (c Caesar) Decode(s string) string {
	s = StringFormat(s)
	return ApplyShift(s, -3)
}

/*
STEP 2 :
Fixed distance Shift Ciphers are no fun though when your kid sister
figures it out. Try amending the code to allow us to specify a shift
distance.

You will implement a more generic Shift Cipher with a flexible shift
distance.
*/

type Shift struct {
	offset int
}

func NewShift(n int) *Shift {
	if n >= 26 || n <= -26 || n == 0 {
		return nil
	}
	return &Shift{offset: n}
}

func (c Shift) Encode(s string) string {
	s = StringFormat(s)
	return ApplyShift(s, c.offset)
}

func (c Shift) Decode(s string) string {
	s = StringFormat(s)
	return ApplyShift(s, -c.offset)
}

/*
With only 26 true possible shift values, your kid sister will figure
this out too. Next lets define a more complex cipher using a string as
key value: a [Vigenère cipher][vc].

Here's an example:

Given the key "aaaaaaaaaaaaaaaaaa", encoding the string
"iamapandabear" would return the original "iamapandabear".

Given the key "ddddddddddddddddd", encoding our string "iamapandabear"
would return the obscured "ldpdsdqgdehdu"
*/

type Vigenere struct {
	offset string
}

func NewVigenere(s string) *Vigenere {
	reg, _ := regexp.Compile(`[^a-z]`)
	if len(s) == 0 || s != reg.ReplaceAllString(s, "") || len(strings.Replace(s, "a", "", -1)) == 0 {
		return nil
	}
	return &Vigenere{offset: s}
}

func (c Vigenere) Encode(s string) string {
	s = StringFormat(s)
	var output string
	var n int
	for key, value := range s {
		n = int([]rune(c.offset)[key%(len(c.offset))] - 'a')
		output += ApplyShift(string(value), n)
	}
	return output
}

func (c Vigenere) Decode(s string) string {
	s = StringFormat(s)
	var output string
	var n int
	for key, value := range s {
		n = int(-[]rune(c.offset)[key%(len(c.offset))] + 'a')
		output += ApplyShift(string(value), n)
	}
	return output
}

// Function StringFormat is used to remove all unnecessary characters from input string
func StringFormat(s string) string {
	reg, _ := regexp.Compile(`[^a-zA-Z]+`)
	return reg.ReplaceAllString(strings.ToLower(s), "")
}

// Function ApplyShift adds n to each letter of a string
func ApplyShift(s string, n int) string {
	runes := []rune(s)
	switch {
	case n > 0 && n < 26:
		for key, value := range runes {
			//runes[key] = rune(value + rune(n) - ('z'+1-'a')*((value+rune(n)-'a')/('z'+1-'a')))
			runes[key] = rune('a' + ((value + rune(n) - 'a') % ('z' + 1 - 'a')))
		}
		return string(runes)
		break

	case n == 0:
		return s
		break

	case n < 0 && n > -26:
		for key, value := range runes {
			runes[key] = rune('z' - (('z' - value - rune(n)) % ('z' + 1 - 'a')))
		}
		return string(runes)
		break
	}
	return string(runes)
}
