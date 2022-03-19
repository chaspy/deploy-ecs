.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o hello

.PHONY: docker
docker:
	docker build . -t chaspy/hello:latest
	docker tag chaspy/hello:latest 655123516369.dkr.ecr.${AWS_REGION}.amazonaws.com/hello:${VERSION}

.PHONY: login
login:
	aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin https://655123516369.dkr.ecr.${AWS_REGION}.amazonaws.com

.PHONY: push
push:
	docker push 655123516369.dkr.ecr.${AWS_REGION}.amazonaws.com/hello:${VERSION}
