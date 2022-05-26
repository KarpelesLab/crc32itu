[![GoDoc](https://godoc.org/github.com/KarpelesLab/crc32itu?status.svg)](https://godoc.org/github.com/KarpelesLab/crc32itu)

# crc32 ITU

ITU I.363.5 algorithm (a.k.a. AAL5 CRC) was popularised by BZIP2 but also used in ATM transmissions.

This is not implemented in golang's `hash/crc32` package, but can be found hidden in `compress/bzip2` and is the equivalent of PHP's `hash('crc32', ...)`.

the algorithm is the same as that in POSIX 1003.2-1992 in Cksum but that stuffs the size into the CRC at the end for extra measure.
