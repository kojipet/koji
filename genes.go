package main

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

func (k *koji) getGenome() string {

	// generate & serialize elements

	// t := fmt.Sprintf("%s", time.Now())

	out := make([]byte, 512)

	hashit := []byte("Father: RockSteady" + k.name)
	sha3.ShakeSum256(out, hashit)

	globalGenes = fmt.Sprintf("%x", out[:])

	k.genome.strand = globalGenes

	k.writeGenes(globalGenes, geneFile)

	return globalGenes

}

func (k *koji) writeGenes(genes, filename string) {
	createDirIfItDontExist(geneDir)
	geneCode := []string{}

	for i := 0; i < len(genes); i = i + 32 {
		if (i + 32) > len(genes) {
			return
		}

		line := fmt.Sprintf("\n%s %s %s %s",
			genes[i:i+8],
			genes[i+8:i+16],
			genes[i+16:i+24],
			genes[i+24:i+32])
		geneCode = append(geneCode, line)
	}

	writeStrings(geneFile, geneCode)
}

func (k *koji) conjugateGenes() {
	// press B to breed, but koji must be a certain age to be able
	// where to place mother and father genes?
	// mother.md father.md could work for now
	// load mother genes
	// load father genes
	motherGenes := readFile(motherFile)
	fatherGenes := readFile(fatherFile)

	// combine them somehow?
	// - maybe hash them both to create a new genome?
	// - maybe combine parts of them?
	// .   but then, how do we determine which parts from mom and which parts
	// .   from father should be inherited?
	// - maybe have inherited genes, but also have parts that are hashed

	out := make([]byte, 512)

	hashit := []byte(motherGenes + fatherGenes)
	sha3.ShakeSum256(out, hashit)

	childGenes := fmt.Sprintf("%x", out[:])
	// return new genome

	k.writeGenes(childGenes, childFile)

	// should the genome have some type of lineage header?
	// maybe thats overcomplicating it.
	// maybe something like
	// mother: <name at time of breeding>
	// father: ''

}
