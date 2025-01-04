Simplest GoLang Redis & RabbitMQ IoT project

To run the project you need to:

1. Run redis with the following command: `docker run --name redis-test -p 6379:6379 -d redis`
2. Run rabbitmq with the following command:
   `docker run -d --hostname my-rabbit --name local-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:4-management`
3. Pull the project and run commands in following order:
    - `go mod download`
    - `go run receiver.go`
    - `go run sender.go`
4. You can see the logs in the receiver terminal and the sender terminal.

Enjoy!