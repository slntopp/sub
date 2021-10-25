package vtt_test

import (
	"testing"

	"github.com/slntopp/sub/pkg/vtt"
)

func TestParseVTTChunk(t *testing.T) {
	_, err := vtt.ParseVTTChunk(VTT_TEST_CHUNK)
	if err != nil {
		t.Errorf("Expected error to be nil, got \"%v\" instead", err)
	}

	_, err = vtt.ParseVTTChunk(VTT_TEST_CHUNK_ERROR_SEQ)
	if err == nil {
		t.Error("Expected to be error")
	} else if err.Error() != "Can't read Chunk sequence ID" {
		t.Errorf("Expected error to be \"Can't read Chunk sequence ID\", got \"%v\" instead", err)
	}

	_, err = vtt.ParseVTTChunk(VTT_TEST_CHUNK_ERROR_NO_SEQ)
	if err == nil {
		t.Error("Expected to be error")
	} else if err.Error() != "Can't read Chunk sequence ID" {
		t.Errorf("Expected error to be \"Can't read Chunk sequence ID\", got \"%v\" instead", err)
	}

	_, err = vtt.ParseVTTChunk(VTT_TEST_CHUNK_ERROR_TIME_FROM)
	if err == nil {
		t.Error("Expected to be error")
	} else if err.Error() != "Can't read Chunk time 'from'" {
		t.Errorf("Expected error to be \"Can't read Chunk time 'from'\", got \"%v\" instead", err)
	}

	_, err = vtt.ParseVTTChunk(VTT_TEST_CHUNK_ERROR_TIME_TO)
	if err == nil {
		t.Error("Expected to be error")
	} else if err.Error() != "Can't read Chunk time 'to'" {
		t.Errorf("Expected error to be \"Can't read Chunk time 'to'\", got \"%v\" instead", err)
	}
}

func TestDumpVTTChunk(t *testing.T) {
	chunk, _ := vtt.ParseVTTChunk(VTT_TEST_CHUNK)
	r := vtt.DumpVTTChunk(*chunk) == VTT_TEST_CHUNK
	if !r {
		t.Error("Expected parsed and dumped chunk to be equal original chunk")
	}
}

func TestParseVTTString(t *testing.T) {
	_, err := vtt.ParseVTTString(VTT_TEST_DATA)
	if err != nil {
		t.Errorf("Expected error to be nil, got: %v", err)
	}

	_, err = vtt.ParseVTTString(VTT_TEST_DATA_CORR_TIME_CHUNK)
	if err == nil {
		t.Error("Expected to be error")
	} else if err.Error() != "Can't read Chunk time 'from'" {
		t.Errorf("Expected error to be \"Can't read Chunk time 'from'\", got \"%v\" instead", err)
	}

	_, err = vtt.ParseVTTString(VTT_TEST_CHUNK)
	if err == nil {
		t.Error("Expected to be error")
	} else if err.Error() != "Not a VTT format" {
		t.Errorf("Expected error to be \"Not a VTT format\", got \"%v\" instead", err)
	}
}

const (
	VTT_TEST_CHUNK = `1
00:01:32.234 --> 00:01:34.754
Radio Moscow.
Director Andreyev. What is it?`
	VTT_TEST_CHUNK_ERROR_TIME_FROM = `1
00:01:32,234 --> 00:01:34.754
Radio Moscow.
Director Andreyev. What is it?`
	VTT_TEST_CHUNK_ERROR_TIME_TO = `1
00:01:32.234 --> 00:01:34,754
Radio Moscow.
Director Andreyev. What is it?`
	VTT_TEST_CHUNK_ERROR_SEQ = `A
00:01:32.234 --> 00:01:34.754
Radio Moscow.
Director Andreyev. What is it?`
	VTT_TEST_CHUNK_ERROR_NO_SEQ = `00:01:32.234 --> 00:01:34.754
Radio Moscow.
Director Andreyev. What is it?`

	VTT_TEST_DATA = `WEBVTT

1
00:01:32.234 --> 00:01:34.754
Radio Moscow.
Director Andreyev. What is it?

2
00:01:38.114 --> 00:01:39.353
Seventeen minutes.

3
00:01:40.154 --> 00:01:43.074
Yes, of course I can ring back
in 17 minutes.

4
00:01:44.034 --> 00:01:45.034
Yeah.

5
00:01:46.073 --> 00:01:48.193
Mm-hm. Yes, I'm writing it down.

6
00:01:49.353 --> 00:01:50.952
I can't get the...

7
00:01:50.954 --> 00:01:53.074
One, five...

8
00:01:54.594 --> 00:01:56.633
Sorry? Was that a nine, as in "fine"?

9
00:01:57.113 --> 00:01:59.474
Or... or another five as in, um...

10
00:02:00.034 --> 00:02:01.633
- "Hive."
- "hive"?

11
00:02:02.834 --> 00:02:04.994
Hello? Hello?

12
00:02:10.474 --> 00:02:12.791
- Hive?
- Who was it?

13
00:02:12.793 --> 00:02:14.992
The Secretariat
of the General Secretariat.

`

VTT_TEST_DATA_CORR_TIME_CHUNK = `WEBVTT

1
00:01:32,234 --> 00:01:34.754
Radio Moscow.
Director Andreyev. What is it?

2
00:01:38.114 --> 00:01:39.353
Seventeen minutes.

`
)