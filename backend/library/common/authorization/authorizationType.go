package authorization

type ClaimData struct {
	JTI string // Token ID, this is important for verification on revoked token
	SUB string // Typically would be username
	AUD uint64 // Typically would be user role id
	ISS string // Typically would be app name
	EXP int64  // expired at time
	IAT int64  // issued at time
}
