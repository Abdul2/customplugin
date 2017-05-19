
default: build plan

deps:
	go install github.com/hashicorp/terraform

build:
	go build -o terraformhoplugin

test:
	go test -v .

plan:
	@terraform plan
