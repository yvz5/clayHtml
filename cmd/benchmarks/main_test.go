package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yvz5/clayHtml/pkg/tokenizer"
)

const tokenCount = 41381

func TestFindStructuralsCompareWithSIMD(t *testing.T) {
	content, err := os.ReadFile("testdata/large.html")
	if err != nil {
		t.Fatalf("failed to read %s: %v", "testdata/large.html", err)
	}
	tokens := tokenizer.FindStructurals(content)
	tokensSIMD := tokenizer.FindStructuralsSIMD(content)

	assert.Equal(t, len(tokens), len(tokensSIMD), "token count mismatch")

	for i := range tokens {
		assert.Equal(t, tokens[i], tokensSIMD[i], "token mismatch at index %d", i)
	}
}

func BenchmarkFindStructurals(b *testing.B) {
	content, err := os.ReadFile("testdata/large.html")
	if err != nil {
		b.Fatalf("failed to read %s: %v", "testdata/large.html", err)
	}

	b.SetBytes(int64(len(content)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tokens := tokenizer.FindStructurals(content)
		if len(tokens) != tokenCount {
			b.Fatalf("token count mismatch: got %d, want %d", len(tokens), tokenCount)
		}
	}
}

func BenchmarkFindStructuralsSIMD(b *testing.B) {
	content, err := os.ReadFile("testdata/large.html")
	if err != nil {
		b.Fatalf("failed to read %s: %v", "testdata/large.html", err)
	}

	b.SetBytes(int64(len(content)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tokens := tokenizer.FindStructuralsSIMD(content)
		if len(tokens) != tokenCount {
			b.Fatalf("token count mismatch: got %d, want %d", len(tokens), tokenCount)
		}
	}
}
