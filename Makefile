SOURCE_PY=solver.py
SOLVER_IMAGE = procon30-solver


build:
	make build-darwin
	make build-linux

build-darwin: main.go
	env GOOS=darwin GOARCH=amd64 go build -x -o procon30_yuge_kyogi_darwin main.go

build-linux: main.go
	env GOOS=linux GOARCH=amd64 go build -x -o procon30_yuge_kyogi_linux main.go

docker-build:
	python3 gen_solver_image.py

docker-build-base: Dockerfile_Base
	docker build -t alpine:procon30-solver-base -f Dockerfile_Base ./

docker-build-solver: Dockerfile_Solver
	docker build -t $(SOLVER_IMAGE) --build-arg SOURCE_PY=$(SOURCE_PY) -f Dockerfile_Solver ./

dist:
	mkdir dist
	make build
	mv procon30_yuge_kyogi_* dist
	cp config.toml dist
	cp Dockerfile_* dist
	cp Makefile dist
	cp gen_solver_image.py dist
	cp -r docs dist
	cp -r solvers dist

clean:
	rm -rf tmp/ procon30_yuge_kyogi_*
