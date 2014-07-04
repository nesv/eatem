PROG = eatem

all: ${PROG}

${PROG}: $(wildcard *.go)
	godep go build -o $@ .

clean:
	rm -f ${PROG}

.PHONY: all clean
