# go-morse

Morse code translator for Go. A generic handler encodes text to Morse and decodes it back using a configurable alphabet and provisional signals, with a ready-made International Morse handler. No third-party dependencies; requires Go 1.23.4 (per `go.mod`).

**Note:** This repository is archived and read-only.

## Installation

```bash
go get github.com/ralvarezdev/go-morse
```

## API

- **`MorseCode`** — interface: `Encode`, `EncodeToBytes`, `Decode`, `DecodeFromBytes`.
- **Types** — `Character` (`NewCharacter(unicode, code)`), `ProvisionalSignal` (`NewProvisionalSignal(signal, code)`), `Alphabet` (`NewAlphabet(...)`), `ProvisionalSignals` (`NewProvisionalSignals(...)`).
- **`NewMorseCodeHandler(alphabet, signals)`** — builds a custom translator; returns an error for an inconsistent alphabet, such as duplicate entries.
- **`international.NewMorseCodeHandler()`** — handler preloaded with `international.Alphabet` and `international.ProvisionalSignals`.
- **Separators** — `CodeSeparator` (`" "`) between letters, `WordSeparator` (`" / "`) between words.

## Usage

```go
h, err := international.NewMorseCodeHandler()
if err != nil {
    panic(err)
}

morse := h.Encode("SOS")
text := h.Decode(morse)
```

## Development

```bash
go build ./...
```

There are no tests.

## License

GNU General Public License v3.0 (see [LICENSE](LICENSE)).
