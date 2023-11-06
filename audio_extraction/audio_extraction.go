package audioextraction

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"

	"github.com/kkdai/youtube/v2"
)

type AudioExtraction struct {
	// Videopath string `json:"videopath"`
	AudioName    string `json:"audioname"`
	ThumbnailURL string `json:"thumbnailurl"`
}

func (avp *AudioExtraction) ExtractAudio(videoStream io.Reader) ([]byte, error) {
	ValidateInputStream(videoStream)
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-vn", "-acodec", "mp3", "-f", "mp3", "pipe:1")
	cmd.Stdin = videoStream
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	fmt.Printf("FFmpeg command: %s\n", strings.Join(cmd.Args, " "))

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error extracting audio:", err)
		fmt.Println("FFmpeg error message:", stderr.String())
		return nil, err
	}

	return stdout.Bytes(), nil
}

func ValidateInputStream(videoStream io.Reader) error {
	// Create a temporary buffer to store a portion of the input data for inspection.
	bufferSize := 4096 // You can adjust the buffer size as needed.
	buffer := make([]byte, bufferSize)

	// Read a portion of the input stream.
	n, err := videoStream.Read(buffer)
	if err != nil && err != io.EOF {
		return err
	}

	if n == 0 {
		log.Println("Input stream is empty or at EOF.")
		return nil
	}

	// Inspect the read data to determine its content or format.
	// You can check for specific patterns, headers, or signatures here.

	// For example, you can check if the data starts with a common video format header.
	if string(buffer[:4]) == "RIFF" {
		log.Println("Input stream appears to be in AVI format.")
	} else if string(buffer[0:3]) == "ID3" {
		log.Println("Input stream appears to be in MP3 format.")
	} else {
		log.Println("Input stream format is not recognized:", string(buffer[:4]))
	}

	return nil
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
