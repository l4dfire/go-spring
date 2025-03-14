module github.com/go-spring/starter-grpc

go 1.14

require (
	github.com/go-spring/spring-base v1.1.0-rc2
	github.com/go-spring/spring-core v1.1.0-rc2
	google.golang.org/grpc v1.41.0
)

replace (
	github.com/go-spring/spring-core => ../../spring/spring-core
)
