.PHONY: bonk
bonk:
	go build -mod=readonly -ldflags="-s -w" -tags=providerless,dockerless -o $@ ./cmd/bonk.go

.PHONY: clean
clean:
	rm -rf bonk
