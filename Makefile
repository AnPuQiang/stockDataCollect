PROGRAMS = stock_main

all: clean $(PROGRAMS)

test: check

stock_main:
	go build -i -o stock_main ./main/main.go
clean:
	@rm -rf $(PROGRAMS)