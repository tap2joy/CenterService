module github.com/tap2joy/CenterService

go 1.14

require (
	github.com/tap2joy/Protocols v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/spf13/viper v1.7.1
	go.elastic.co/apm/module/apmgrpc v1.11.0
	google.golang.org/grpc v1.37.0
	google.golang.org/grpc/examples v0.0.0-20210409234925-fab5982df20a // indirect
)

replace github.com/tap2joy/Protocols => ./proto
