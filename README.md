# ggd - xxd in pure go

`ggd` is the rewrite of the `xxd` binary in pure go using **only the standard library**. Contains *NSFW* feature for making the hex dump filled with rainbow colors. Streams over stdin and buffers output reducing syscalls at most.

```bash
~/ $ cat loremipsum.txt | ./ggd
00000000 (0000)  |  4c6f 7265 6d20 6970 7375 6d20 646f 6c6f  |  Lorem ipsum dolo
00000010 (0016)  |  7220 7369 7420 616d 6574 2c20 636f 6e73  |  r sit amet, cons
00000020 (0032)  |  6563 7465 7475 7220 6164 6970 6973 6369  |  ectetur adipisci
00000030 (0048)  |  6e67 2065 6c69 742c 2073 6564 2064 6f20  |  ng elit, sed do 
00000040 (0064)  |  6569 7573 6d6f 6420 7465 6d70 6f72 2069  |  eiusmod tempor i
00000050 (0080)  |  6e63 6964 6964 756e 7420 7574 206c 6162  |  ncididunt ut lab
00000060 (0096)  |  6f72 6520 6574 2064 6f6c 6f72 6520 6d61  |  ore et dolore ma
00000070 (0112)  |  676e 6120 616c 6971 7561 2e20 5574 2065  |  gna aliqua. Ut e
00000080 (0128)  |  6e69 6d20 6164 206d 696e 696d 2076 656e  |  nim ad minim ven
00000090 (0144)  |  6961 6d2c 2071 7569 7320 6e6f 7374 7275  |  iam, quis nostru
```

```bash
~/coding/ggd $ ./ggd -h
Usage of ./ggd:
  -bf int
        Buffer (MiB) size to assignate to the buffered reader. Default: 1 MiB (default 1)
  -c int
        Number of columns in which to display the hex dump. Default: 8 columns. (default 8)
  -color
        Choose wether to print a fully rainbowed hex dump. Default: false :(
```

## Building from source

Build binary

```bash
git pull git@github.com:gomills/ggd.git
cd ggd
go build -o ggd .
```

and run it

```bash
./ggd -h
```