package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
)

const (
	kSampleRate = 44100
	kChannels   = 2
	kBitDepth   = 16
)

var (
	Version    string
	GoVersion  string
	Commit     string
	CommitTime string
)

func playSilentAudio(ctx context.Context, duration time.Duration, interval time.Duration) error {
	// generate silent audio
	silentData := make([]byte, kSampleRate*kChannels*kBitDepth/8*int(duration.Seconds()))

	op := &oto.NewContextOptions{
		SampleRate:   kSampleRate,
		ChannelCount: kChannels,
		Format:       oto.FormatSignedInt16LE,
	}

	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		return err
	}

	// wait hardware audio devices to be ready
	<-readyChan

	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return nil

		case <-ticker.C:
			player := otoCtx.NewPlayer(bytes.NewReader(silentData))
			player.Play()
			if err := player.Close(); err != nil {
				return err
			}
		}
	}
}

func help() {
	fmt.Println("Usage: marshall-bluetooth-keep [version]")
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "version" {
			fmt.Println("Version:", Version)
			fmt.Println("Go Version:", GoVersion)
			fmt.Println("Commit:", Commit)
			fmt.Println("Commit Time:", CommitTime)
			fmt.Println("")

			help()

			return
		} else {
			fmt.Println("Unknown command:", os.Args[1])

			help()

			return
		}
	}

	// play silent audio for 3 seconds every 5 minutes
	playSilentAudio(context.Background(), 3*time.Second, 5*time.Minute)
}
