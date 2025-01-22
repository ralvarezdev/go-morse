package international

import (
	gomorse "github.com/ralvarezdev/go-morse"
)

var (
	// Morse code characters
	A = gomorse.NewCharacter('A', ".-")
	B = gomorse.NewCharacter('B', "-...")
	C = gomorse.NewCharacter('C', "-.-.")
	D = gomorse.NewCharacter('D', "-..")
	E = gomorse.NewCharacter('E', ".")
	F = gomorse.NewCharacter('F', "..-.")
	G = gomorse.NewCharacter('G', "--.")
	H = gomorse.NewCharacter('H', "....")
	I = gomorse.NewCharacter('I', "..")
	J = gomorse.NewCharacter('J', ".---")
	K = gomorse.NewCharacter('K', "-.-")
	L = gomorse.NewCharacter('L', ".-..")
	M = gomorse.NewCharacter('M', "--")
	N = gomorse.NewCharacter('N', "-.")
	O = gomorse.NewCharacter('O', "---")
	P = gomorse.NewCharacter('P', ".--.")
	Q = gomorse.NewCharacter('Q', "--.-")
	R = gomorse.NewCharacter('R', ".-.")
	S = gomorse.NewCharacter('S', "...")
	T = gomorse.NewCharacter('T', "-")
	U = gomorse.NewCharacter('U', "..-")
	V = gomorse.NewCharacter('V', "...-")
	W = gomorse.NewCharacter('W', ".--")
	X = gomorse.NewCharacter('X', "-..-")
	Y = gomorse.NewCharacter('Y', "-.--")
	Z = gomorse.NewCharacter('Z', "--..")

	// Morse code numbers
	One   = gomorse.NewCharacter('1', ".----")
	Two   = gomorse.NewCharacter('2', "..---")
	Three = gomorse.NewCharacter('3', "...--")
	Four  = gomorse.NewCharacter('4', "....")
	Five  = gomorse.NewCharacter('5', ".....")
	Six   = gomorse.NewCharacter('6', "-....")
	Seven = gomorse.NewCharacter('7', "--...")
	Eight = gomorse.NewCharacter('8', "---..")
	Nine  = gomorse.NewCharacter('9', "----.")
	Zero  = gomorse.NewCharacter('0', "-----")

	// Morse code punctuation
	Period           = gomorse.NewCharacter('.', ".-.-.-")
	Comma            = gomorse.NewCharacter(',', "--..--")
	QuestionMark     = gomorse.NewCharacter('?', "..--..")
	LeftParenthesis  = gomorse.NewCharacter('(', "-.--.")
	RightParenthesis = gomorse.NewCharacter(')', "-.--.-")
	Apostrophe       = gomorse.NewCharacter('\'', ".----.")
	Semicolon        = gomorse.NewCharacter(';', "-.-.-.")
	Colon            = gomorse.NewCharacter(':', "---...")
	QuotationMark    = gomorse.NewCharacter('"', ".-..-.")
	Hyphen           = gomorse.NewCharacter('-', "-....-")
	Slash            = gomorse.NewCharacter('/', "-..-.")
	At               = gomorse.NewCharacter('@', ".--.-.")
	Equals           = gomorse.NewCharacter('=', "-...-")
	Plus             = gomorse.NewCharacter('+', ".-.-.")
	Ampersand        = gomorse.NewCharacter('&', ".-...")
	Dollar           = gomorse.NewCharacter('$', "...-..-")
	Exclamation      = gomorse.NewCharacter('!', "-.-.--")

	// Morse code prosigns
	Wait = gomorse.NewProvisionalSignal("WAIT", ".-...")
	Over = gomorse.NewProvisionalSignal("OVER", "-...-.")
	End  = gomorse.NewProvisionalSignal("END", "...-.-")

	// Morse code error
	Error = gomorse.NewProvisionalSignal("ERROR", "........")

	// Morse code invitation to transmit
	InvitationToTransmit = gomorse.NewProvisionalSignal(
		"INVITATION TO TRANSMIT",
		"-.-",
	)

	// Morse code starting signal
	StartingSignal = gomorse.NewProvisionalSignal("STARTING SIGNAL", "-.-.-")

	// Morse code distress signal
	DistressSignal = gomorse.NewProvisionalSignal(
		"DISTRESS SIGNAL",
		"...---...",
	)

	// Alphabet is a list of all Morse code characters
	Alphabet = gomorse.NewAlphabet(
		A,
		B,
		C,
		D,
		E,
		F,
		G,
		H,
		I,
		J,
		K,
		L,
		M,
		N,
		O,
		P,
		Q,
		R,
		S,
		T,
		U,
		V,
		W,
		X,
		Y,
		Z,
		One,
		Two,
		Three,
		Four,
		Five,
		Six,
		Seven,
		Eight,
		Nine,
		Zero,
		Period,
		Comma,
		QuestionMark,
		LeftParenthesis,
		RightParenthesis,
		Apostrophe,
		Semicolon,
		Colon,
		QuotationMark,
		Hyphen,
		Slash,
		At,
		Equals,
		Plus,
		Ampersand,
		Dollar,
		Exclamation,
	)

	// Morse code provisional signals
	ProvisionalSignals = gomorse.NewProvisionalSignals(
		Wait,
		Over,
		End,
		Error,
		InvitationToTransmit,
		StartingSignal,
		DistressSignal,
	)
)

// NewMorseCodeHandler creates a new Morse code handler with the international alphabet
func NewMorseCodeHandler() (*gomorse.MorseCodeHandler, error) {
	return gomorse.NewMorseCodeHandler(Alphabet, ProvisionalSignals)
}
