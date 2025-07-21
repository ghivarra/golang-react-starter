package auth

type ClaimData struct {
	JTI string   // Token ID, this is important for verification on revoked token
	SUB string   // Typically would be username
	ISS string   // Typically would be app name
	AUD []string // Typically would be user role id
	EXP int64    // expired at time
	IAT int64    // issued at time
}

type GeneratedToken struct {
	Token string
	Error error
	Data  ClaimData
}

type RefreshTokenData struct {
	AccessToken  string
	RefreshToken string
	Error        error
}
