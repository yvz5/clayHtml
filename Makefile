bench:
	cd cmd/benchmarks && go test -bench=. -benchmem

analyze:
	fieldalignment ./...

.PHONY: bench analyze