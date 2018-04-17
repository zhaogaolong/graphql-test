.FORCE:

dev:
	go build -o test main/main.go
	DEV_ENV=dev ./test
