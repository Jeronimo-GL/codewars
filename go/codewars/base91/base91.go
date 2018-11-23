package base91

import(
	"fmt"
	"unicode/utf8"
)

var encodeTable []byte
var decodeTable []byte
var ebq, en, dbq, dn, dv int

// Init Initializes the arrays used for conversion
func Init(){
	encodeTable=StringToASCIIBytes("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\"")
	decodeTable=make([]byte, 256)

	for i:=0; i<256; i++ {
		decodeTable[i]=255
	}

	for i:=0; i<91; i++ {
		decodeTable[encodeTable[i]]=byte(i)
	}
	EncReset()
	DecReset()
}


// EncReset sends back the counters to the initial position
func EncReset() {
	ebq = 0
	en = 0
}

// DecReset sends back the counters to the initial position
func DecReset(){
	dbq = 0
	dn = 0
	dv = 1
}

//StringToASCIIBytes  converts a string to bytes
func StringToASCIIBytes(s string) []byte {
	t := make([]byte, utf8.RuneCountInString(s))
	i := 0
	for _, r := range s {
		t[i] = byte(r)
		i++
	}
	return t
}


// PrintTables Prints the values if the translation tables
func PrintTables(){
	fmt.Println(encodeTable)
	fmt.Println(decodeTable)
}
