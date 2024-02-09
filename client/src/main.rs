use std::time::Duration;

use lapin::{options::BasicPublishOptions, protocol::basic::AMQPProperties, Connection, ConnectionProperties};

#[tokio::main]
async fn main() -> Result<(), lapin::Error> {
    let conn = Connection::connect("amqp://renas:root@localhost:5671", ConnectionProperties::default()).await?;

    println!("CONNECTED");

    let channel = conn.create_channel().await?;

    let two_sec = Duration::from_secs(2);
    loop {
        let message = "This is a message from my Rust Producer.".as_bytes();
        channel.basic_publish("first-exchange", "", BasicPublishOptions::default(), message, AMQPProperties::default()).await?;

        println!("Sended a message");
        std::thread::sleep(two_sec)
    }
}