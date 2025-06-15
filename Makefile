clean:
	$(MAKE) -C src clean

build:
	$(MAKE) -C src build

clean.build: clean build

fmt:
	$(MAKE) -C src fmt

test:
	$(MAKE) -C src test

test.v:
	$(MAKE) -C src test.v

test.cov:
	$(MAKE) -C src test.cov

test.bench:
	$(MAKE) -C src test.bench

test.bench.profile:
	$(MAKE) -C src test.bench.profile

.PHONY: test
