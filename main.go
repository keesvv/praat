package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/diamondburned/arikawa/v3/discord"
	"keesvv.nl/praat/voice"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	channel := flag.String("c", "", "channel ID")
	flag.Parse()

	chanSf, err := discord.ParseSnowflake(*channel)
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	s := voice.NewSession(token)
	if err := s.Open(ctx, discord.ChannelID(chanSf)); err != nil {
		log.Fatalln(err)
	}

	defer s.Close()

	if err := s.Stream(os.Stdin); err != nil &&
		!errors.Is(err, context.Canceled) {
		log.Fatalln(err)
	}
}
