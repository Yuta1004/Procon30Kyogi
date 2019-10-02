docker-build:
	make docker-build-base
	make docker-build-solver

docker-build-base: Dockerfile_Base
	docker build -t alpine:procon30-solver-base -f Dockerfile_Base ./

docker-build-solver: Dockerfile_Solver
	docker build -t procon30-solver -f Dockerfile_Solver ./
