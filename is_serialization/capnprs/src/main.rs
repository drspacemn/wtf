mod rpc;
use std::time::Instant;

pub mod starwars_capnp {
    include!(concat!(env!("OUT_DIR"), "/starwars_capnp.rs"));
}

fn main() {
    let arg = std::env::args().nth(1);
    match arg {
        Some(a) => {
            match a.as_ref() {
                "client" => run_client(),
                _ => println!("didn't match"),
            }
        },
        None => rpc::server::run().expect("cannot run server"),
    }
}

fn run_client() {
    let tot = 10;
    let mut count = 0u32;
    println!("starting client...");

    let now = Instant::now();
    loop {
        if count == tot {
            break;
        }
        count += 1;
        let thang = rpc::client::show_human("inny".to_string());
        println!("resp.... {:?}", thang);
    }
    let elapsed = now.elapsed();
    println!("Elapsed: {:.2?} per run for {:?} cycles", elapsed/tot, tot);
}
