package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/sha3"
)

func (k *koji) getGenome() string {

	// generate & serialize elements

	t := fmt.Sprintf("%s", time.Now())

	out := make([]byte, 512)

	hashit := []byte("Father: RockSteady" + k.name + t)
	sha3.ShakeSum256(out, hashit)

	globalGenes = fmt.Sprintf("%x", out[:])

	k.genome.strand = globalGenes

	k.writeGenes(globalGenes, geneFile)

	return globalGenes

}

func (k *koji) writeGenes(genes, filename string) {

	geneCode := []string{}

	for i := 0; i < len(globalGenes); i = i + 32 {
		if (i + 32) > len(globalGenes) {
			return
		}

		line := fmt.Sprintf("\n%s %s %s %s",
			globalGenes[i:i+8],
			globalGenes[i+8:i+16],
			globalGenes[i+16:i+24],
			globalGenes[i+24:i+32])
		geneCode = append(geneCode, line)
	}

	writeStrings(geneFile, geneCode)
}
