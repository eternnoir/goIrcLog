package Model

import (
	"time"
)

type Message struct {
	Message, Nick,Channel string
	Time time.Time
}
