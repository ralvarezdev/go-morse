package international

import (
	gomorse "github.com/ralvarezdev/go-morse"
)

// NewMorseCodeHandler creates a new Morse code handler with the international alphabet
//
// Returns:
//
// *gomorse.MorseCodeHandler: Morse code handler
// error: error if any
func NewMorseCodeHandler() (*gomorse.MorseCodeHandler, error) {
	return gomorse.NewMorseCodeHandler(Alphabet, ProvisionalSignals)
}
