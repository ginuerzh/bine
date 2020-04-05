# bine - BINary Editor

## Edit binary file with one command.

### Usage

```
bine -f FILENAME [-s OFFSET] [-n LENGTH] VALUE [VALUE...]
```

### Example

Write the magic number '0xaa55' to an MBR binary file.

```
$ touch mbr.bin
$ bine -f mbr.bin -s 510 -n 2 0x55 0xaa
$ hexdump -C mbr.bin
00000000  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
000001f0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 55 aa  |..............U.|
00000200
```
