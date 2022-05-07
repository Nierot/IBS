ifndef $(GOPATH)
	GOPATH = $(shell go env GOPATH)
	export GOPATH
endif

ifeq ($(OS),Windows_NT)
	AIR = air
else
	AIR = $(GOPATH)/bin/air;
endif
export AIR

dev:
	$(AIR)

install_air:
	go install github.com/cosmtrek/air@latest