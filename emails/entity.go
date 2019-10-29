package emails

import (
	"math/rand"
	"time"
)

//Email ... email entity
type Email struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func (e Email) Handle() error {
	r := rand.Intn(200)
	time.Sleep(time.Duration(r) * time.Millisecond)
	return nil
}
