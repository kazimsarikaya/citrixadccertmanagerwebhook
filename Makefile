.PHONY: container build

all: build

build:
	./build.sh $(filter-out $@,$(MAKECMDGOALS))

container:
	./container.sh $(filter-out $@,$(MAKECMDGOALS))

project:
	./project.sh $(filter-out $@,$(MAKECMDGOALS))

%:
	@:
