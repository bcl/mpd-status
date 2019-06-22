package main

import (
	"flag"
	"fmt"
	"github.com/fhs/gompd/mpd"
	"log"
	"strings"
	"time"
	"unicode/utf8"
)

/* commandline flags */
type cmdlineArgs struct {
	Volume  bool // Show volume percentage 0-100
	Elapsed bool // Show the duration:elapsed time
	Width   int  // Maximum width
	Debug   bool // Output debugging info
}

/* commandline defaults */
var cfg = cmdlineArgs{
	Volume:  false,
	Elapsed: false,
	Width:   60,
	Debug:   false,
}

/* parseArgs handles parsing the cmdline args and setting values in the global cfg struct */
func parseArgs() {
	flag.BoolVar(&cfg.Volume, "volume", cfg.Volume, "Include the volume percentage 0-100")
	flag.BoolVar(&cfg.Elapsed, "elapsed", cfg.Elapsed, "Include the duration:elapsed time")
	flag.IntVar(&cfg.Width, "width", cfg.Width, "Maximum width of output")
	flag.BoolVar(&cfg.Debug, "debug", cfg.Debug, "Output debug information")

	flag.Parse()
}

func main() {
	parseArgs()

	// Connect to MPD server
	conn, err := mpd.Dial("tcp", "localhost:6600")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// Get the status and current song
	status, err := conn.Status()
	if err != nil {
		log.Fatalln(err)
	}
	song, err := conn.CurrentSong()
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Printf("STATUS- %s\n", status)
	// fmt.Printf("SONG- %s\n", song)

	// Build the optional parts of the output
	var optional strings.Builder
	if cfg.Volume {
		fmt.Fprintf(&optional, " %s%%", status["volume"])
	}
	if cfg.Elapsed {
		duration, _ := time.ParseDuration(status["duration"] + "s")
		elapsed, _ := time.ParseDuration(status["elapsed"] + "s")
		fmt.Fprintf(&optional, " %s/%s", elapsed.Truncate(time.Second), duration.Truncate(time.Second))
	}

	// Build the final output string
	var s strings.Builder
	switch status["state"] {
	case "play":
		s.WriteString("▶ ")
	case "stop":
		s.WriteString("◼ ")
	case "pause":
		s.WriteString("‖ ")
	default:
		s.WriteString("  ")
	}

	// Build the Artist + title part (do I want to make artist optional? album?)
	songStr := fmt.Sprintf("%s - %s", song["Artist"], song["Title"])

	// Calculate how much title to trim
	trim := cfg.Width - utf8.RuneCountInString(optional.String()) - utf8.RuneCountInString(s.String())
	trim = utf8.RuneCountInString(songStr) - trim
	if trim < 0 {
		trim = 0
	} else if trim > utf8.RuneCountInString(songStr) {
		trim = utf8.RuneCountInString(songStr)
	}

	// Trim the title so that it will fit into the width
	fmt.Printf("%s%s%s", s.String(), songStr[trim:], optional.String())
}
