.PHONY: run

ca:
	git add .
	git commit -m "update"
	git push origin head
run:
	go run main.go
