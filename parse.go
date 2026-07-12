package validator

func Parse(s string) (Number, error) {
	normalised, err := normalise(s)
	if err != nil {
		return "", err
	}

	if !checksumValid(normalised) {
		return "", ErrInvalidChecksum
	}

	return Number(normalised), nil
}

func Validate(s string) error {
	_, err := Parse(s)
	return err
}

func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}
