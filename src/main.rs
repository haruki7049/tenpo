use bevy::prelude::*;

fn main() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, tenpo::setup)
        .add_systems(Update, tenpo::sprite_movement)
        .add_systems(Update, bevy::input::keyboard::keyboard_input_system)
        .run();
}
