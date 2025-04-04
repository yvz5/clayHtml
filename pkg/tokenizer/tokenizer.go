package tokenizer

type TokenType uint8

const (
	TokenStartTag TokenType = iota
	TokenEndTag
	TokenAttrName
	TokenAttrValue
	TokenText
	TokenComment
	TokenDoctype
)
