package tokenizer

// HTMLTape is a compact 64-bit encoded token (like simdjson tape).
type HTMLTape uint64

const (
	tapeShiftType  = 58 // shift for type (6 bits)
	tapeShiftStart = 26 // shift for start (32 bits)
	tapeShiftLen   = 0  // shift for length (26 bits)

	tapeMaskType  = uint64(0b111111) << tapeShiftType
	tapeMaskStart = uint64(0xFFFFFFFF) << tapeShiftStart
	tapeMaskLen   = uint64((1 << 26) - 1)
)

func newHTMLTape(tokenType TokenType, start, length uint32) HTMLTape {
	return HTMLTape(
		(uint64(tokenType&0b111111) << tapeShiftType) |
			(uint64(start) << tapeShiftStart) |
			(uint64(length) << tapeShiftLen),
	)
}

func (t HTMLTape) Type() TokenType {
	return TokenType((uint64(t) >> tapeShiftType) & 0b111111)
}

func (t HTMLTape) Start() uint32 {
	return uint32((uint64(t) >> tapeShiftStart) & 0xFFFFFFFF)
}

func (t HTMLTape) Len() uint32 {
	return uint32(uint64(t) & tapeMaskLen)
}
