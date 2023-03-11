package environment

import "os"

const (
	Development = "development"
	Test        = "test"
	Production  = "production"

	// Other constants in your app
	ApplicationName = "hackaton"
	SessionName     = "_hackaton_session"
)

func Current() string {
	if env := os.Getenv("GO_ENV"); env != "" {
		return env
	}

	return Development
}
