package vtt_test

import (
	"testing"

	"github.com/slntopp/sub/pkg/vtt"
)

func TestParseVTTChunk(t *testing.T) {
	r, err := vtt.ParseVTTChunk(VTT_TEST_CHUNK)
	if err != nil {
		t.Errorf("Expected error to be nil, got: %v", err)
	}
	t.Log("Result", r)
}

func TestParseVTTString(t *testing.T) {
	r, err := vtt.ParseVTTString(VTT_TEST_DATA)
	if err != nil {
		t.Errorf("Expected error to be nil, got: %v", err)
	}
	t.Log("Result", r)
}

const (
	VTT_TEST_CHUNK = `1
00:01:32.234 --> 00:01:34.754
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
)