all:
	docker-compose up --build -d

kill:
	docker-compose down 

PROTO = proto/task_service.proto

generate_proto:
	@echo "Generating for task-service..."
	protoc \
		--go_out=task-service/ --go_opt=paths=source_relative,Mproto/task.proto=github.com/AbhinitKumarRai/task-management-system/task-service/ \
		--go-grpc_out=task-service/ --go-grpc_opt=paths=source_relative,Mproto/task.proto=github.com/AbhinitKumarRai/task-management-system/task-service/ \
		$(PROTO)

	@echo "Generating for user-service..."
	protoc \
		--go_out=user-service/ --go_opt=paths=source_relative,Mproto/task.proto=github.com/AbhinitKumarRai/task-management-system/user-service/ \
		--go-grpc_out=user-service/ --go-grpc_opt=paths=source_relative,Mproto/task.proto=github.com/AbhinitKumarRai/task-management-system/user-service/ \
		$(PROTO)
