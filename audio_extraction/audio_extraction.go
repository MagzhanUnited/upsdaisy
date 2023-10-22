package audioextraction

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/kkdai/youtube/v2"
)

type AudioExtraction struct {
	// Videopath string `json:"videopath"`
	AudioName    string `json:"audioname"`
	ThumbnailURL string `json:"thumbnailurl"`
}

func (avp *AudioExtraction) ExtractAudio(videoStream io.Reader) ([]byte, error) {
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-vn", "-acodec", "mp3", "-f", "mp3", "pipe:1")
	cmd.Stdin = videoStream
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error extracting audio:", err)
		fmt.Println("FFmpeg error message:", stderr.String())
		return nil, err
	}
	return stdout.Bytes(), nil
}

func DownloadFromYoutube(videoURL string) (io.ReadCloser, *AudioExtraction, error) {
	// Replace the YouTube video URL with the one you want to download
	client := youtube.Client{}

	video, err := client.GetVideo(videoURL)
	if err != nil {
		log.Fatalf("Error getting video info: %v", err)
	}
	var thumbnailURL string
	if len(video.Thumbnails) > 0 {
		thumbnailURL = video.Thumbnails[0].URL
	}
	stream, _, err := client.GetStream(video, &video.Formats.Quality("tiny")[0])
	if err != nil {
		log.Fatalf("Error getting video stream: %v", err)
	}

	return stream, &AudioExtraction{AudioName: video.Title, ThumbnailURL: thumbnailURL}, nil
}
