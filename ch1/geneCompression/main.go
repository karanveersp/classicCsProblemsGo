package main

import (
	"fmt"
	"math/big"
	"strings"
)

// Reverse reverses the given string.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// CompressedGene models a compressed format gene sequence
type CompressedGene struct {
	bitString *big.Int
}

func (c *CompressedGene) compress(gene string) {
	c.bitString = big.NewInt(1) // sentinel value

	// fmt.Printf("Length of string %d\n", len(gene))
	for _, nucleotide := range strings.ToUpper(gene) {
		c.bitString = c.bitString.Lsh(c.bitString, 2)
		switch nucleotide {
		case 'A':
			c.bitString = c.bitString.Or(c.bitString, big.NewInt(0b00))
		case 'C':
			c.bitString = c.bitString.Or(c.bitString, big.NewInt(0b01))
		case 'G':
			c.bitString = c.bitString.Or(c.bitString, big.NewInt(0b10))
		case 'T':
			c.bitString = c.bitString.Or(c.bitString, big.NewInt(0b11))
		default:
			panic(fmt.Errorf("Invalid nucleotide: %c", nucleotide))
		}
	}
}

// Decompress expands the compressed bit sequence into
// a string of nucleotides A,C,G,T.
func (c *CompressedGene) Decompress() string {
	gene := ""

	// bitlen - 1 to ignore sentinel value
	for i := uint(0); i < uint(c.bitString.BitLen()-1); i += 2 {
		rightShifted := &big.Int{} // local value to not mutate the original
		rightShifted.Rsh(c.bitString, i)
		// fmt.Print("Right shifted: ")
		// display(rightShifted)

		// Getting the two, rightmost bits.
		bits := (rightShifted.And(rightShifted, big.NewInt(0b11))).Uint64()
		// fmt.Print("Bits: ")
		// display(big.NewInt(int64(bits)))
		switch bits {
		case 0b00:
			gene = gene + "A"
		case 0b01:
			gene = gene + "C"
		case 0b10:
			gene = gene + "G"
		case 0b11:
			gene = gene + "T"
		default:
			panic(fmt.Errorf("Invalid bits: %d", bits))
		}
		// fmt.Println(gene)
	}
	return Reverse(gene)
}

func (c *CompressedGene) String() string {
	return c.Decompress()
}

// NewCompressedGene compresses the given gene string into
// a binary string of 0s and 1s
func NewCompressedGene(gene string) *CompressedGene {
	c := CompressedGene{}
	c.compress(gene)
	return &c
}

func display(n *big.Int) {
	fmt.Printf("%b\n", n)
}

func main() {
	s := strings.Repeat("TAGGGATTAACCGTTATATATATATAGCCATGGATCGATTATATAGGGATTAACCGTTATATATATATAGCCATGGATCGATTATA", 100)
	c := NewCompressedGene(s)
	fmt.Printf("String bytes: %d\n", len([]byte(s)))
	fmt.Printf("Compressed bytes: %d\n", len(c.bitString.Bytes()))
	d := c.String()
	fmt.Println(d == s)
}
