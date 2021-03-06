ARG DOCKER_IMAGE

FROM $DOCKER_IMAGE as builder

# This flags mut be set to ENV before this Dockerfile is run

ARG LDFLAG_DEVELOPER
ARG LDFLAG_HOMEPAGE
ARG LDFLAG_LICENSE
ARG LDFLAG_VERSION
ARG LDFLAG_EXPIRE
ARG LDFLAG_GIT
ARG LDFLAG_BUILD

RUN echo "LDFLAG_DEVELOPER: $LDFLAG_DEVELOPER" > /tmp/report \
    && echo "LDFLAG_HOMEPAGE: $LDFLAG_HOMEPAGE" >> /tmp/report \
    && echo "LDFLAG_LICENSE: $LDFLAG_LICENSE" >> /tmp/report \
    && echo "LDFLAG_VERSION: $LDFLAG_VERSION" >> /tmp/report \
    && echo "LDFLAG_EXPIRE: $LDFLAG_EXPIRE" >> /tmp/report \
    && echo "LDFLAG_GIT: $LDFLAG_GIT" >> /tmp/report \
    && echo "LDFLAG_BUILD: $LDFLAG_BUILD" >> /tmp/report \
    && cat /tmp/report \
    && mkdir -p /var/lib/dbus \
    && echo "machine-id" > /var/lib/dbus/machine-id

# add current checkout directory to Docker context

ADD . /go/src/discover/
WORKDIR /go/src/discover

# untar vendor.tar.gz to "vendor" directoy with its contents. the file vendor.tar.gz must be kept up-to-date to the project by the developer with the "vendor.bat" Batch file

RUN tar -xvf vendor.tar.gz > /dev/null

# just here to report to Teamcity build log the content of the Docker context
# RUN find

# build Windows-amd64

RUN mkdir dist-windows-amd64 \
	&& cp CHANGES.txt dist-windows-amd64 \
	&& cp README.pdf dist-windows-amd64 \
    && export CGO_ENABLED=0 \
    && export GOOS=windows \
    && export GOARCH=amd64 \
    && go build -mod=vendor -ldflags="-s -w -X main.LDFLAG_DEVELOPER=$LDFLAG_DEVELOPER -X main.LDFLAG_HOMEPAGE=$LDFLAG_HOMEPAGE -X main.LDFLAG_LICENSE=$LDFLAG_LICENSE -X main.LDFLAG_VERSION=$LDFLAG_VERSION -X main.LDFLAG_EXPIRE=$LDFLAG_EXPIRE -X main.LDFLAG_GIT=$LDFLAG_GIT -X main.LDFLAG_BUILD=$LDFLAG_BUILD" -o dist-windows-amd64/discover.exe .

# build Linux-amd64, also with Docker support

RUN mkdir dist-linux-amd64 \
	&& cp CHANGES.txt dist-linux-amd64 \
	&& cp README.pdf dist-linux-amd64 \
	&& cp Dockerfile-run dist-linux-amd64/Dockerfile \
	&& cp docker-compose.yml dist-linux-amd64 \
	&& cp docker-compose-up.bat dist-linux-amd64 \
	&& cp docker-compose-up.sh dist-linux-amd64 \
	&& cp docker-compose-down.bat dist-linux-amd64 \
	&& cp docker-compose-down.sh dist-linux-amd64 \
    && export CGO_ENABLED=0 \
    && export GOOS=linux \
    && export GOARCH=amd64 \
    && go build -mod=vendor -ldflags="-s -w -X main.LDFLAG_DEVELOPER=$LDFLAG_DEVELOPER -X main.LDFLAG_HOMEPAGE=$LDFLAG_HOMEPAGE -X main.LDFLAG_LICENSE=$LDFLAG_LICENSE -X main.LDFLAG_VERSION=$LDFLAG_VERSION -X main.LDFLAG_EXPIRE=$LDFLAG_EXPIRE -X main.LDFLAG_GIT=$LDFLAG_GIT -X main.LDFLAG_BUILD=$LDFLAG_BUILD" -o dist-linux-amd64/discover .

# build and tar for Linux-arm6

RUN mkdir dist-linux-arm6 \
	&& cp CHANGES.txt dist-linux-arm6 \
	&& cp README.pdf dist-linux-arm6 \
    && export CGO_ENABLED=0 \
    && export GOOS=linux \
    && export GOARCH=arm \
    && go build -mod=vendor -ldflags="-s -w -X main.LDFLAG_DEVELOPER=$LDFLAG_DEVELOPER -X main.LDFLAG_HOMEPAGE=$LDFLAG_HOMEPAGE -X main.LDFLAG_LICENSE=$LDFLAG_LICENSE -X main.LDFLAG_VERSION=$LDFLAG_VERSION -X main.LDFLAG_EXPIRE=$LDFLAG_EXPIRE -X main.LDFLAG_GIT=$LDFLAG_GIT -X main.LDFLAG_BUILD=$LDFLAG_BUILD" -o dist-linux-arm6/discover .

# tar all builds along with discover-tests.html test result file

RUN tar -czvf discover-$LDFLAG_VERSION-$LDFLAG_BUILD-windows-amd64.tar.gz -C dist-windows-amd64/ . \
    && tar -czvf discover-$LDFLAG_VERSION-$LDFLAG_BUILD-linux-amd64.tar.gz -C dist-linux-amd64/ . \
    && tar -czvf discover-$LDFLAG_VERSION-$LDFLAG_BUILD-linux-arm6.tar.gz -C dist-linux-arm6/ .

# report artifacts

RUN ls discover*.tar.gz > /tmp/report \
    && cat /tmp/report
