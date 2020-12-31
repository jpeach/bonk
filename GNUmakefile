VERS := v1.20.1

.PHONY: bin/bonk
bin/bonk: deps/kubernetes@$(VERS)
	@mkdir -p bin
	go build -mod=readonly -ldflags="-s -w" -tags=providerless,dockerless -o $@ ./cmd/bonk.go
	for k in kubelet kubeadm kubectl kube-scheduler kube-proxy kube-controller-manager kube-apiserver ; do \
		( cd bin && ln -sf bonk $$k ); \
	done

# Modules in k/k have generated files that are necessary to build certain
# packages that we may depend on. Here we checkout the branch that we want,
# then generated the build files. The package is replaced in go.mod.
deps/kubernetes@$(VERS):
	@git clone --depth 1 --branch $(VERS) https://github.com/kubernetes/kubernetes.git $@
	@$(MAKE) -C $@ generated_files

.PHONY: clean
clean:
	rm -rf bin
	rm -rf deps
