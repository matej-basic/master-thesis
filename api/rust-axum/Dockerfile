FROM rust:1-slim-buster AS build
RUN cargo new --bin app
WORKDIR /app
COPY Cargo.toml /app/
RUN cargo build --release  
COPY src /app/src
RUN touch src/main.rs
RUN cargo build --release

FROM debian:buster-slim  
COPY --from=build /app/target/release/rust-axum /app/main  
EXPOSE 8080
CMD "/app/main"
