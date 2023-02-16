help: ## show help, shown by default if no target is specified
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

bench: ## run benchmarks
	go test -bench=. -benchtime 1s -cpu 1

benchmark-perflock: ## run benchmarks using perflock - https://github.com/aclements/perflock
	go install golang.org/x/perf/cmd/benchstat@latest
	perflock -governor 80% go test -count 10 -cpu 8 -run=^# -bench=Slice. | tee .bench.output
	benchstat .bench.output
