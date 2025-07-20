package authorization

type ClaimData struct {
	JTI string
	SUB string
	ISS string
	AUD uint64
	EXP int64
	IAT int64
}
