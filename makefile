run: stop up
run-d: stop up-d
run-build: stop up-b
run-build-d: stop up-b-d

build-pb:
	protoc --go-grpc_out=paths=source_relative:. --go_out=paths=source_relative:. ./models/pb/*.proto

build-app:
	docker-compose -f docker-compose.yaml build app

up: 
	docker-compose -f docker-compose.yaml up

up-d: 
	docker-compose -f docker-compose.yaml up -d

up-b: build-app up

up-b-d: build-app up-d

stop:
	docker-compose -f docker-compose.yaml stop

down:
	docker-compose -f docker-compose.yaml down

test:
	env CALCULATOR_ENV=test go test ${ARGS} ./...

run-stand-alone:
	env CALCULATOR_ENV=dev go run main.go


run-client-example:
	env CALCULATOR_ENV=dev go run cmd/main.go -rent 1700 -rent-inc-rate 2.5 -down-payment -2000 -intrest 2.91 -term 20 -price 250000 -p-tax 2 -t-tax 0 -res 11

run-client:
	env CALCULATOR_ENV=dev go run cmd/main.go ${ARGS}