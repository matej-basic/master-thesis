use axum::{routing::get, Router};
use std::net::SocketAddr;
use std::env;

#[tokio::main]
async fn main() {
    let app = Router::new().route("/benchmark", get(handler));
    let addr = SocketAddr::from(([0,0,0,0], 3000));

    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn handler() -> &'static str {
    "Simple Rust benchmark for Axum"
}
