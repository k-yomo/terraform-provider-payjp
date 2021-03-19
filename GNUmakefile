default: testacc

VERSION=0.0.3-snapshot

# Build provider binary and place it plugins directory to be able to sideload the built provider.
.PHONY: build
build:
	go build -o ~/.terraform.d/plugins/registry.terraform.io/k-yomo/payjp/${VERSION}/darwin_amd64/terraform-provider-payjp

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
