package hello

// Hello  Return a string with given name appended.
// name (string): Name to be appended after "Hello".
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name
}
