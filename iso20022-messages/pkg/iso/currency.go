package iso

type Currency string

func (a Currency) MarshalText() ([]byte, error) {
	return []byte(a), nil
}

func (a Currency) Validate() error {
	// TODO
	_, err := a.MarshalText()
	return err
}
