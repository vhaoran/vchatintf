module github.com/vhaoran/vchatintf

go 1.13

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/go-kit/kit v0.9.0
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/vhaoran/vchat v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.2.0
	sigs.k8s.io/yaml v1.2.0 // indirect

)

replace github.com/vhaoran/vchat => ../vchat
