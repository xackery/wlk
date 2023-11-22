build:
	@echo "Building examples..."
	cd examples && make
run-%:
	@echo "Running examples/$*..."
	@cd examples/$* && wine64 $*