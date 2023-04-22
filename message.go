package app

import "time"

type Message struct {
	Sender     string
	SenderRole string
	Content    string
	Recipient  string 
	Time       time.Time
	IsRead     bool
}
