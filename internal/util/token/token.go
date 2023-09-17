package token

type TokenMaker interface {
	GenerateToken(uid int64, name string, role int8) (string, error)
	VerifyToken(tokenString string) (*Payload, error)
}
