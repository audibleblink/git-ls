APP=gitls
FLAGS=-trimpath -ldflags="-s -w -buildid="
PLATFORMS=windows linux darwin
OS=$(word 1, $@)

all: ${PLATFORMS}

${PLATFORMS}:
	GOOS=${OS} go build ${FLAGS} -o ${APP}_${OS}

clean:
	rm ${APP}_*
