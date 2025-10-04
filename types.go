package go_morse

import (
	"fmt"
	"strings"
)

type (
	// MorseCodeHandler is the struct for the Morse code handler
	MorseCodeHandler struct {
		alphabet      Alphabet
		signals       ProvisionalSignals
		unicodeToCode map[rune]string
		codeToUnicode map[string]rune
	}

	// Character is the struct for the Morse code character
	Character struct {
		unicode rune
		code    string
	}

	// ProvisionalSignal is the struct for the provisional signal
	ProvisionalSignal struct {
		signal string
		code   string
	}

	// Alphabet is the Morse code alphabet
	Alphabet *[]*Character

	// ProvisionalSignals is the provisional signals
	ProvisionalSignals *[]*ProvisionalSignal
)

// NewCharacter creates a new Morse code character
//
// Parameters:
//
//   - unicode: The Unicode character
//   - code: The Morse code representation
//
// Returns:
//
//   - *Character: A pointer to the newly created Character struct
func NewCharacter(unicode rune, code string) *Character {
	return &Character{
		unicode,
		code,
	}
}

// GetUnicode returns the character in Unicode
//
// Returns:
//
//   - rune: The Unicode character
func (c Character) GetUnicode() rune {
	return c.unicode
}

// GetCode returns the code
//
// Returns:
//
//   - string: The Morse code representation
func (c Character) GetCode() string {
	return c.code
}

// NewProvisionalSignal creates a new Morse code provisional signal
//
// Parameters:
//
//   - signal: The signal representation (e.g., SOS)
//   - code: The Morse code representation of the signal
//
// Returns:
//
//   - *ProvisionalSignal: A pointer to the newly created ProvisionalSignal struct
func NewProvisionalSignal(signal string, code string) *ProvisionalSignal {
	return &ProvisionalSignal{
		signal,
		code,
	}
}

// GetSignal returns the signal
//
// Returns:
//
// - string: The signal representation (e.g., SOS)
func (p ProvisionalSignal) GetSignal() string {
	return p.signal
}

// GetCode returns the code
//
// Returns:
//
// - string: The Morse code representation of the signal
func (p ProvisionalSignal) GetCode() string {
	return p.code
}

// NewAlphabet creates a new Morse code alphabet
//
// Parameters:
//
//   - characters: A variadic list of pointers to Character structs
//
// Returns:
//
//   - Alphabet: A pointer to the newly created Alphabet
func NewAlphabet(characters ...*Character) Alphabet {
	return &characters
}

// NewProvisionalSignals creates a new Morse code provisional signals
//
// Parameters:
//
//   - signals: A variadic list of pointers to ProvisionalSignal structs
//
// Returns:
//
//   - ProvisionalSignals: A pointer to the newly created ProvisionalSignals
func NewProvisionalSignals(signals ...*ProvisionalSignal) ProvisionalSignals {
	return &signals
}

// NewMorseCodeHandler creates a new Morse code handler
//
// Parameters:
//
//   - alphabet: The Morse code alphabet
//
// - signals: The Morse code provisional signals
//
// Returns:
//
// - *MorseCodeHandler: A pointer to the newly created MorseCodeHandler struct
// - error: An error if the Unicode or code already exists in the alphabet
func NewMorseCodeHandler(
	alphabet Alphabet,
	signals ProvisionalSignals,
) (*MorseCodeHandler, error) {
	// Create the Unicode to code and code to Unicode maps
	unicodeToCode := make(map[rune]string)
	codeToUnicode := make(map[string]rune)

	// Populate the maps
	var unicode int32
	var code string
	for _, character := range *alphabet {
		// Check if the character is already in the maps
		unicode = character.GetUnicode()
		if _, ok := unicodeToCode[unicode]; ok {
			return nil, fmt.Errorf(ErrUnicodeAlreadyExists, unicode, unicode)
		}

		// Check if the code is already in the maps
		code = character.GetCode()
		if _, ok := codeToUnicode[code]; ok {
			return nil, fmt.Errorf(ErrCodeAlreadyExists, code)
		}

		// Add the Unicode and code to the maps
		unicodeToCode[unicode] = code
		codeToUnicode[code] = unicode
	}

	return &MorseCodeHandler{
		alphabet,
		signals,
		unicodeToCode,
		codeToUnicode,
	}, nil
}

// Encode encodes the text to Morse code
//
// Parameters:
//
//   - text: The text to encode
//
// Returns:
//
// - string: The encoded Morse code
func (m MorseCodeHandler) Encode(text string) string {
	// Encode the text to Morse code
	var code string
	var words []string
	var codes []string
	for _, unicode := range strings.ToUpper(strings.Trim(text, " ")) {
		// Check if the Unicode is a space
		if unicode == ' ' {
			// Append the Morse code for the word
			words = append(words, strings.Join(codes, CodeSeparator))
			codes = []string{}
			continue
		}

		// Get the code for the Unicode
		code = m.unicodeToCode[unicode]

		// Append the Morse code
		codes = append(codes, code)
	}

	// Check if there are any codes left
	if len(codes) > 0 {
		words = append(words, strings.Join(codes, CodeSeparator))
	}

	return strings.Join(words, WordSeparator)
}

// EncodeToBytes encodes the text to Morse code bytes
//
// Parameters:
//
// - text: The text to encode
//
// Returns:
//
// - []byte: The encoded Morse code bytes
func (m MorseCodeHandler) EncodeToBytes(text string) []byte {
	return []byte(m.Encode(text))
}

// Decode decodes the Morse code to text
//
// Parameters:
//
// - morseCode: The Morse code to decode
//
// Returns:
//
// - string: The decoded text
func (m MorseCodeHandler) Decode(morseCode string) string {
	// Decode the Morse code to text
	var unicodes []rune
	var words []string
	var codes []string
	for _, word := range strings.Split(
		strings.Trim(morseCode, " "),
		WordSeparator,
	) {
		// Check if the word is empty
		if word == "" {
			continue
		}

		// Get the codes for the word
		codes = strings.Split(word, CodeSeparator)

		// Decode the codes to Unicode
		unicodes = []rune{}
		for _, code := range codes {
			// Get the Unicode for the code
			unicode := m.codeToUnicode[code]

			// Append the Unicode to the Unicode list
			unicodes = append(unicodes, unicode)
		}

		// Append the Unicode list to the words list
		words = append(words, string(unicodes))
	}

	return strings.Join(words, " ")
}

// DecodeFromBytes decodes the Morse code bytes to text
//
// Parameters:
//
// - morseCode: The Morse code bytes to decode
//
// Returns:
//
// - string: The decoded text
func (m MorseCodeHandler) DecodeFromBytes(morseCode []byte) string {
	return m.Decode(string(morseCode))
}
