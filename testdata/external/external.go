package external

type ExString string

const (
	ExStr ExString = "ex"
)

type ExInt int

type ExFloat float64

type ExComplex complex128

type ExBool bool

type ExRune rune

type privExString string

const (
	PrivExStr privExString = "priv_ex"
)

func TakesPrivExString(s privExString) {}

var PrivExStringChan = make(chan privExString)

var PrivExStringMap = map[privExString]int{
	privExString("a"): 1,
	privExString("b"): 2,
	privExString("c"): 3,
}
