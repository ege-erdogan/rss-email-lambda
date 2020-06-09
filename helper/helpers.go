package helper

// Check panics with error
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
