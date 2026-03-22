package tracker

import (
	"log"
	"ppclimb-client/internal/models"
)

type Tracker struct {
	messages <-chan models.OsuMessage
}

func NewTracker(messages <-chan models.OsuMessage) *Tracker {
	return &Tracker{messages: messages}
}

func (t *Tracker) Run() {
	var prev models.OsuMessage

	for msg := range t.messages {
		if prev.State.Name == "play" && msg.State.Name == "resultScreen" {
			if msg.ResultsScreen.PP.Current > 0 {
				go t.submitScore(msg)
				prev = msg
			}
			continue
		}

		prev = msg
	}
}

func (t *Tracker) submitScore(msg models.OsuMessage) {
	log.Printf("Submitting score: %+v", msg)
}
