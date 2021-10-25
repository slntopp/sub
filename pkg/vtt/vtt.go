package vtt

import (
	"errors"
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
	from, err := t.Parse("15:04:05.000", time[0])
	if err != nil {
		return nil, errors.New("Can't read Chunk time 'from'")
	}
	to, err := t.Parse("15:04:05.000", time[1])
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