# Go Subtitles

Go Packages for subtitles parsing and converting

## Contents

1. `sbtl` CLI tool
2. `srt` Go package
3. `vtt` Go package

### `sbtl` CLI tool

### Download

1. Download Binary matching to your platform from Releases page
2. `mv sbtl-<your-os>-<your-arch> /bin/sbtl`

### Use

`sbtl` support only one action right now - convert.
To use it, you can run:

```shell
sbtl /path/to/input.srt /path/to/output.vtt
```

### Build

> Golang >1.17 must be installed
> If you want to compile for another platfrom/os, read about `GOOS` and `GOARCH` env variables

1. Clone this repo
2. `cd repo/location`
3. `go build`

Supported formats are:

* SRT
* VTT

### `core` Go Package

### `srt` Go package

### `vtt` Go package
