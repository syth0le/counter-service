.SILENT:

run:
	docker-compose up -d

rebuild:
	docker-compose up -d --build

proto:
	protoc -I proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative proto/internalapi/counter_service.proto
