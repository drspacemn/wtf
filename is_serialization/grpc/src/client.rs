use thang::say_client::SayClient;
use thang::SayRequest;
use async_std::prelude::*;

pub mod thang {
    tonic::include_proto!("thang");
}

#[async_std::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let channel = tonic::transport::Channel::from_static("http://[::1]:8085")
        .connect()
        .await?;

    let mut client = SayClient::new(channel);

    let request = tonic::Request::new(
        SayRequest {
            name:String::from("bgoebel")
        },
    );

    let response = client.send(request).await?.into_inner();
    println!("RESEPONSE={:?}", response);
    Ok(())
}