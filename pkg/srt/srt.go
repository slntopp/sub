package srt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	t "time"

	"github.com/slntopp/sub/pkg/core"
	"github.com/slntopp/sub/pkg/vtt"
)

type SRT struct {
	Chunks []core.Chunk
}


type SRTChunk struct {
	seq int
	from t.Time
	to t.Time
	text string
}

const SRT_TIME_FORMAT = "15:04:05,000"

func ParseSRTString(vtt string) (*SRT, error) {
	var r SRT
	chunks := strings.Split(vtt, "\n\n")
	for _, chunk := range chunks {
		if chunk == "" {
			continue
		}
	    res, err := ParseSRTChunk(chunk)
		if err != nil {
			return nil, err
		}
		r.Chunks = append(r.Chunks, res)
	}

	return &r, nil
}

func ParseSRTChunk(chunk string) (*SRTChunk, error) {
	data := strings.Split(chunk, "\n")
	seq, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, errors.New("Can't read Chunk sequence ID")
	}
	time := strings.Split(data[1], " --> ")
	from, err := t.Parse(vtt.VTT_TIME_FORMAT, strings.Replace(time[0], ",", ".", 1))
	if err != nil {
		return nil, errors.New("Can't read Chunk time 'from'")
	}
	to, err := t.Parse(vtt.VTT_TIME_FORMAT, strings.Replace(time[1], ",", ".", 1))
	if err != nil {
		return nil, errors.New("Can't read Chunk time 'to'")
	}
	return &SRTChunk{
		seq: seq,
		from: from,
		to: to,
		text: strings.Join(data[2:], "\n"),
	}, nil
}

func (chunk *SRTChunk) Dump() (r string) {
	r += (strconv.Itoa(chunk.seq) + "\n")
	r += fmt.Sprintf("%s --> %s\n", chunk.from.Format(SRT_TIME_FORMAT), chunk.to.Format(SRT_TIME_FORMAT))
	r += chunk.text
	return r
}

func DumpSRTChunk(chunk SRTChunk) (r string) {
	return chunk.Dump()
}

func (srt *SRT) Dump() (r string) {
	for _, chunk := range srt.Chunks {
		r += chunk.(*SRTChunk).Dump() + "\n\n"
	}
	return r
}

func DumpSRT(vtt SRT) (r string) {
	return vtt.Dump()
}

func (chunk *SRTChunk) Seq() int {
	return chunk.seq
}

func (chunk *SRTChunk) From() t.Time {
	return chunk.from
}

func (chunk *SRTChunk) To() t.Time {
	return chunk.to
}

func (chunk *SRTChunk) Text() string {
	return chunk.text
}