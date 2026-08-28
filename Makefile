.PHONY: build clean run test fmt

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
