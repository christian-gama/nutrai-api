package unit

const (
	// Byte is a byte, 1 byte.
	Byte = 1

	// Kilobyte is a kilobyte, 1024 bytes.
	Kilobyte = 1024

	// Megabyte is a megabyte, 1024 kilobytes.
	Megabyte = 1024 * Kilobyte

	// Gigabyte is a gigabyte, 1024 megabytes.
	Gigabyte = 1024 * Megabyte
)

var (
	// Alphabet is a list of all the letters in the alphabet, both upper and lower case.
	Alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// Numbers is a list of all the numbers.
	Numbers = []rune("0123456789")

	// AlphaNumeric is a list of all the letters in the alphabet,
	// both upper and lower case, and all the numbers.
	AlphaNumeric = append(Alphabet, Numbers...)

	// Symbols is a list of all the symbols.
	Symbols = []rune("!@#$%^&*()_+-=[]{};':,./<>?|\\`~\"'")
)
