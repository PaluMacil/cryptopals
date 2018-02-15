/* Package set1 Basics
Fixed XOR
Write a function that takes two equal-length buffers and produces their XOR combination.

If your function works properly, then when you feed it the string:

1c0111001f010100061a024b53535009181c
... after hex decoding, and when XOR'd against:

686974207468652062756c6c277320657965
... should produce:

746865206b696420646f6e277420706c6179
*/
package set1

import (
	"encoding/hex"
)

func HexDecodeAndXOR(val1, val2 []byte) ([]byte, error) {
	decoded := make([]byte, hex.DecodedLen(len(val1)))
	_, err := hex.Decode(decoded, val1)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(decoded))
	n := len(decoded)
	if len(val2) < n {
		n = len(val2)
	}
	for i := 0; i < n; i++ {
		result[i] = decoded[i] ^ val2[i]
	}

	return result, nil
}
