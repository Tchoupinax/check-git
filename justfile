lint:
  golangci-lint run

fmt:
  gofmt -s

test:
  ginkgo -r --randomize-all --randomize-suites --race ./...

test-watch:
  ginkgo watch -r --randomize-all --randomize-suites --race ./...

test-details:
  ginkgo -r --randomize-all --randomize-suites --race --trace --cover ./...

coverage:
  go tool cover -html=coverage.out
