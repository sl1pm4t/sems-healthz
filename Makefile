default: build

build: 
	gox -osarch="darwin/amd64 linux/amd64" -output "sems_healthz_{{.OS}}_{{.Arch}}"

tools:
	go get -v github.com/mitchellh/gox

.PHONY: build default tools