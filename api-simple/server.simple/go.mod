module server.simple

go 1.13

require (
	github.com/gorilla/mux v1.7.3
	pkg/handler v0.0.1
	pkg/mw v0.0.1
)

replace (
	pkg/handler => ./pkg/handler
	pkg/mw => ./pkg/mw
)
