.PHONY: lint format test

MODULES := $(shell find . -name 'go.mod' -exec dirname {} \;)

lint:
	@for mod in $(MODULES); do \
		echo "Linting $$mod..."; \
		cd $$mod && golangci-lint run ./... && cd - > /dev/null; \
	done

format:
	@for mod in $(MODULES); do \
		echo "Formatting $$mod..."; \
		cd $$mod && golangci-lint fmt ./... && cd - > /dev/null; \
	done

test:
	@for mod in $(MODULES); do \
		echo "Testing $$mod..."; \
		cd $$mod && go test -v ./... && cd - > /dev/null; \
	done
