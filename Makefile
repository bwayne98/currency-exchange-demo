testc: 
	@echo "測試開始"
	go test ./... -coverprofile="./coverage.out"
	go tool cover -func=coverage.out

.PHONY: testc