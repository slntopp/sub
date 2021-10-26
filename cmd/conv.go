/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/slntopp/sub/pkg/core"
	"github.com/slntopp/sub/pkg/srt"
	"github.com/slntopp/sub/pkg/vtt"
	"github.com/spf13/cobra"
)

var TYPES = map[string]bool {
	".srt": true,
	".vtt": true,
}

// convCmd represents the conv command
var convCmd = &cobra.Command{
	Use:   "conv",
	Short: "Convert Subtitles between formats",
	Long: ``,
	Args: func(cmd *cobra.Command, args []string) (error) {
		if len(args) < 2 {
			return errors.New("Not enough arguments given")
		}

		info, err := os.Stat(args[0]);
		if errors.Is(err, os.ErrNotExist) {
			return errors.New(fmt.Sprintf("File %s doesn't exist", args[0]))
		}
		if info.IsDir() {
			return errors.New("Is a directory")
		}
		if _, err := os.Stat(args[1]); !errors.Is(err, os.ErrNotExist) {
			return errors.New(fmt.Sprintf("File %s already exists", args[1]))
		}

		if !TYPES[filepath.Ext(args[0])] {
			return errors.New(fmt.Sprintf("Type %s isn't supported", filepath.Ext(args[0])))
		}

		if !TYPES[filepath.Ext(args[1])] {
			return errors.New(fmt.Sprintf("Type %s isn't supported", filepath.Ext(args[1])))
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log := log.Default()
		log.Println("conv called")
		log.Println(args)

		data, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatalf("Unexpected error while reading file %s:\n\t%s", args[0], err)
		}
		var reader core.Subtitles
		switch filepath.Ext(args[0]) {
		case ".srt":
			err = srt.Parse(&reader, string(data))
		case ".vtt":
			err = vtt.Parse(&reader, string(data))
		default:
			log.Fatalf("Unsupported subtitles format... anyhow...")
		}
		if err != nil {
			log.Fatalf("Unexpected error while parsing file:\n\t%s", err)
		}

		var result string
		switch filepath.Ext(args[1]) {
		case ".srt":
			result = srt.Dump(&reader)
		case ".vtt":
			result = vtt.Dump(&reader)
		default:
			log.Fatalf("Unsupported subtitles format... anyhow...")
		}

		err = ioutil.WriteFile(args[1], []byte(result), 0644)
		if err != nil {
			log.Fatalf("Unexpected error while writing file:\n\t%s", err)
		}
		fmt.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(convCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
