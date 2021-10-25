package core

import (
	t "time"
)

type Chunk interface {
	Seq() int;
	From() t.Time;
	To() t.Time;
	Text() string;
}

type Subtitles interface {
	Dump() string;
}