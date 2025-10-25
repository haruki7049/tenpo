use bevy::prelude::*;
use clap::Parser;
use tenpo::cli::CLIArgs;

fn main() {
    let _args: CLIArgs = CLIArgs::parse();

    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup) // (旧 setup_audio)
        .add_systems(
            Update,
            (
                play_metronome, // (旧 play_sound_on_timer)
                check_key_press,
            ),
        )
        .run();
}

// メトロノームの状態を管理するリソース
// (Timerの代わり)
#[derive(Resource)]
struct MetronomeState {
    // 1拍の秒数 (BPM 60 の場合 1.0)
    seconds_per_beat: f32,
    // 次に鳴らすべき拍のインデックス (1.0秒地点, 2.0秒地点...)
    next_beat_index: u64,
}

// 起動時にメトロノームの初期状態を設定
fn setup(mut commands: Commands) {
    commands.insert_resource(MetronomeState {
        seconds_per_beat: 1.0,
        // 最初の拍は 1.0 * 1 = 1.0秒 地点
        next_beat_index: 1,
    });
}

// アプリ起動からの絶対時間に基づき、メトロノーム音を再生
fn play_metronome(
    mut commands: Commands,
    time: Res<Time>, // アプリ起動時からの総時間
    mut metronome: ResMut<MetronomeState>,
    asset_server: Res<AssetServer>,
) {
    let current_time = time.elapsed_secs();
    let mut next_beat_time = metronome.next_beat_index as f32 * metronome.seconds_per_beat;

    // 現在時刻が「次に鳴らすべき時間」を超えるまでループ
    // (フレーム落ちで処理が遅れても、鳴らすべきだった音を再生し追いつくため)
    while current_time >= next_beat_time {
        // メトロノーム音を再生
        commands.spawn(AudioPlayer::new(
            asset_server.load("sounds/maou_se_8bit02.ogg"),
        ));

        // 次の拍に進める
        metronome.next_beat_index += 1;
        // 次の拍の時間を計算
        next_beat_time = metronome.next_beat_index as f32 * metronome.seconds_per_beat;
    }
}

// 絶対時間に基づき、キー入力を判定
fn check_key_press(
    mut commands: Commands,
    input: Res<ButtonInput<KeyCode>>,
    time: Res<Time>,                // SoundTimer の代わりに Time を使用
    metronome: Res<MetronomeState>, // seconds_per_beat を参照
    asset_server: Res<AssetServer>,
) {
    if input.just_pressed(KeyCode::Space) {
        const ALLOWED_ERROR_MARGIN: f32 = 0.05;
        let spb = metronome.seconds_per_beat; // 1.0

        // キーが押された瞬間の「絶対時間」
        let current_time = time.elapsed_secs();

        // 現在時刻に「最も近い」拍の時間を計算
        // (例: 2.98秒 -> 3.0秒, 3.03秒 -> 3.0秒, 3.4秒 -> 3.0秒)
        let nearest_beat_time = (current_time / spb).round() * spb;

        // 差の絶対値
        let delta = (current_time - nearest_beat_time).abs();

        if delta <= ALLOWED_ERROR_MARGIN {
            // 成功音
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
