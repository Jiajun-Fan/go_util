package main

import (
	"bufio"
	"flag"
	"io"
	"os"
)

func parseArgs() OssConfig {
	config := OssConfig{}
	pt := flag.String("t", "aliyun", "cloud storage provider")
	pk := flag.String("k", "ok", "API key")
	ps := flag.String("s", "ok", "API secret")
	pb := flag.String("b", "ok", "bucket name")
	pf := flag.String("f", "ok", "file name")
	pe := flag.String("e", "ok", "API end point")
	pw := flag.Bool("w", false, "write mode")
	flag.Parse()
	config.Type = *pt
	config.Key = *pk
	config.Secret = *ps
	config.Bucket = *pb
	config.FileName = *pf
	config.EndPoint = *pe
	config.Write = *pw
	return config
}

func pipe(reader io.Reader, writer io.Writer) {
	buff := make([]byte, 4096, 4096)
	for {
		nr, err := reader.Read(buff)
		if nr > 0 {
			offset := 0
			for {
				nw, _ := writer.Write(buff[offset:nr])
				offset += nw
				if nr == offset {
					break
				} else if nr < offset {
					Fatal("pipe error")
				}
			}
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	config := parseArgs()

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	oss := MakeOss(config)

	if config.Write {
		pipe(reader, oss)
	} else {
		pipe(oss, writer)
	}
	oss.Close()
	writer.Flush()
}
