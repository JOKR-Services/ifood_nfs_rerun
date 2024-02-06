tidy:
	@export GOPRIVATE=github.com/JOKR-Services/*; go mod tidy && go mod vendor