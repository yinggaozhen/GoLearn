Go ?= $(which go)
Hello := "hello"
World := "world"

all: install

.PHONY: install
install:
	echo "install"
	echo "install2"

# world 是 hello 的前置依赖项
.PHONY: hello
hello: world
	echo $(Hello)

.PHONY: world
world:
	echo $(World)