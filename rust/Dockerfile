FROM --platform=$BUILDPLATFORM messense/rust-musl-cross:${TARGETARCH}-musl AS builder
ARG NAME=rust
ARG TARGETARCH
RUN <<EOF
set -eu
if [ $TARGETARCH = "amd64" ]; then
  echo "x86_64" > /arch;
elif [ $TARGETARCH = "arm64" ]; then
  echo "aarch64" > /arch;
else
  echo "Unsupported platform: $TARGETARCH";
  exit 1;
fi
EOF
COPY Cargo.* .
RUN <<EOF
set -eu
mkdir -p src
echo 'fn main() {}' > src/main.rs
cargo build --release --target $(cat /arch)-unknown-linux-musl
EOF
COPY src src
RUN <<EOF
set -eu
CARGO_BUILD_INCREMENTAL=true cargo build --release --target $(cat /arch)-unknown-linux-musl
cp target/$(cat /arch)-unknown-linux-musl/release/${NAME} target/release/app
EOF

FROM alpine
COPY --from=builder /home/rust/src/target/release/app /opt/app
EXPOSE 3000
CMD ["/opt/app"]
