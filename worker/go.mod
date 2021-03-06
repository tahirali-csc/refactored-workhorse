module github.com/workhorse/worker

go 1.14

replace github.com/workhorse/api => ../api/

replace github.com/workhorse/client => ../client/

replace github.com/workhorse/commons => ../commons/

require github.com/workhorse/client v0.0.0

require github.com/workhorse/api v0.0.0

require github.com/workhorse/commons v0.0.0

require (
	github.com/containerd/containerd v1.4.0 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v17.12.0-ce-rc1.0.20200424210312-4839b27a1fb9+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/go-git/go-git/v5 v5.1.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/r3labs/sse v0.0.0-20200819143619-1491ab50668f
	github.com/sirupsen/logrus v1.6.0 // indirect
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073
	google.golang.org/grpc v1.31.1 // indirect
	gopkg.in/yaml.v2 v2.2.4
)
