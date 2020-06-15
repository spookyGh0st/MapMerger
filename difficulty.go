package main

import "encoding/json"

// DifficultyJSON is the new beatmap difficulty file
type DifficultyJSON struct {
	Version string `json:"_version"`

	Events    []Event    `json:"_events"`
	Notes     []Note     `json:"_notes"`
	Obstacles []Obstacle `json:"_obstacles"`

	CustomData DifficultyCustomData `json:"_customData,omitempty"`
}

// DifficultyCustomData is the custom data of new beatmap difficulties
type DifficultyCustomData struct {
	BPMChanges []BPMChange `json:"_BPMChanges,omitempty"`
	Bookmarks  []Bookmark  `json:"_bookmarks,omitempty"`

	Time int `json:"_time,omitempty"`
}

// BPMChange MM BPM Change
type BPMChange struct {
	BPM             float64 `json:"_BPM"`
	Time            float64 `json:"_time"`
	BeatsPerBar     int     `json:"_beatsPerBar"`
	MetronomeOffset int     `json:"_metronomeOffset"`
}

// Event Beatmap Event
type Event struct {
	Time       float64         `json:"_time"`
	Type       int             `json:"_type"`
	Value      int             `json:"_value"`
	CustomData EventCustomData `json:"_customData,omitempty"`
}

// EventCustomData Custom Data
type EventCustomData struct {
	Color    []float64 `json:"_color,omitempty"`
	Rotation []float64 `json:"_rotation,omitempty"`
}

// Note Beatmap Note
type Note struct {
	Time         float64        `json:"_time"`
	LineIndex    int            `json:"_lineIndex"`
	LineLayer    int            `json:"_lineLayer"`
	Type         int            `json:"_type"`
	CutDirection int            `json:"_cutDirection"`
	CustomData   NoteCustomData `json:"_customData,omitempty"`
}

// NoteCustomData Custom Data
type NoteCustomData struct {
	Position     []float64 `json:"_position,omitempty"`
	Rotation     []float64 `json:"_rotation,omitempty"`
	CutDirection *float64  `json:"_cutDirection,omitempty"`
	Flip         []float64 `json:"_flip,omitempty"`
	Color        []float64 `json:"_color,omitempty"`
}

// Obstacle Beatmap Obstacle
type Obstacle struct {
	Time       float64            `json:"_time"`
	LineIndex  int                `json:"_lineIndex"`
	Type       int                `json:"_type"`
	Duration   float64            `json:"_duration"`
	Width      int                `json:"_width"`
	CustomData ObstacleCustomData `json:"_customData,omitempty"`
}

// ObstacleCustomData Custom Data
type ObstacleCustomData struct {
	Position      []float64 `json:"_position,omitempty"`
	Scale         []float64 `json:"_scale,omitempty"`
	Rotation      []float64 `json:"_rotation,omitempty"`
	LocalRotation []float64 `json:"_flip,omitempty"`
	Color         []float64 `json:"_color,omitempty"`
}

// Bookmark MM Bookmark
type Bookmark struct {
	Time float64 `json:"_time"`
	Name string  `json:"_name"`
}

// Bytes Convert to byte array
func (i DifficultyJSON) Bytes() ([]byte, error) {
	return json.Marshal(i)
}
