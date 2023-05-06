module github.com/Booking-Platform/accommodation-booking-webapp/api_gateway

go 1.19

replace github.com/Booking-Platform/accommodation-booking-webapp/common => ../common

require (
	github.com/Booking-Platform/accommodation-booking-webapp/common v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/rs/cors v1.9.0
	google.golang.org/grpc v1.55.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
