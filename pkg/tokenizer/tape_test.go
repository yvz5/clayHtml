package tokenizer

import "testing"

func TestHTMLTapePacking(t *testing.T) {
	tests := []struct {
		typeVal TokenType
		start   uint32
		length  uint32
	}{
		{TokenStartTag, 0, 10},
		{TokenAttrName, 123456, 32},
		{TokenAttrValue, 1 << 24, 1 << 20},
		{TokenComment, 0xFFFFFFFF, 0x3FFFFFF}, // max values
	}

	for _, tt := range tests {
		tape := newHTMLTape(tt.typeVal, tt.start, tt.length)
		if got := tape.Type(); got != tt.typeVal {
			t.Errorf("Type() = %v, want %v", got, tt.typeVal)
		}

		if got := tape.Start(); got != tt.start {
			t.Errorf("Start() = %v, want %v", got, tt.start)
		}

		if got := tape.Len(); got != tt.length {
			t.Errorf("Len() = %v, want %v", got, tt.length)
		}
	}
}
