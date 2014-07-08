package Model

import (
	"time"
)

type Message struct {
	Message, Nick,Channel,Server string
	Time time.Time
}
