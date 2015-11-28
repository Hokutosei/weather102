build:
	#build for linux
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/linux_amd64 -v
	docker build -t gcr.io/chat-app-proto01/weather102:latest .
	gcloud docker push gcr.io/chat-app-proto01/weather102:latest
