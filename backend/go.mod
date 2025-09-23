module github.com/georgirtodorov/protein-bot/backend

go 1.22.0

require (
	github.com/georgirtodorov/protein-bot v0.0.0
	github.com/gorilla/handlers v1.5.2
	github.com/gorilla/mux v1.8.1
	github.com/lib/pq v1.10.9
)

require github.com/felixge/httpsnoop v1.0.3 // indirect

replace github.com/georgirtodorov/protein-bot v0.0.0 => ../backend
