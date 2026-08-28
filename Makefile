.PHONY: build clean run test fmt lint

build:
	@./scripts/build.nu

clean:
	@./scripts/clean.nu

run: build
	./target/tenpo

test:
	@./scripts/test.nu

fmt:
	@./scripts/fmt.nu

lint:
	@./scripts/lint.nu
