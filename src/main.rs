use bevy::prelude::*;
use clap::Parser;
use tenpo::cli::CLIArgs;

fn main() {
    let _args: CLIArgs = CLIArgs::parse();

    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, tenpo::setup)
        .add_systems(Update, tenpo::sprite_movement)
        .add_systems(Update, bevy::input::keyboard::keyboard_input_system)
        .add_systems(Update, update_system)
        .run();
}

fn update_system(keyboard: Res<ButtonInput<KeyCode>>) {
    for key in keyboard.get_just_pressed() {
        match key {
            KeyCode::KeyA
            | KeyCode::KeyB
            | KeyCode::KeyC
            | KeyCode::KeyD
            | KeyCode::KeyE
            | KeyCode::KeyF
            | KeyCode::KeyG
            | KeyCode::KeyH
            | KeyCode::KeyI
            | KeyCode::KeyJ
            | KeyCode::KeyK
            | KeyCode::KeyL
            | KeyCode::KeyM
            | KeyCode::KeyN
            | KeyCode::KeyO
            | KeyCode::KeyP
            | KeyCode::KeyQ
            | KeyCode::KeyR
            | KeyCode::KeyS
            | KeyCode::KeyT
            | KeyCode::KeyU
            | KeyCode::KeyV
            | KeyCode::KeyW
            | KeyCode::KeyX
            | KeyCode::KeyY
            | KeyCode::KeyZ => {
                println!("{:?} Pressed", key)
            }
            _ => (),
        }
    }
}
