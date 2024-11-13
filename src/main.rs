use remote_device_agent::{agent, network};

fn main() {
    println!("Starting Remote Device Agent...");
    agent::start_agent();

    if let Err(e) = network::start_server() {
        eprintln!("Failed to start server: {}", e);
    }
}