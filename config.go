package main

var (

	// app state

	logLevel         int    = 2
	kojiMessageQueue        = []string{}
	globalGenes      string = ""
	geneDir          string = "./genes/"
	geneFile         string = geneDir + "genes.md"
	motherFile       string = geneDir + "mother.md"
	fatherFile       string = geneDir + "father.md"
	childFile        string = geneDir + "child.md"
)

var (

	// koji sanity

	kojiMaxFood  = 10
	kojiMaxSleep = 10
	kojiMaxHappy = 10
)

type koji struct {

	// our koji

	name   string
	age    int
	awake  bool
	alive  bool
	vitals vitals
	genome genome
}

type genome struct {
	strand string
}

type vitals struct {

	// this kills the koji

	food  int
	sleep int
	happy int
}

var (

	// emotions of koji

	happyFaces = []string{
		"(*- ̬ -)",
		"(- ̬ -*)",
		"(^-^* )",
		"( *^-^)"}

	stressFaces = []string{
		"(  /_')",
		"(  o O)",
		"('o_o')",
		"('-' ;)"}

	sadFaces = []string{
		"(Q_Q  )",
		"(/_\\, )",
		"(-__- )",
		"( -__-)"}

	deadFaces = []string{
		"(X__X)",
		"(x__x)",
		"(x__X)",
		"(X__x)"}

	sleepFaces = []string{
		"(=__=)",
		"(-__-).",
		"(=__=).z",
		"(-__-).zZ"}

	eggFaces = []string{
		" O",
		"O",
		" o",
		"o "}
)
