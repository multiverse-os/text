module github.com/multiverse-os/text

go 1.19

replace (
	github.com/multiverse-os/text/banner => ./banner
	github.com/multiverse-os/text/symbols => ./symbols
)

require golang.org/x/text v0.13.0
