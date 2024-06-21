package addressbook

func AreStringsEqual(a, b string) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	return a == b
}
