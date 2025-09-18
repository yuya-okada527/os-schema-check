.PHONY: run

ca:
	git add .
	git commit -m "update"
	git push origin head
test:
	go run ./cmd/os-schema-check/main.go sample/schema.json sample/bulk_ok.jsonl
