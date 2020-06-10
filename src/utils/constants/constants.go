package constants

import "os"

var (

	// SETUP
	DefaultPort      = "8080"
	DefaultFlagDebug = false
	DefaultFlagTLS   = false

	// VERSIONING
	API               = "api/"
	CurrentAPIVersion = "v1/"

	// CONTROLLERS
	MaximumFetch   = 10
	LimitFastFetch = 5

	// AUTH
	RegistrationKey = os.Getenv("REG_KEY")
	JwtTokenKey     = os.Getenv("JWT_TKN")
)
