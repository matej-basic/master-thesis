use axum::{routing::get, Router};
use std::net::SocketAddr;
use std::env;

#[tokio::main]
async fn main() {
    let app = Router::new().route("/benchmark", get(handler));
    let port = env::var_os("FUNCTIONS_CUSTOMHANDLER_PORT")
        .map(|value| {
		value.to_string_lossy()
   		     .parse::<i32>()
		     .unwrap_or(3000)
	})
	.unwrap_or(3000);

    let addr = SocketAddr::from(([0,0,0,0], port as u16));

    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn handler() -> &'static str {
    "Simple Rust benchmark for Axum"
}
