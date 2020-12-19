.PHONY: bin/bonk
bin/bonk:
	@mkdir -p bin
	go build -mod=readonly -ldflags="-s -w" -tags=providerless,dockerless -o $@ ./cmd/bonk.go
	for k in kubelet kubeadm kubectl kube-scheduler kube-proxy kube-controller-manager kube-apiserver ; do \
		( cd bin && ln -sf bonk $$k ); \
	done

.PHONY: clean
clean:
	rm -rf bin
