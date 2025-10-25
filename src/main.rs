use bevy::prelude::*;
use clap::Parser;
use tenpo::cli::CLIArgs;

fn main() {
    let _args: CLIArgs = CLIArgs::parse();

    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup_audio)
        .add_systems(Update, play_sound_on_timer)
        .run();
}

// サウンド再生間隔を管理するタイマー
#[derive(Resource)]
struct SoundTimer(Timer);

// 起動時にアセットをロードし、タイマーをリソースとして登録
fn setup_audio(mut commands: Commands) {
    // 1秒ごとに繰り返すタイマーを設定
    commands.insert_resource(SoundTimer(Timer::from_seconds(1.0, TimerMode::Repeating)));
}

// タイマーを更新し、指定時間が経過したらサウンドを再生
fn play_sound_on_timer(
    mut commands: Commands,
    time: Res<Time>,               // 経過時間を取得
    mut timer: ResMut<SoundTimer>, // タイマーリソース（可変）
    asset_server: Res<AssetServer>,
) {
    // 毎フレーム、タイマーを更新
    timer.0.tick(time.delta());

    // タイマーがちょうど完了したか（1秒経過したか）をチェック
    if timer.0.just_finished() {
        // サウンドを再生
        commands.spawn(AudioPlayer::new(
            asset_server.load("sounds/maou_se_8bit02.ogg"),
        ));
    }
}
