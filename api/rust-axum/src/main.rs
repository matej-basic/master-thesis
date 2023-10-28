use axum::{routing::get, Router};
use std::net::SocketAddr;
use std::env;

#[tokio::main]
async fn main() {
    let app = Router::new().route("/benchmark", get(handler));
    let port = match env::var("FUNCTIONS_CUSTOMHANDLER_PORT") {
        Some(v) => v.into_int().unwrap(),
        None => "3000"
    };

    let addr = SocketAddr::from(([0,0,0,0], 3000));

    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn handler() -> &'static str {
    "Simple Rust benchmark for Axum"
}
