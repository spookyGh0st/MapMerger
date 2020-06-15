package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	info := parseInfo()
	var final DifficultyBeatmap
	for _, diffSet := range info.DifficultyBeatmapSets {
		for _, diff := range diffSet.DifficultyBeatmaps {
			str := strings.ToLower(diff.CustomData.DifficultyLabel)
			if str == "final" || str == "merged" {
				final = diff
			}
		}
	}

	finalDiff := final.DiffJSON
	println("Appending to", final.BeatmapFilename)
	for _, diffSet := range info.DifficultyBeatmapSets {
		for _, diff := range diffSet.DifficultyBeatmaps {
			println("checking", diff.BeatmapFilename)
			d := diff.DiffJSON
			if diff.isLabel("notes") {
				finalDiff.Notes = append(finalDiff.Notes, d.Notes...)
			}
			if diff.isLabel("bombs") {
				finalDiff.Notes = append(finalDiff.Notes, d.Notes...)
			}
			if diff.isLabel("lights") {
				finalDiff.Events = append(finalDiff.Events, d.Events...)
			}
			if diff.isLabel("walls") {
				finalDiff.Obstacles = append(finalDiff.Obstacles, d.Obstacles...)
			}
			if diff.isLabel("obstacles") {
				finalDiff.Obstacles = append(finalDiff.Obstacles, d.Obstacles...)
			}
		}
	}
	final.backup()
	str, _ := json.Marshal(finalDiff)
	ioutil.WriteFile(final.BeatmapFilename, str, 0666)
}

func (d DifficultyBeatmap) backup() {
	oldName := d.BeatmapFilename
	newName := oldName + ".bak"
	err := os.Link(oldName, newName)
	if err != nil {
		println("Failed to create Backup, Hard links may not be supported by your System")
	}

}

func (d DifficultyBeatmap) isLabel(s string) bool {
	label := strings.ToLower(d.CustomData.DifficultyLabel)
	str := strings.ToLower(s)
	return strings.Contains(label, str)
}

func parseInfo() InfoJSON {
	info := InfoJSON{}
	infoStr, err := ioutil.ReadFile("info.dat")
	if err != nil {
		println("Failed to read info.dat. Is this Program placed in you map-folder?")
		os.Exit(1)
	}
	if json.Unmarshal(infoStr, &info) != nil {
		println("Failed to parse info json. Is it valid json?")
		os.Exit(1)
	}

	for _, diffSet := range info.DifficultyBeatmapSets {
		for _, diff := range diffSet.DifficultyBeatmaps {
			diffStr, err := ioutil.ReadFile(diff.BeatmapFilename)
			if err != nil {
				println("Failed to read Difficulty:", diff.BeatmapFilename)
				os.Exit(1)
			}
			diff.DiffJSON = &DifficultyJSON{}
			err = json.Unmarshal(diffStr, diff.DiffJSON)
			if err != nil {
				println("Failed to parse difficulty. Is it valid json?: ", diff.BeatmapFilename)
				println(err.Error())
				os.Exit(1)
			}
		}
	}
	return info
}
