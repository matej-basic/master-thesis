use std::env;
use warp::{http::Response, Filter};
use std::net::Ipv4Addr;

#[tokio::main]
async fn main() {
    let response = warp::any().map(|| {
        Response::builder()
        .body("Simple Rust benchmark")
    });
    let server = warp::get()
        .and(response);

    let port_key = "FUNCTIONS_CUSTOMHANDLER_PORT";
    let port: u16 = match env::var(port_key) {
        Ok(val) => val.parse().expect("Custom Handler port is not a number!"),
        Err(_) => 3000,
    };

    warp::serve(server).run((Ipv4Addr::LOCALHOST, port)).await
}
