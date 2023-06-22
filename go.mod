module iot

go 1.20

require (
	github.com/go-redis/redis v6.15.9+incompatible
	huaweicloud.com/go-runtime v0.0.0-00010101000000-000000000000
)

require (
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.27.8 // indirect
)

replace huaweicloud.com/go-runtime => ./go-runtime
