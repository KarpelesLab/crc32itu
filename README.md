[![GoDoc](https://godoc.org/github.com/KarpelesLab/crc32itu?status.svg)](https://godoc.org/github.com/KarpelesLab/crc32itu)

# crc32itu

Package crc32itu implements the ITU I.363.5 CRC-32 algorithm (also known as AAL5 CRC). This algorithm was popularised by BZIP2 and is also used in ATM transmissions.

**This package is particularly useful for interoperability with PHP**, as it produces the same output as PHP's `hash('crc32', ...)` function. Go's standard library `hash/crc32` uses a different polynomial and bit ordering, making it incompatible with PHP's crc32 hash.

The polynomial used is `0x04C11DB7`, with reversed bit shifts compared to the IEEE polynomial.

The algorithm is the same as that in POSIX 1003.2-1992 cksum, except cksum appends the message length to the CRC at the end.

## A note on CRC32 variants

"CRC32" is not a single algorithm—there are many variants that produce different results despite sharing the same name. When integrating systems or verifying checksums, it's important to know which variant you need:

| Variant | Polynomial | Used by |
|---------|------------|---------|
| **CRC-32/ISO-HDLC** (IEEE) | 0xEDB88320 | Go's `hash/crc32`, PNG, ZIP, gzip, Ethernet |
| **CRC-32/BZIP2** (ITU) | 0x04C11DB7 | **This package**, PHP's `hash('crc32', ...)`, BZIP2, AAL5 |
| **CRC-32C** (Castagnoli) | 0x1EDC6F41 | iSCSI, SCTP, ext4 |

If you're getting mismatched checksums between systems, you're likely using the wrong CRC32 variant. Common pitfalls:

- Go's `hash/crc32` (IEEE) does **not** match PHP's `hash('crc32', ...)`
- PHP's `hash('crc32b', ...)` uses IEEE and **does** match Go's `hash/crc32`
- Python's `binascii.crc32()` and `zlib.crc32()` use IEEE

Use this package when you need compatibility with PHP's `hash('crc32', ...)` or other systems using the ITU/BZIP2 variant.

## Installation

```bash
go get github.com/KarpelesLab/crc32itu
```

## Usage

### Simple checksum

```go
package main

import (
	"fmt"
	"github.com/KarpelesLab/crc32itu"
)

func main() {
	data := []byte("Hello, World!")
	checksum := crc32itu.Checksum(data)
	fmt.Printf("CRC-32 ITU: %08x\n", checksum)
}
```

### Incremental updates

```go
package main

import (
	"fmt"
	"github.com/KarpelesLab/crc32itu"
)

func main() {
	crc := crc32itu.Checksum([]byte("Hello, "))
	crc = crc32itu.Update(crc, []byte("World!"))
	fmt.Printf("CRC-32 ITU: %08x\n", crc)
}
```

### Using hash.Hash32 interface

```go
package main

import (
	"fmt"
	"github.com/KarpelesLab/crc32itu"
)

func main() {
	h := crc32itu.New()
	h.Write([]byte("Hello, "))
	h.Write([]byte("World!"))
	fmt.Printf("CRC-32 ITU: %08x\n", h.Sum32())
}
```

## API

- `Checksum(data []byte) uint32` - Returns the ITU I.363.5 checksum of data.
- `Update(crc uint32, data []byte) uint32` - Returns the result of adding bytes to an existing CRC.
- `New() hash.Hash32` - Returns a new hash.Hash32 for computing CRC-32 ITU checksums.
- `Size` - The size of a CRC-32 checksum in bytes (4).

## License

This code is extracted from the Go standard library (`compress/bzip2`) and is licensed under a BSD-style license. See the [LICENSE](LICENSE) file for details.
