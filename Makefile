all: kubectl-tap

kubectl-tap:
	go build -o bin/kubectl-tap ./cmd/kubectl-tap

release: changelog
	GOOS=darwin GOARCH=amd64 go build -o bin/kubectl-tap-darwin-amd64 ./cmd/kubectl-tap
	GOOS=linux GOARCH=amd64 go build -o bin/kubectl-tap-linux-amd64 ./cmd/kubectl-tap

changelog:
	docker run -it --rm -v "$(pwd)":/usr/local/src/your-app ferrarimarco/github-changelog-generator:1.15.0 -u erwinvaneyk -p kubectl-tap