package core

import (
	t "time"
)

type Chunk struct {
	Seq int
	From t.Time
	To t.Time
	Text string
}

type Subtitles struct {
	Chunks []Chunk
}