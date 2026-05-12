use clap::Parser;
use std::sync::LazyLock;
use tenpo_xtask::builder::{Builder, TenpoBuilder};
use tenpo_xtask::cli::CLIArgs;

type Result = std::result::Result<(), Box<dyn std::error::Error>>;

fn main() -> Result {
    tracing_subscriber::fmt::init();

    tracing::debug!("Parsing CLI arguments...");
    let args = CLIArgs::parse();
    tracing::debug!("Parsed CLI arguments.");

    tracing::debug!("Creating a TenpoBuilder...");
    let builder = TenpoBuilder::new(args, CARGO.to_string());
    tracing::debug!("Created a TenpoBuilder.");

    builder.run()?;

    Ok(())
}

static CARGO: LazyLock<String> =
    LazyLock::new(|| std::env::var("CARGO").unwrap_or_else(|_| "cargo".to_string()));
