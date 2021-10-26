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

const SRT_TIME_FORMAT = "15:04:05,000"

func Parse(r *core.Subtitles, srt string) (error) {
	chunks := strings.Split(srt, "\n\n")
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

func ParseChunk(chunk string) (*core.Chunk, error) {
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
	return &core.Chunk{
		Seq: seq,
		From: from,
		To: to,
		Text: strings.Join(data[2:], "\n"),
	}, nil
}

func DumpChunk(chunk core.Chunk) (r string) {
	r += (strconv.Itoa(chunk.Seq) + "\n")
	r += fmt.Sprintf("%s --> %s\n", chunk.From.Format(SRT_TIME_FORMAT), chunk.To.Format(SRT_TIME_FORMAT))
	r += chunk.Text
	return r
}

func Dump(srt *core.Subtitles) (r string) {
	for _, chunk := range srt.Chunks {
		r += DumpChunk(chunk) + "\n\n"
	}
	return r
}
