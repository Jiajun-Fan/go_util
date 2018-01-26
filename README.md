# go_util

A few utilities make life easier.

enc
=============================
Encrypt/decrypt data read from STDIN and prints the result into STDOUT.
```
Usage of ./enc:
  -d    decrypt mode
  -k string
        encrypt/decrypt key (default "ok")
  -t string
        encrypt/decrypt type (default "aes")
```

oss
============================
Read data from STDIN and put it to cloud storage.
Get data from cloud storage and print it into STDOUT.
```
Usage of ./oss:
  -b string
        bucket name (default "ok")
  -e string
        API end point (default "ok")
  -f string
        file name (default "ok")
  -k string
        API key (default "ok")
  -s string
        API secret (default "ok")
  -t string
        cloud storage provider (default "aliyun")
  -w    write mode
```

icrop
=============================
Read input image and crop it using specified rect and write it output file.
```
Usage of /home/ubuntu/icrop:
  -a int
        X min
  -b int
        Y min
  -c int
        X max (default 1)
  -d int
        Y max (default 1)
  -i string
        input file name (default "ok")
  -o string
        output file name (default "ok_output")
```

iresize
=============================
resize a image
```
Usage of /Users/fanjiajun/iresize:
  -h int
    	height (default 256)
  -i string
    	input file name (default "ok")
  -o string
    	output file name (default "ok_output")
  -w int
    	width (default 256)
```

irotate
=============================
rotate a image
```
Usage of ./irotate:
  -i string
    	input file name (default "ok")
  -o string
    	output file name (default "ok_output")
  -p float
    	angle (default 1)
  -s int
    	size (default 1)
  -x int
    	center x (default 1)
  -y int
    	center y (default 1)
```
