# 1,简单赋值( := ) 编程语言中常规理解的赋值方式，只对当前语句的变量有效
# 2,递归赋值( = )赋值语句可能影响多个变量，所有目标变量相关的其他变量都受影响
# 3,条件赋值( ?= )如果变量未定义，则使用符号中的值定义变量。如果该变量已经赋值，则该赋值语句无效。
# 4,追加赋值( += ）原变量用空格隔开的方式追加一个新值

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