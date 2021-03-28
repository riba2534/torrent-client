package main

import (
	"fmt"
	"log"
	"os"

	"github.com/riba2534/torrent-client/torrentfile"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("param error, please check path!")
		os.Exit(0)
	}
	inPath := os.Args[1]
	outPath := os.Args[2]
	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}
	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
