package main

import (
	"encoding/json"
	"html"
	"log"
	"os/exec"
	"strings"

	"github.com/grokify/html-strip-tags-go"
)

type DunstctlHistoryData struct {
	Type string `json:"type"`
	Data [][]DunstNotification
}

type DunstNotification struct {
	Body struct {
		Type string `json:"type"`
		Data string `json:"data"`
	} `json:"body"`
	Message struct {
		Type string `json:"type"`
		Data string `json:"data"`
	} `json:"message"`
	Summary struct {
		Type string `json:"type"`
		Data string `json:"data"`
	} `json:"summary"`
	Appname struct {
		Type string `json:"type"`
		Data string `json:"data"`
	} `json:"appname"`
	Category struct {
		Type string `json:"type"`
		Data string `json:"data"`
	} `json:"category"`
	DefaultActionName struct {
		Type string `json:"type"`
		Data string `json:"data"`
	} `json:"default_action_name"`
	IconPath struct {
		Type string `json:"type"`
		Data string `json:"data"`
	} `json:"icon_path"`
	ID struct {
		Type string `json:"type"`
		Data int    `json:"data"`
	} `json:"id"`
	Timestamp struct {
		Type string `json:"type"`
		Data int64  `json:"data"`
	} `json:"timestamp"`
	Timeout struct {
		Type string `json:"type"`
		Data int64  `json:"data"`
	} `json:"timeout"`
	Progress struct {
		Type string `json:"type"`
		Data int    `json:"data"`
	} `json:"progress"`
}

func getDunstctlHistoryJSON() string {
	cmd := exec.Command("dunstctl", "history")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(output)
}

func getDunstctlHistory() []DunstNotification {
	var dunstctlHistory DunstctlHistoryData
	err := json.Unmarshal([]byte(getDunstctlHistoryJSON()), &dunstctlHistory)
	if err != nil {
		log.Fatal(err)
	}
	result := dunstctlHistory.Data[0]
	for i := range result {
		result[i].Body.Data = strings.ReplaceAll(
			html.UnescapeString(
				strip.StripTags(result[i].Body.Data),
			), "\n", " ",
		)
	}
	return result
}
