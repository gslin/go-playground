#
.DEFAULT:	all
.PHONY:		all clean

#
DIST?=	dist/
GO?=	go

#
all:
	"${GO}" build -o "${DIST}" ./cmd/https

clean:
	rm -rf "${DIST}"
