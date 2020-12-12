.PHONY : build
.PHONY : clean


build: deps bin ## Install dependancies and Build binaries

#Dependancies
deps: ## Install dependancies
	go get github.com/Skycoin/dmsg/cipher

#Bin
bin: ## Build `activevisors`, `visorcount`, `visorverify`
	go build -o ./activevisors ./cmd/activevisors.go
	go build -o ./visorcount  ./cmd/visorcount.go
	go build -o ./visorverify ./cmd/visorverify.go
	go build -o ./genkeypair ./cmd/genkeypair.go

#Clean
clean: ## Remove binaries
	rm -fR ./activevisors
	rm -fR ./visorcount
	rm -fR ./visorverify
	rm -fR ./genkeypair
