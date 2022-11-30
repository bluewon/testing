utils_test:
	go test -v -count=1 -timeout 30s -run TestSquare github.com/bluewon/testing/test
covarage:
	go test ./... -coverprofile cover.out
tool_coverage:
	go tool cover -html=coverage.out -o coverage.html
integration_test:
	go test -v -count=0 -timeout 30s -run TestGetUserByEmail github.com/bluewon/testing/integration-test/user
	go test -v -count=1 -timeout 30s -run TestGetUserByEmail github.com/bluewon/testing/db/postgres