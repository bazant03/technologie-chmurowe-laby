# syntax=docker/dockerfile:1

FROM scratch AS build
ADD filesystem/alpine-minirootfs-3.23.3-x86_64.tar.gz / 

RUN apk update && apk add --no-cache go git openssh-client

RUN mkdir -p -m 0600 ~/.ssh && \
    ssh-keyscan github.com >> ~/.ssh/known_hosts

RUN --mount=type=ssh git clone git@github.com:bazant03/tchlab6.git /app

WORKDIR /app

# Program jest kompilowany/linkowany statycznie
RUN CGO_ENABLED=0 go build main.go

# ---------------------------------------------------------
FROM nginx

LABEL org.opencontainers.image.source="https://github.com/bazant03/tchlab6" \
      org.opencontainers.image.description="Example webapp in Go with nginx as reverse proxy" \
      org.opencontainers.image.authors="Marek Ruszecki <s101505@pollub.edu.pl>"

ARG VERSION
ENV VERSION=${VERSION}

COPY --from=build /app/main /srv/app
COPY nginx/app.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

HEALTHCHECK --interval=5s --timeout=1s  \ 
    CMD curl -f http://127.0.0.1 || exit 1

#Uruchamiamy aplikację wraz z przekazaniem jej informacji do wyświetlania na stronie
#a następnie uruchamiamy nginx
ENTRYPOINT [ "/bin/bash", "-c", "/srv/app $VERSION $HOSTNAME & nginx -g 'daemon off;'"]