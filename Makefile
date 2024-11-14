gen-chat-app:
	protoc --go_out=. --go-grpc_out=. internal/handlers/chat_app.proto