package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// LongInt ...
type LongInt struct {
	val  []int8
	sign bool
}

// NewLongInt ...
func NewLongInt(str string) (*LongInt, error) {
	ret := &LongInt{
		sign: true,
		val:  make([]int8, 1, utf8.RuneCountInString(str)+1),
	}
	var nonZero bool
	for i, r := range str {
		if i == 0 && r == '-' {
			ret.sign = false
			continue
		}

		v, err := strconv.ParseInt(string(r), 10, 8)
		if err != nil {
			return nil, err
		}

		if v != 0 {
			nonZero = true
		}
		ret.val = append(ret.val, int8(v))
	}

	if !nonZero {
		ret.sign = true
	}

	return ret, nil
}

// Inc ...
func (l *LongInt) Inc() {
	if l.sign {
		var saved int8 = 1
		for i := len(l.val) - 1; i >= 0 && saved == 1; i-- {
			l.val[i] += saved
			saved = l.val[i] / 10
			l.val[i] %= 10
		}
		return
	}

	for i := len(l.val) - 1; i > 0; i-- {
		l.val[i]--
		if l.val[i] != -1 {
			break
		}
		l.val[i] = 9
	}
}

// Print ...
func (l *LongInt) Print() {
	// var hasNonZero bool
	// for _, v := range l.val {
	// 	if hasNonZero {
	// 		fmt.Print(v)
	// 		continue
	// 	}

	// 	if v != 0 {
	// 		hasNonZero = true
	// 		if !l.sign {
	// 			fmt.Print("-")
	// 		}
	// 		fmt.Print(v)
	// 	}
	// }
	// if !hasNonZero {
	// 	fmt.Print(0)
	// }

	var nonZero int
	for ; nonZero != len(l.val); nonZero++ {
		if l.val[nonZero] != 0 {
			break
		}
	}

	if nonZero == len(l.val) {
		fmt.Printf("0")
		return
	}

	if !l.sign {
		fmt.Printf("-")
	}

	for i := nonZero; i != len(l.val); i++ {
		fmt.Printf("%v", l.val[i])
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s filePath\n", os.Args[0])
		os.Exit(1)
	}

	var str string
	if dat, err := ioutil.ReadFile(os.Args[1]); err != nil {
		fmt.Printf("Problem with file %s: %s\n", os.Args[1], err)
		os.Exit(1)
	} else {
		str = strings.TrimSpace(string(dat))
	}

	if str == "" || str == "-" {
		fmt.Printf("File %s is malformed\n", os.Args[1])
		os.Exit(1)
	}

	lint, err := NewLongInt(str)
	if err != nil {
		log.Fatalf("Incorrect data in file: %v", err)
	}

	lint.Inc()
	lint.Print()
}
