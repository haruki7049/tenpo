package assets_test

import (
	"testing"

	"github.com/haruki7049/tenpo/assets"
)

func TestAssetsEmbedded(t *testing.T) {
	f, err := assets.Assets.Open("audio/defaultSource.wav")
	if err != nil {
		t.Fatalf("expected Wav file to be embedded, got error: %v", err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		t.Fatalf("expected to stat font file, got error: %v", err)
	}
	if stat.Size() == 0 {
		t.Errorf("expected Wav file size to be greater than 0, got 0")
	}
}

func TestAssetsInvalidFile(t *testing.T) {
	_, err := assets.Assets.Open("nonexistent_path/file.txt")
	if err == nil {
		t.Errorf("expected error when opening non-existent file, got nil")
	}
}
