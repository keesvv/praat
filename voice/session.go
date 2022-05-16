package voice

import (
	"context"
	"fmt"
	"io"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/voice"
	"github.com/diamondburned/oggreader"
)

type Session struct {
	Token string
	state *state.State
	vsess *voice.Session
}

func NewSession(token string) *Session {
	return &Session{Token: token}
}

func (s *Session) Open(ctx context.Context, id discord.ChannelID) error {
	st := state.New(fmt.Sprint("Bot ", s.Token))
	voice.AddIntents(st)

	if err := st.Open(ctx); err != nil {
		return err
	}

	v, err := voice.NewSession(st)
	if err != nil {
		return err
	}

	if err := v.JoinChannelAndSpeak(st.Context(), id, false, true); err != nil {
		return err
	}

	s.state = st
	s.vsess = v
	return nil
}

func (s *Session) Stream(r io.Reader) error {
	return oggreader.DecodeBuffered(s.vsess, r)
}

func (s *Session) Close() error {
	s.vsess.Leave(s.state.Context())
	return s.state.Close()
}
