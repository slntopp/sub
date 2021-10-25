package vtt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	t "time"
)

type VTT map[int]VTTChunk

type VTTChunk struct {
	seq int
	from t.Time
	to t.Time
	text string
}

const VTT_TIME_FORMAT = "15:04:05.000"

func ParseVTTString(vtt string) (*VTT, error) {
	r := make(VTT)
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
		r[res.seq] = *res
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