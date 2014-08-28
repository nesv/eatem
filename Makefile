PROG = terf

all: ${PROG}

${PROG}: $(wildcard cmd/${PROG}/*.go)
	godep go build -o $@ github.com/nesv/terf/cmd/terf

clean:
	rm -f ${PROG}

.PHONY: all clean
