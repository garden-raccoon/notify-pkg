module github.com/garden-raccoon/notify-pkg

go 1.23.2

require (
	github.com/gocql/gocql v1.7.0
	github.com/misnaged/scriptorium v0.0.0-20231207043744-47446928a2b9
	github.com/segmentio/kafka-go v0.4.47
	google.golang.org/grpc v1.67.1
	google.golang.org/protobuf v1.35.1
)

require (
	github.com/golang/snappy v0.0.3 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241104194629-dd2ea8efbc28 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.67.1
