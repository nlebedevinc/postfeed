lint:
	golangci-lint run -c ./golangci.yml ./...

# Generate docs
# Require gomarkdoc (https://github.com/princjef/gomarkdoc)
docs:
	gomarkdoc -o README.md -e .