package ast

// import ""

type MixinStatement struct {
	Token *Token
	Ident *Token
	Block *Block
}

func (stm MixinStatement) CanBeStatement() {}
func (stm MixinStatement) String() string  { return "" }

func NewMixinStatementWithToken(tok *Token) *MixinStatement {
	return &MixinStatement{Token: tok}
}
