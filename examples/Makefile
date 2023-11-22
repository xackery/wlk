build-all:
	@echo "Building all examples..."
	@for dir in */; do \
		echo "Building $$dir..."; \
		cd $$dir && GOOS=windows go build .; \
		cd ..; \
	done
run-%:
	@echo "Running $*..."
	@cd $* && wine64 $*