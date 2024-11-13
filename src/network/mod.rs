use std::net::{TcpListener, TcpStream};
use std::io::{Read, Write};
use std::thread;

/// Starts a TCP server to listen for incoming connections (for local testing initially).
pub fn start_server() -> std::io::Result<()> {
    let listener = TcpListener::bind("127.0.0.1:7878")?;
    println!("Server listening on 127.0.0.1:7878");

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                thread::spawn(|| handle_client(stream));
            }
            Err(e) => eprintln!("Connection failed: {}", e),
        }
    }
    Ok(())
}

fn handle_client(mut stream: TcpStream) {
    let mut buffer = [0; 512];
    stream.read(&mut buffer).expect("Failed to read from stream");
    println!("Received a request!");

    let response = b"Hello from the agent!";
    stream.write(response).expect("Failed to write to stream");
}