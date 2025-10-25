use bevy::prelude::*;
use clap::Parser;
use tenpo::cli::CLIArgs;

fn main() {
    let _args: CLIArgs = CLIArgs::parse();

    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup) // (Formerly setup_audio)
        .add_systems(
            Update,
            (
                play_metronome, // (Formerly play_sound_on_timer)
                check_key_press,
            ),
        )
        .run();
}

// Resource to manage the metronome state
// (Replaces Timer)
#[derive(Resource)]
struct MetronomeState {
    // Seconds per beat (1.0 for BPM 60)
    seconds_per_beat: f32,
    // Index of the next beat to play (at 1.0s, 2.0s...)
    next_beat_index: u64,
}

// Set the initial metronome state at startup
fn setup(mut commands: Commands) {
    commands.insert_resource(MetronomeState {
        seconds_per_beat: 1.0,
        // The first beat is at 1.0 * 1 = 1.0s
        next_beat_index: 1,
    });
}

// Play metronome sound based on absolute time since app start
fn play_metronome(
    mut commands: Commands,
    time: Res<Time>, // Total time since app start
    mut metronome: ResMut<MetronomeState>,
    asset_server: Res<AssetServer>,
) {
    let current_time = time.elapsed_secs();
    let mut next_beat_time = metronome.next_beat_index as f32 * metronome.seconds_per_beat;

    // Loop until the current time exceeds the next beat time
    // (This catches up by playing missed beats if frames are dropped)
    while current_time >= next_beat_time {
        // Play the metronome sound
        commands.spawn(AudioPlayer::new(
            asset_server.load("sounds/maou_se_8bit02.ogg"),
        ));

        // Advance to the next beat
        metronome.next_beat_index += 1;
        // Calculate the time for the next beat
        next_beat_time = metronome.next_beat_index as f32 * metronome.seconds_per_beat;
    }
}

// Check key press based on absolute time
fn check_key_press(
    mut commands: Commands,
    input: Res<ButtonInput<KeyCode>>,
    time: Res<Time>,                // Use Time instead of SoundTimer
    metronome: Res<MetronomeState>, // Reference seconds_per_beat
    asset_server: Res<AssetServer>,
) {
    if input.just_pressed(KeyCode::Space) {
        const ALLOWED_ERROR_MARGIN: f32 = 0.05;
        let spb = metronome.seconds_per_beat; // 1.0

        // "Absolute time" when the key was pressed
        let current_time = time.elapsed_secs();

        // Calculate the "nearest" beat time to the current time
        // (e.g., 2.98s -> 3.0s, 3.03s -> 3.0s, 3.4s -> 3.0s)
        let nearest_beat_time = (current_time / spb).round() * spb;

        // Absolute difference
        let delta = (current_time - nearest_beat_time).abs();

        if delta <= ALLOWED_ERROR_MARGIN {
            // Success sound
            commands.spawn(AudioPlayer::new(
                asset_server.load("sounds/maou_se_8bit16.ogg"),
            ));

            info!(
                "OK!! time: {current_time:.3}, nearest_beat: {nearest_beat_time:.1}, delta: {delta:.3}"
            );
        } else {
            info!(
                "Mistake. time: {current_time:.3}, nearest_beat: {nearest_beat_time:.1}, delta: {delta:.3}"
            );
        }
    }
}
