start_dev:
	go run main.go

start_prod:
	go run main.go

# download dependecies
download_deps:
	go mod download

# run tests
tests:
	go test -run ""