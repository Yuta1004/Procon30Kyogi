build:
	make build-darwin

build-darwin: main.go
	env GOOS=darwin GOARCH=amd64 go build -x -o procon30_yuge_kyogi_darwin main.go

build-linux: main.go
	env GOOS=linux GOARCH=amd64 go build -x -o procon30_yuge_kyogi_linux main.go

docker-build:
	make docker-build-base
	make docker-build-solver

docker-build-base: Dockerfile_Base
	docker build -t alpine:procon30-solver-base -f Dockerfile_Base ./

docker-build-solver: Dockerfile_Solver
	docker build -t procon30-solver -f Dockerfile_Solver ./
