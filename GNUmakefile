#
.DEFAULT:	all
.PHONY:		all clean

#
DIST?=	dist/
GO?=	go

#
all: dist/asm dist/https
	@true

dist/asm: cmd/asm/*
	"${GO}" build -o "${DIST}" ./cmd/asm

dist/https: cmd/https/*
	"${GO}" build -o "${DIST}" ./cmd/https

clean:
	rm -rf "${DIST}"
