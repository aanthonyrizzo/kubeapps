# syntax = docker/dockerfile:1

FROM rust:1.51 as builder

WORKDIR /pinniped-proxy

# Create a release build of pinniped-proxy itself.
COPY ./cmd/pinniped-proxy ./
RUN --mount=type=cache,target=/usr/local/cargo/registry \
    --mount=type=cache,target=/pinniped-proxy/target \
    cargo build --release
RUN --mount=type=cache,target=/pinniped-proxy/target \
    cp /pinniped-proxy/target/release/pinniped-proxy /pinniped-proxy/pinniped-proxy

FROM bitnami/minideb:buster
RUN apt-get update && apt-get install -y ca-certificates libssl1.1 && rm -rf /var/lib/apt/lists/*
COPY --from=builder /pinniped-proxy/pinniped-proxy /pinniped-proxy
ENV PATH="/:$PATH"

EXPOSE 3333
USER 1001
ENTRYPOINT [ "pinniped-proxy" ]
CMD [ "--help" ]
