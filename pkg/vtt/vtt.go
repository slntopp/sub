package vtt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	t "time"

	"github.com/slntopp/sub/pkg/core"
)

type VTT struct {
	Chunks []core.Chunk
}

type VTTChunk struct {
	seq int
	from t.Time
	to t.Time
	text string
}

const VTT_TIME_FORMAT = "15:04:05.000"

func ParseVTTString(vtt string) (*VTT, error) {
	var r VTT
	chunks := strings.Split(vtt, "\n\n")
	if chunks[0] != "WEBVTT" {
		return nil, errors.New("Not a VTT format")
	}
	chunks = chunks[1:]
	for _, chunk := range chunks {
		if chunk == "" {
			continue
		}
	    res, err := ParseVTTChunk(chunk)
		if err != nil {
			return nil, err
		}
		r.Chunks = append(r.Chunks, res)
	}

	return &r, nil
}

func ParseVTTChunk(chunk string) (*VTTChunk, error) {
	data := strings.Split(chunk, "\n")
	seq, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, errors.New("Can't read Chunk sequence ID")
	}
	time := strings.Split(data[1], " --> ")
	from, err := t.Parse(VTT_TIME_FORMAT, time[0])
	if err != nil {
		return nil, errors.New("Can't read Chunk time 'from'")
	}
	to, err := t.Parse(VTT_TIME_FORMAT, time[1])
	if err != nil {
		return nil, errors.New("Can't read Chunk time 'to'")
	}
	return &VTTChunk{
		seq: seq,
		from: from,
		to: to,
		text: strings.Join(data[2:], "\n"),
	}, nil
}

func (chunk *VTTChunk) Dump() (r string) {
	r += (strconv.Itoa(chunk.seq) + "\n")
	r += fmt.Sprintf("%s --> %s\n", chunk.from.Format(VTT_TIME_FORMAT), chunk.to.Format(VTT_TIME_FORMAT))
	r += chunk.text
	return r
}

func DumpVTTChunk(chunk VTTChunk) (r string) {
	return chunk.Dump()
}

func (vtt *VTT) Dump() (r string) {
	r += "WEBVTT\n\n"
	for _, chunk := range vtt.Chunks {
		r += chunk.(*VTTChunk).Dump() + "\n\n"
	}
	return r
}

func DumpVTT(vtt VTT) (r string) {
	return vtt.Dump()
}

func (chunk *VTTChunk) Seq() int {
	return chunk.seq
}

func (chunk *VTTChunk) From() t.Time {
	return chunk.from
}

func (chunk *VTTChunk) To() t.Time {
	return chunk.to
}

func (chunk *VTTChunk) Text() string {
	return chunk.text
}