package go_morse

type (
	// MorseCode is the interface for the Morse code
	MorseCode interface {
		Encode(text string) string
		EncodeToBytes(text string) []byte
		Decode(morseCode string) string
		DecodeFromBytes(morseCode []byte) string
	}
)
