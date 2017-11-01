package main

import (
    "bufio"
    "flag"
    "io"
    "os"
)

func parseArgs() EncryptConfig {
    config := EncryptConfig{}
    pt := flag.String("type", "aes", "encrypt/decrypt type")
    pk := flag.String("key", "ok", "encrypt/decrypt key")
    pe := flag.Bool("dec", false, "decrypt mode")
    flag.Parse()
    config.Type = *pt 
    config.Key = *pk 
    config.Enc = !*pe
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
    
    if config.Enc {
        enc := MakeEncryptor(config, writer)

        pipe(reader, enc)

        enc.Close()
    } else {
        dec := MakeDecryptor(config, reader)

        pipe(dec, writer)

        dec.Close()
    }
    writer.Flush()
}
