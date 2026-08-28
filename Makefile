.PHONY: build clean run test

build:
	@./scripts/build.nu

clean:
	@./scripts/clean.nu

run: build
	./target/tenpo

test:
	@./scripts/test.nu
