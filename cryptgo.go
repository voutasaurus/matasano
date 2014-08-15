package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	//"unicode"
)

var hexTo64In = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

func hexTo64(hexStr string) (string, error) {
	hexbytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hexbytes), nil
}

var hexXorIn1 = "1c0111001f010100061a024b53535009181c"
var hexXorIn2 = "686974207468652062756c6c277320657965"

func xor(buf1, buf2 string) (string, error) {
	if len(buf1) != len(buf2) {
		return "", errors.New("Inputs of different lengths are not allowed.")
	}
	hex1, err := hex.DecodeString(buf1)
	if err != nil {
		return "", err
	}
	hex2, err := hex.DecodeString(buf2)
	if err != nil {
		return "", err
	}

	hex3 := make([]byte, len(hex1))

	for i := range hex1 {
		hex3[i] = hex1[i] ^ hex2[i]
	}

	return hex.EncodeToString(hex3), nil
}

var breakXorIn = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func breakXor(message string) (string, error) {
	hexCipher, err := hex.DecodeString(message)
	if err != nil {
		return "", err
	}

	//possible := []byte{}
	plainText := ""
	bestScore := 0.01
	for b := byte(0); b < 255; b++ {
		hexPlain := make([]byte, len(hexCipher))
		for i := range hexCipher {
			hexPlain[i] = hexCipher[i] ^ b
		}
		current := string(hexPlain)
		if currentScore := score(current); currentScore >= bestScore {
			fmt.Print(current + "\t\tscore: ")
			fmt.Println(currentScore)
			plainText = current
			bestScore = currentScore
		}
	}

	return plainText, nil
}

func main() {

	//fmt.Println(score("test and something"))

	plain, _ := breakXor(breakXorIn)
	fmt.Println(plain)
}

//fmt.Println(breakXorIn)
//hexCipher, _ := hex.DecodeString(breakXorIn)
//fmt.Println(hexCipher)
//fmt.Println(string(hexCipher))
