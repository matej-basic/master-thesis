

// Provjereno radi ali je output pun JSON govana
use lambda_http::{service_fn, Error, Request};


async fn strhandler() -> Result<String, Error> {
    Ok("Simple Rust benchmark".to_string())
}

#[tokio::main]
async fn main() -> Result<(), Error> {
    lambda_http::run(service_fn(|_request: Request| {
        strhandler()
    })).await?;
    Ok(())
} 