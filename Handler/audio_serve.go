package handler

import (
	"fmt"
	audioextraction "mytelegrambot/audio_extraction"
)

type VideoData struct {
	Videourl string `json:"videourl"`
}

func (videoData VideoData)ServeAudio() ([]byte, *audioextraction.AudioExtraction) {
	stream, audioExtraction, err := audioextraction.DownloadFromYoutube(videoData.Videourl)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer stream.Close()

	audioData, err := audioExtraction.ExtractAudio(stream)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	fmt.Println("audioExtraction.AudioName:", audioExtraction.AudioName)
	return audioData, audioExtraction
}
