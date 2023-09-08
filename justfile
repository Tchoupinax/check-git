watch:
  npx nodemon -e go  --exec "go run *.go perso || exit 1"

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
  go tool cover -html=coverprofile.out
