use crate::starwars_capnp;
use capnp::capability::Promise;
use capnp_rpc::{pry, RpcSystem};
use capnp_rpc::twoparty::{VatNetwork};
use capnp_rpc::rpc_twoparty_capnp::{Side};
use futures::{AsyncReadExt, FutureExt};
use futures::task::LocalSpawnExt;

struct StarWars;

impl starwars_capnp::star_wars::Server for StarWars {
    fn show_human(&mut self, params: starwars_capnp::star_wars::ShowHumanParams, mut results: starwars_capnp::star_wars::ShowHumanResults) -> Promise<(), capnp::Error> {
        println!("showing human...");

        let request_reader = pry!(params.get());
        let _id = request_reader.get_id();
        results.get().set_name("Luke");
        results.get().set_appears_in(starwars_capnp::star_wars::human::AppearsIn::NewHope);
        results.get().set_home_planet("Mars");
        Promise::ok(())
    }

    fn create_human(&mut self, params: starwars_capnp::star_wars::CreateHumanParams, mut results: starwars_capnp::star_wars::CreateHumanResults) -> Promise<(), capnp::Error> {
        println!("create human...");

        let request_reader = pry!(params.get());
        results.get().set_id("123");
        results.get().set_name(request_reader.get_name().unwrap());
        results.get().set_appears_in(request_reader.get_appears_in().unwrap());
        results.get().set_home_planet(request_reader.get_home_planet().unwrap());
        Promise::ok(())
    }
}

pub fn run() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "0.0.0.0:8085";
    let mut exec = futures::executor::LocalPool::new();
    let spawner = exec.spawner();

    exec.run_until(async move {
        let listener = async_std::net::TcpListener::bind(&addr).await.unwrap();

        println!("metadata: tcp://{}", addr);
        let client: starwars_capnp::star_wars::Client = capnp_rpc::new_client(StarWars);

        loop {
            let (stream, _) = listener.accept().await.unwrap();
            stream.set_nodelay(true).unwrap();
            let (reader, writer) = stream.split();
            let network = VatNetwork::new(
                reader,
                writer,
                Side::Server,
                Default::default(),
            );

            let rpc_system = RpcSystem::new(Box::new(network), Some(client.clone().client));

            spawner.spawn_local(Box::pin(rpc_system.map(|_|())))?;
        }
    })
}