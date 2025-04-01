package tokenizer

type StructuralToken struct {
	Offset uint32  // Offset into the buffer
	Char   byte    // Structural char found ('<', '>', '=', '/', '"' ...)
	_      [3]byte // Explicit padding to 8 bytes
}

func FindStructurals(buf []byte) []StructuralToken {
	structurals := make([]StructuralToken, 0, len(buf)/4)

	for i := range buf {
		c := buf[i]
		switch c {
		case '<', '>', '=', '/', '"', '\'', '!', '-', '?':
			structurals = append(structurals, StructuralToken{
				Offset: uint32(i),
				Char:   c,
			})
		}
	}

	return structurals
}

func FindStructuralsSIMD(buf []byte) []StructuralToken {
	out := make([]uint32, len(buf)/4)
	n := findStructuralsSIMD(&buf[0], len(buf), &out[0], len(out))

	structurals := make([]StructuralToken, 0, n)
	for _, pos := range out[:n] {
		structurals = append(structurals, StructuralToken{
			Offset: pos,
			Char:   buf[pos],
		})
	}

	// do we have remainder < 32 bytes?
	if remainder := len(buf) % 32; remainder != 0 {
		start := len(buf) - remainder
		for i := start; i < len(buf); i++ {
			switch buf[i] {
			case '<', '>', '/', '=', '"', '\'', '!', '-', '?':
				structurals = append(structurals, StructuralToken{
					Offset: uint32(i),
					Char:   buf[i],
				})
			}
		}
	}

	return structurals
}

//go:noescape
func findStructuralsSIMD(buf *byte, length int, structurals *uint32, maxStructurals int) int
