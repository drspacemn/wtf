fn main() {
    capnpc::CompilerCommand::new()
        .src_prefix("schema")
        .file("schema/starwars.capnp")
        .run().expect("schema compiler command");
}
