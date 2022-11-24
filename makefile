utils_test:
	go test -v -count=1 -timeout 30s -run TestSquare github.com/bluewon/testing/test
covarage:
	go test ./... -coverprofile cover.out
tool_coverage:
	go tool cover -html=coverage.out -o coverage.html