package main

import (
  "os"
  "bufio"
  "io"
  "math"
)

/*
  fmt.Println("Result:")
  rotating := 0
  for i := 0; i < 10000; i++ {
    fmt.Print(rotating, ", ")
    if rotating == 14 {
      fmt.Println()
    }
    rotating = incrementMod(rotating,15)
  }
  fmt.Println()
*/
func incrementMod(index, modulus int) int {
  index++
  return int(math.Mod(float64(index), float64(modulus)))
}

func bufferDecryptFileRXE(input, output, secret string) {
  bufferEncryptFileRXE(input, output, secret) // encryption and decryption are the same
}

// probably split this into a file function and a bufio readr/writer function
// it's a bit long
func bufferEncryptFileRXE(input, output, secret string) {
  // open input file
  fi, err := os.Open(input)
  if err != nil {
    panic(err)
  }

  // close fi on exit and check for its returned error
  defer func() {
    if err := fi.Close(); err != nil {
      panic(err)
    }
  }()

  // make a read buffer
  r := bufio.NewReader(fi)

  // open output file
  fo, err := os.Create(output)
  if err != nil {
    panic(err)
  }

  // close fo on exit and check for its returned error
  defer func() {
    if err := fo.Close(); err != nil {
      panic(err)
    }
  }()

  // make a write buffer
  w := bufio.NewWriter(fo)

  secretIndex := 0
  secretBuf := []byte(secret)

  // make a buffer to keep chunks that are read
  buf := make([]byte, 1024)
  for {
    // read a chunk
    n, err := r.Read(buf)
    if err != nil && err != io.EOF {
      panic(err)
    }
    if n == 0 {
      break
    }

    for i := range buf[:n] {
      buf[i] ^= secretBuf[secretIndex] // xor is both to encrypt AND decrypt
      secretIndex = incrementMod(secretIndex, len(secret))
    }

    // write a chunk
    if _, err := w.Write(buf[:n]); err != nil {
      panic(err)
    }
  }

  if err = w.Flush(); err != nil {
    panic(err)
  }

  return
}
