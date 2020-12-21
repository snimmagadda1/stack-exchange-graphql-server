package constants

type dialects struct {
	MySQL string
}

var (
	// Dialects is server supported languages
	Dialects = dialects{
		MySQL: "mysql",
	}
)
