package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/karanveersp/utils-go/strfuncs"
)

// CompressedGene models a compressed format gene sequence
type CompressedGene struct {
	bits *big.Int
}

func (c *CompressedGene) compress(gene string) error {
	c.bits = big.NewInt(1) // sentinel value

	for _, nucleotide := range strings.ToUpper(gene) {
		c.bits.Lsh(c.bits, 2)
		switch nucleotide {
		case 'A':
			c.bits.Or(c.bits, big.NewInt(0b00))
		case 'C':
			c.bits.Or(c.bits, big.NewInt(0b01))
		case 'G':
			c.bits.Or(c.bits, big.NewInt(0b10))
		case 'T':
			c.bits.Or(c.bits, big.NewInt(0b11))
		default:
			return fmt.Errorf("Invalid nucleotide: %c", nucleotide)
		}
	}
	return nil
}

// Decompress expands the compressed bit sequence into
// a string of nucleotides A,C,G,T.
func (c *CompressedGene) Decompress() (string, error) {
	gene := ""

	// bitlen - 1 to ignore sentinel value
	for i := uint(0); i < uint(c.bits.BitLen()-1); i += 2 {
		rightShifted := &big.Int{} // local value to not mutate the original
		rightShifted.Rsh(c.bits, i)
		// fmt.Print("Right shifted: ")
		// display(rightShifted)

		// Getting the two, rightmost bits.
		rightBits := (rightShifted.And(rightShifted, big.NewInt(0b11))).Uint64()
		// fmt.Print("Bits: ")
		// display(big.NewInt(int64(bits)))
		switch rightBits {
		case 0b00:
			gene = gene + "A"
		case 0b01:
			gene = gene + "C"
		case 0b10:
			gene = gene + "G"
		case 0b11:
			gene = gene + "T"
		default:
			return "", fmt.Errorf("Invalid bits: %d", rightBits)
		}
		// fmt.Println(gene)
	}
	return strfuncs.Reverse(gene), nil
}

func (c *CompressedGene) String() string {
	s, _ := c.Decompress()
	return s
}

// NewCompressedGene compresses the given gene string into
// a binary string of 0s and 1s
func NewCompressedGene(gene string) (*CompressedGene, error) {
	c := CompressedGene{}
	err := c.compress(gene)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func display(n *big.Int) {
	fmt.Printf("%b\n", n)
}

func main() {
	s := strings.Repeat("TAGGGATTAACCGTTATATATATATAGCCATGGATCGATTATATAGGGATTAACCGTTATATATATATAGCCATGGATCGATTATA", 100)
	c, err := NewCompressedGene(s)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("String bytes: %d\n", len([]byte(s)))
	fmt.Printf("Compressed bytes: %d\n", len(c.bits.Bytes()))
	d, err := c.Decompress()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Matches original gene: %v\n", d == s)
}
