package main

import (
	"encoding/json"
)

// InfoJSON New Info JSON
type InfoJSON struct {
	Version string `json:"_version"`

	SongName        string `json:"_songName"`
	SongSubName     string `json:"_songSubName"`
	SongAuthorName  string `json:"_songAuthorName"`
	LevelAuthorName string `json:"_levelAuthorName"`

	BeatsPerMinute float64 `json:"_beatsPerMinute"`
	SongTimeOffset float64 `json:"_songTimeOffset"`
	Shuffle        float64 `json:"_shuffle"`
	ShufflePeriod  float64 `json:"_shufflePeriod"`

	PreviewStartTime float64 `json:"_previewStartTime"`
	PreviewDuration  float64 `json:"_previewDuration"`

	OldSongFilename    string `json:"-"`
	SongFilename       string `json:"_songFilename"`
	CoverImageFilename string `json:"_coverImageFilename"`

	EnvironmentName string `json:"_environmentName"`

	CustomData InfoCustomData `json:"_customData,omitempty"`

	DifficultyBeatmapSets []DifficultyBeatmapSet `json:"_difficultyBeatmapSets"`

	Hash string `json:"-"`
}

// InfoCustomData Custom JSON Data for root info.json
type InfoCustomData struct {
	Contributors []Contributor `json:"_contributors,omitempty"`

	CustomEnvironment     string `json:"_customEnvironment,omitempty"`
	CustomEnvironmentHash string `json:"_customEnvironmentHash,omitempty"`
}

// Contributor New Info JSON Contributors
type Contributor struct {
	Role     string `json:"_role"`
	Name     string `json:"_name"`
	IconPath string `json:"_iconPath"`
}

// BeatmapColor Beatmap Lighting Color
type BeatmapColor struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
}

// DifficultyBeatmap Beatmap Difficulty Info
type DifficultyBeatmap struct {
	Difficulty      string `json:"_difficulty"`
	DifficultyRank  int    `json:"_difficultyRank"`
	BeatmapFilename string `json:"_beatmapFilename"`

	NoteJumpMovementSpeed   float64 `json:"_noteJumpMovementSpeed"`
	NoteJumpStartBeatOffset float64 `json:"_noteJumpStartBeatOffset"`

	CustomData BeatmapCustomData `json:"_customData,omitempty"`

	DiffJSON *DifficultyJSON `json:"-"`
}

// BeatmapCustomData Custom JSON Data for a DifficultyBeatmap
type BeatmapCustomData struct {
	DifficultyLabel string `json:"_difficultyLabel,omitempty"`

	EditorOffset    int `json:"_editorOffset,omitempty"`
	EditorOldOffset int `json:"_editorOldOffset,omitempty"`

	ColorLeft  *BeatmapColor `json:"_colorLeft,omitempty"`
	ColorRight *BeatmapColor `json:"_colorRight,omitempty"`

	Warnings     []string `json:"_warnings,omitempty"`
	Information  []string `json:"_information,omitempty"`
	Suggestions  []string `json:"_suggestions,omitempty"`
	Requirements []string `json:"_requirements,omitempty"`
}

// DifficultyBeatmapSet Set of DifficultyBeatmap structs
type DifficultyBeatmapSet struct {
	BeatmapCharacteristicName string              `json:"_beatmapCharacteristicName"`
	DifficultyBeatmaps        []DifficultyBeatmap `json:"_difficultyBeatmaps"`
}

// Bytes Convert to byte array
func (i InfoJSON) Bytes() ([]byte, error) {
	return json.Marshal(i)
}
