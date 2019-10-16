package main

import (
	"github.com/farzaa/RedditPlaylist/reddit"
	log "github.com/sirupsen/logrus"
	"os"
)

func initLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	initLogging()
	log.Debug("Starting up...")
	reddit.GetTopPosts("hiphopheads")
}
