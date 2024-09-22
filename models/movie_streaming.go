package models

type MovieStreaming struct {
	ID          int
	MovieID     int
	StreamingID int
	MovieURL    string
	Streaming   Streaming
}
