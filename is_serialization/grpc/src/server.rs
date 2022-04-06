use tonic::{transport::Server, Request, Response, Status};
use thang::say_server::{Say, SayServer};
use thang::{SayResponse, SayRequest};
use async_std::prelude::*;

pub mod thang {
    tonic::include_proto!("thang");
}

#[derive(Default)]
pub struct MySay {}

#[tonic::async_trait]
impl Say for MySay {
    async fn send(&self, request:Request<SayRequest>) -> Result<Response<SayResponse>, Status> {
        Ok(Response::new(SayResponse{
            message:format!("hello {}", request.get_ref().name),
        }))
    }
}

#[async_std::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:8085".parse().unwrap();

    let say = MySay::default();

    println!("server listening on {}... ", addr);

    Server::builder()
        .add_service(SayServer::new(say))
        .serve(addr)
        .await?;
    Ok(())
}