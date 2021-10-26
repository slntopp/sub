package vtt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	t "time"

	"github.com/slntopp/sub/pkg/core"
)

// Format used by VTT 'timestamps'
const VTT_TIME_FORMAT = "15:04:05.000"

// Parses string as VTT and stores into Subtitles
func Parse(r *core.Subtitles, vtt string) (error) {
	chunks := strings.Split(vtt, "\n\n")
	if chunks[0] != "WEBVTT" {
		return errors.New("Not a VTT format")
	}
	chunks = chunks[1:]
	for _, chunk := range chunks {
		if chunk == "" {
			continue
		}
	    res, err := ParseChunk(chunk)
		if err != nil {
			return err
		}
		r.Chunks = append(r.Chunks, *res)
	}

	return nil
}

// Parses single VTT Chunk
func ParseChunk(chunk string) (*core.Chunk, error) {
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
	return &core.Chunk{
		Seq: seq,
		From: from,
		To: to,
		Text: strings.Join(data[2:], "\n"),
	}, nil
}

// Dumps single VTT Chunk
func DumpChunk(chunk core.Chunk) (r string) {
	r += (strconv.Itoa(chunk.Seq) + "\n")
	r += fmt.Sprintf("%s --> %s\n", chunk.From.Format(VTT_TIME_FORMAT), chunk.To.Format(VTT_TIME_FORMAT))
	r += chunk.Text
	return r
}

// Dumps Subtitles into VTT string
func Dump(vtt *core.Subtitles) (r string) {
	r += "WEBVTT\n\n"
	for _, chunk := range vtt.Chunks {
		r += DumpChunk(chunk) + "\n\n"
	}
	return r
}