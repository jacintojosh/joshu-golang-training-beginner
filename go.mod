module joshua-golang-training-beginner

go 1.13

replace github.com/jacintojosh/weekzero => ./week-0

require (
	github.com/jacintojosh/rest v0.0.0-00010101000000-000000000000
	github.com/jacintojosh/weekone v0.0.0-00010101000000-000000000000
	github.com/jacintojosh/weekzero v0.0.0-00010101000000-000000000000
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.2.2
)

replace github.com/jacintojosh/weekone => ./week-1

replace github.com/jacintojosh/rest => ./rest
