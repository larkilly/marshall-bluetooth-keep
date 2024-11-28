package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/generators"
	"github.com/gopxl/beep/v2/speaker"
)

const (
	kSampleRate = 44100

	kDuration = 1 * time.Second
	kInterval = 5 * time.Minute
)

var (
	Version    string
	GoVersion  string
	Commit     string
	CommitTime string
)

func help() {
	fmt.Println("Usage: marshall-bluetooth-keep [version]")
}

func playSilentAudioService(ctx context.Context, duration time.Duration, interval time.Duration) error {
	sr := beep.SampleRate(kSampleRate)
	speaker.Init(sr, sr.N(time.Second/10))

	silence := generators.Silence(-1)

	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return nil

		case <-ticker.C:
			done := make(chan bool)
			speaker.Play(beep.Seq(beep.Take(sr.N(duration), silence), beep.Callback(func() {
				done <- true
			})))
			<-done
		}
	}
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

	// play silent audio for 1 seconds every 5 minutes
	if err := playSilentAudioService(context.Background(), kDuration, kInterval); err != nil {
		panic(err)
	}
}
