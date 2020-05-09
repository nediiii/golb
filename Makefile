.PYHON: tidy
tidy:
	go mod tidy

.PYHON: run
run:
	go run server.go

.PYHON: gen
gen:
	go run github.com/99designs/gqlgen

.PYHON: start
start:
	make gen
	make run


.PYHON: up
up:
	docker-compose down
	docker-compose up -d

.PYHON: down
down:
	docker-compose down



.PYHON: dps
dps:
	docker ps -a

dcp:
	docker container prune -f

dip:
	docker image prune -f

dnp:
	docker network prune -f

dvp:
	docker volume prune -f

dap:
	make dcp dip dnp dvp
