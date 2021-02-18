.PHONY : build
.PHONY : clean


build: deps bin ## Install dependancies and Build binaries

#Dependancies
deps: ## Install dependancies
	go get github.com/skycoin/dmsg/cipher

#Bin
bin: ## Build `activevisors`, `visorcount`, `visorverify`
	go build -o ./activevisors ./cmd/activevisors.go
	go build -o ./visorcount  ./cmd/visorcount.go
	go build -o ./visorverify ./cmd/visorverify.go
	go build -o ./genkeypair ./cmd/genkeypair.go
	go build -o ./skywire-scan ./cmd/skywire-scan.go
	go build -o ./skywire-connect ./cmd/skywire-connect.go
	go build -o ./wifi-scan ./cmd/wifi-scan.go
	go build -o ./address_gen ./cmd/address_gen/address_gen.go
	go build -o ./auxaddr_gen ./cmd/address_gen/auxaddr_gen_v0.1.go
	go build -o ./auxaddr_verify ./cmd/address_gen/auxaddr_verify.go
	go build -o ./swap ./cmd/swap.go
	go build -o ./skywire-get-remote-ap-pk ./cmd/skywire-get-remote-ap-pk.go

#Clean
clean: ## Remove binaries
	rm -fR ./activevisors
	rm -fR ./visorcount
	rm -fR ./visorverify
	rm -fR ./genkeypair
	rm -fR ./skywire-scan
	rm -fR ./wifi-scan
