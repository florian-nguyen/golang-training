package utils

func CheckNilString(str string) *string {
	if len(str) == 0 {
		return nil
	}

	return &str
}
