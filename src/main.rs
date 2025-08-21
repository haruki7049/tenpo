use tenpo::cli::CLIArgs;
use bevy::prelude::*;
use clap::Parser;

fn main() {
    let _args: CLIArgs = CLIArgs::parse();

    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, tenpo::setup)
        .add_systems(Update, tenpo::sprite_movement)
        .add_systems(Update, bevy::input::keyboard::keyboard_input_system)
        .run();
}
