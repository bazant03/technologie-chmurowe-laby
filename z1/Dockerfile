#syntax=docker/dockerfile:1

# --------- ETAP 1 ------------------------
FROM rust AS build

RUN apt-get update && apt-get install -y --no-install-recommends musl-tools 

RUN useradd -u 10001 weather
USER weather

COPY --chown=weather:weather weather-app-rust /weather-app-rust
WORKDIR /weather-app-rust

# Kopilujemy aplikację do statycznego pliku wykonywalnego wykorzystującego musl 
# (implementacja libc lżejsza od domyślnej glibc
RUN rustup target add x86_64-unknown-linux-musl && \
    CC_x86_64_unknown_linux_musl=musl-gcc \
    cargo build --release --target x86_64-unknown-linux-musl

# --------- ETAP 2 ------------------------
FROM scratch AS final

LABEL org.opencontainers.image.authors="Marek Ruszecki <s101505@pollub.edu.pl>"

# UID i GID zgodne z użytkownikiem z etapu budowania
USER 10001:10001

COPY --from=build \ 
     --chown=10001:10001 \
     /weather-app-rust/target/x86_64-unknown-linux-musl/release/weather-rs \ 
     /weather-rs

# Wgrywamy certyfikat ca API do listy zaufanych certyfikatów aby łączyć z nim przez HTTPS
COPY --from=build \ 
     --chown=10001:10001 \
     /etc/ssl/certs/ca-certificates.crt \ 
     /etc/ssl/certs/ca-certificates.crt

EXPOSE 80

# Aplikacja ma tryb --healthcheck, który sprawdza czy główna aplikacja działa
HEALTHCHECK --interval=5s --timeout=3s --retries=3 CMD ["/weather-rs", "--healthcheck"]

ENTRYPOINT ["/weather-rs"]