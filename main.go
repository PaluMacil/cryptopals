package main

import (
	"fmt"
	"log"

	"github.com/PaluMacil/cryptopals/set1"
)

func main() {
	fmt.Println("Set 1, Ex 1: Convert Hex to Base 64")
	b64, err := set1.HexToBase64([]byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("%s\n", b64)

	fmt.Println("Set 1, Ex 2: Fixed XOR")
	xor, err := set1.HexDecodeAndXOR([]byte("1c0111001f010100061a024b53535009181c"), []byte("686974207468652062756c6c277320657965"))
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("%x\n", xor)
}
