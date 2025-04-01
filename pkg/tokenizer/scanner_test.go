package tokenizer

import "testing"

func TestFindStructuralsSIMD(t *testing.T) {
	type expected struct {
		Pos  uint32
		Char byte
	}

	tests := []struct {
		name  string
		input string
		want  []expected
	}{
		{
			name:  "basic tag",
			input: "<div>",
			want: []expected{
				{0, '<'},
				{4, '>'},
			},
		},
		{
			name:  "tag with attr and double quotes",
			input: `<div id="x">`,
			want: []expected{
				{0, '<'},
				{7, '='},
				{8, '"'},
				{10, '"'},
				{11, '>'},
			},
		},
		{
			name:  "end tag",
			input: "</div>",
			want: []expected{
				{0, '<'},
				{1, '/'},
				{5, '>'},
			},
		},
		{
			name:  "attr with single quotes",
			input: `<a href='x'>`,
			want: []expected{
				{0, '<'},
				{7, '='},
				{8, '\''},
				{10, '\''},
				{11, '>'},
			},
		},
		{
			name:  "comment open/close",
			input: `<!-- hi -->`,
			want: []expected{
				{0, '<'},
				{1, '!'},
				{2, '-'},
				{3, '-'},
				{8, '-'},
				{9, '-'},
				{10, '>'},
			},
		},
		{
			name:  "doctype",
			input: `<!DOCTYPE html>`,
			want: []expected{
				{0, '<'},
				{1, '!'},
				{14, '>'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			structurals := FindStructuralsSIMD([]byte(tt.input))

			if len(structurals) != len(tt.want) {
				t.Errorf("got %d structurals, want %d", len(structurals), len(tt.want))
			}

			for i, s := range structurals {
				if s.Offset != tt.want[i].Pos || s.Char != tt.want[i].Char {
					t.Errorf("entry %d: got (%d, %q), want (%d, %q)",
						i, s.Offset, s.Char, tt.want[i].Pos, tt.want[i].Char)
				}
			}
		})
	}
}
