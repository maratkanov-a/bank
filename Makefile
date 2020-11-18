INTERNAL_PKG_PATH=./internal/pkg
PROTO_FILES_PATH="./api"
LOCAL_BIN:=$(CURDIR)/bin
MINIMOCK_BIN:=$(LOCAL_BIN)/minimock

GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.32.2

GENVALIDATE_TAG:=0.3.0
GOOSE_TAG:=2.6.0

CLIENTS_PATH=pkg
ACCOUNTS_CLIENT_PATH=$(CLIENTS_PATH)/accounts
PAYMENTS_CLIENT_PATH=$(CLIENTS_PATH)/payments

IMPLEMENTATIONS_PATH=internal/app
ACCOUNTS_IMPLEMENTATION_PATH=$(IMPLEMENTATIONS_PATH)/accounts
PAYMENTS_IMPLEMENTATION_PATH=$(IMPLEMENTATIONS_PATH)/payments

PKGMAP:=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/api.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor,$\
        Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/source_context.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/type.proto=github.com/gogo/protobuf/types,$\
        Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types

# deps installation
.PHONY: install-minimock
install-minimock:
ifeq ($(wildcard $(MINIMOCK_BIN)),)
	$(info #Downloading minimock v$(MINIMOCK_TAG))
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && go get -d github.com/gojuno/minimock/v3@v$(MINIMOCK_TAG) && \
		go build -ldflags "-X 'main.version=$(MINIMOCK_TAG)' -X 'main.commit=test' -X 'main.buildDate=test'" -o $(LOCAL_BIN)/minimock github.com/gojuno/minimock/v3/cmd/minimock && \
		rm -rf $$tmp
MINIMOCK_BIN:=$(LOCAL_BIN)/minimock
endif

.PHONY: install-validate
install-validate:
	$(info #Downloading protoc-gen-validate v$(GENVALIDATE_TAG))
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && go get -d github.com/envoyproxy/protoc-gen-validate@v$(GENVALIDATE_TAG) && \
		go build -o $(LOCAL_BIN)/protoc-gen-validate github.com/envoyproxy/protoc-gen-validate && \
		rm -rf $$tmp

.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info #Downloading golangci-lint v$(GOLANGCI_TAG))
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && go get -d github.com/golangci/golangci-lint@v$(GOLANGCI_TAG) && \
		go build -ldflags "-X 'main.version=$(GOLANGCI_TAG)' -X 'main.commit=test' -X 'main.date=test'" -o $(LOCAL_BIN)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint && \
		rm -rf $$tmp
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

.PHONY: install-goose
install-goose:
	$(info #Downloading goose v$(GOOSE_TAG))
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && go get -d github.com/pressly/goose@v$(GOOSE_TAG) && \
		go build -o $(LOCAL_BIN)/goose github.com/pressly/goose && \
		rm -rf $$tmp

.PHONY: install-gendoc
install-gendoc:
	$(info #Downloading protoc-gen-doc)
	go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

REL_PATH_FROM_ACCOUNTS_TO_ROOT=$(shell echo $(ACCOUNTS_CLIENT_PATH) | perl -F/ -lane 'print "../"x scalar(@F)')
REL_PATH_FROM_PAYMENTS_TO_ROOT=$(shell echo $(PAYMENTS_CLIENT_PATH) | perl -F/ -lane 'print "../"x scalar(@F)')
.PHONY: .generate
.generate:
	mkdir -p $(ACCOUNTS_CLIENT_PATH) && cd $(ACCOUNTS_CLIENT_PATH) && protoc --plugin=protoc-gen-goclay=$(LOCAL_BIN)/protoc-gen-goclay \
		--plugin=protoc-gen-gofast=$(LOCAL_BIN)/protoc-gen-gofast \
		-I$(REL_PATH_FROM_ACCOUNTS_TO_ROOT)api/:$(CURDIR)/vendor.pb \
		--gofast_out=$(PKGMAP),plugins=grpc:. \
		--goclay_out=$(PKGMAP),impl=true,impl_path=$(REL_PATH_FROM_ACCOUNTS_TO_ROOT)$(ACCOUNTS_IMPLEMENTATION_PATH),impl_type_name_tmpl=Implementation:. \
		$(REL_PATH_FROM_ACCOUNTS_TO_ROOT)api/accounts.proto
	mkdir -p $(PAYMENTS_CLIENT_PATH) && cd $(PAYMENTS_CLIENT_PATH) && protoc --plugin=protoc-gen-goclay=$(LOCAL_BIN)/protoc-gen-goclay \
		--plugin=protoc-gen-gofast=$(LOCAL_BIN)/protoc-gen-gofast \
		-I$(REL_PATH_FROM_PAYMENTS_TO_ROOT)api/:$(CURDIR)/vendor.pb \
		--gofast_out=$(PKGMAP),plugins=grpc:. \
		--goclay_out=$(PKGMAP),impl=true,impl_path=$(REL_PATH_FROM_PAYMENTS_TO_ROOT)$(PAYMENTS_IMPLEMENTATION_PATH),impl_type_name_tmpl=Implementation:. \
		$(REL_PATH_FROM_PAYMENTS_TO_ROOT)api/payments.proto

.generate-validation:
	protoc --plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate -I $(PROTO_FILES_PATH):./vendor.pb \
    		--validate_out="$(PKGMAP),lang=gogo:$(ACCOUNTS_CLIENT_PATH)" $(PROTO_FILES_PATH)/accounts.proto
	protoc --plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate -I $(PROTO_FILES_PATH):./vendor.pb \
    		--validate_out="$(PKGMAP),lang=gogo:$(PAYMENTS_CLIENT_PATH)" $(PROTO_FILES_PATH)/payments.proto


.PHONY: generate
generate: .generate .generate-validation

.PHONY: generate-mock
generate-mock: install-minimock
	find . -name '*_mock.go' -delete
	$(MINIMOCK_BIN) -g -i "$(INTERNAL_PKG_PATH)/repository.*" -o "$(INTERNAL_PKG_PATH)/repository/mock/" -s "_mock.go"

.PHONY: generate-docs
generate-docs:
	protoc -I $(PROTO_FILES_PATH):./vendor.pb --doc_out=./docs \
	--doc_opt=markdown,api.md $(PROTO_FILES_PATH)/*.proto

.PHONY: build
build:
	$(info #Building...)
	GOOS=linux GOARCH=amd64 go build -o $(LOCAL_BIN)/bank ./cmd/bank

.PHONY: lint
lint: install-lint
	$(info #Running lint...)
	$(GOLANGCI_BIN) run --new-from-rev=origin/master --config=.lint.yaml ./...

.PHONY: test
test:
	go test ./internal/... -count=1 -parallel=8 -run=$(run)

# integration part
.PHONY: test-integration
test-integration:
	go test ./test -count=1 -run=$(run)

.PHONY: test-repos
test-repos:
	 go test ./test/repos -timeout=60s -count=1 -run=$(run)

POSTGRES_SETUP_TEST:=user=test password=test dbname=bank_test host=localhost port=6432 sslmode=disable

# migrations for test db
.PHONY: test-migrations-up
test-migrations-up:
	goose -dir "$(INTERNAL_PKG_PATH)/db/migrations" postgres "${POSTGRES_SETUP_TEST}" up

.PHONY: test-migrations-down
test-migrations-down:
	goose -dir "$(INTERNAL_PKG_PATH)/db/migrations" postgres "${POSTGRES_SETUP_TEST}" down

# docker-compose aliases
.PHONY: compose-up
compose-up:
	docker-compose -p bank -f ./deployments/docker-compose.yml up -d

.PHONY: compose-rs
compose-rs:
	make compose-rm
	make compose-up

.PHONY: compose-rm
compose-rm:
	docker-compose -p bank -f ./deployments/docker-compose.yml rm -fvs

.PHONY: compose-down
compose-down:
	docker-compose -p bank -f ./deployments/docker-compose.yml stop

.PHONY: run
run:
	go run ./cmd/bank/main.go

.PHONY: create-image
create-image: build
	docker build --tag bank:$(version) --file ./build/ci/Dockerfile .

