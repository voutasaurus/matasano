package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"bufio"
	//"unicode"
)

// EXERCISE 1
var hexTo64In = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

func hexTo64(hexStr string) (string, error) {
	hexbytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hexbytes), nil
}

// EXERCISE 2
var hexXorIn1 = "1c0111001f010100061a024b53535009181c"
var hexXorIn2 = "686974207468652062756c6c277320657965"

func xorStr(buf1, buf2 string) (string, error) {
	hex1, err := hex.DecodeString(buf1)
	if err != nil {
		return "", errors.New(fmt.Sprint("Cannot decode", buf1, err.Error()))
	}
	hex2, err := hex.DecodeString(buf2)
	if err != nil {
		return "", errors.New(fmt.Sprint("Cannot decode", buf2, err.Error()))
	}
	hex3, err := xor(hex1, hex2)
	if err != nil {
		return "", err	
	}
	return hex.EncodeToString(hex3), nil
}

func xor(buf1, buf2 []byte) ([]byte, error) {
	if len(buf1) != len(buf2) {
		return []byte{}, errors.New("Inputs of different lengths are not allowed.")
	}

	buf3 := make([]byte, len(buf1))
	for i := range buf1 {
		buf3[i] = buf1[i] ^ buf2[i]
	}
	return buf3, nil
}

// EXERCISE 3
var breakXorIn = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	// plain, b, _ := breakXor(breakXorIn)
	// fmt.Println("Decoded message:", plain)
	// fmt.Printf("Private Key: %c\n", rune(b))

func breakXor(message string) (string, byte, float64, error) {
	hexCipher, err := hex.DecodeString(message)
	if err != nil {
		return "", byte(0), 0.0, err
	}

	plainText := ""
	bestScore := 0.01 // ignore any worse than this
	var private byte

	for b := byte(0); b < 255; b++ {
		hexPlain := make([]byte, len(hexCipher))
		for i := range hexCipher {
			hexPlain[i] = hexCipher[i] ^ b
		}
		current := string(hexPlain)

		if currentScore := score(current); currentScore >= bestScore {
			// fmt.Print(current + "\t\tscore: ")
			// fmt.Println(currentScore)
			plainText = current
			bestScore = currentScore
			private = b
		}
	}

	return plainText, private, bestScore, nil
}

		// An alternative here would be to construct a message
		// of the same length made from b and xor them, but
		// there are type/format issues so it's not really
		// more convenient, as you can see:

		// secret := make([]byte, len(message))

		// for i := range secret {
		// 	secret[i] = b
		// }
		// temp, err := xor(hexCipher, secret)
		// if err != nil {
		// 	return "", byte(0), err
		// }
		// current := hex.EncodeToString(temp)

// Exercise 4
	// fmt.Println("Result:")
	// plain, err := findMessage("4.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(plain)
	// return

func findMessage(fileHandle string) (string, error) {
	// open file
	file, err := os.Open(fileHandle)
	if err != nil {
		return "", errors.New(fmt.Sprint("opening:", fileHandle, err))
	}
	defer file.Close()
	results := make([]CrackByteXor,0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		c := CrackByteXor{}
		c.Plain, c.Secret, c.Score, err = breakXor(line)
		if err != nil {
			return "", errors.New(fmt.Sprint("breakingXor:", err))
		}
		if c.Plain != "" {
			c.Cipher = line
			results = append(results, c)
		}
	}
	if err = scanner.Err(); err != nil {
		return "", errors.New(fmt.Sprint("reading from file:", fileHandle, err))
	}

	best := CrackByteXor{Score: 0.1}
	for _, c := range results {
		if c.Score >= best.Score {
			best = c
		}
	}

	//fmt.Println(best)

	return best.Plain, nil
}

type CrackByteXor struct{
	Cipher string
	Plain string
	Secret byte
	Score float64
}

// Exercise 5
	// 	fmt.Println("Result:")
	// 	message := `Burning 'em, if you ain't quick and nimble
	// I go crazy when I hear a cymbal`
	// 	crypt, err := repeatXorEncrypt(message, "ICE")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(crypt)
	// 	return
func repeatXorEncrypt(plain, secret string) (string, error) {
	// create full length secret
	fullSecret := ""
	for len(fullSecret) < len(plain) {
		fullSecret += secret
	}
	fullSecret = fullSecret[:len(plain)]

	encrypted, err := xor([]byte(plain), []byte(fullSecret))
	if err != nil {
		return "", errors.New(fmt.Sprint("xorStr failed:", err))
	}

	return hex.EncodeToString([]byte((encrypted))), nil
}

func rxeFile(fileRead, fileWrite string) error {
	file, err := os.Open(fileRead)
	if err != nil {
		return errors.New(fmt.Sprint("opening:", fileRead, err))
	}
	defer file.Close()

	return nil
}

// Exercise 5B: encrypt a bunch of stuff with repeating XOR

func main() {
	fmt.Println("Result:")

}


