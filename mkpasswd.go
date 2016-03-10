//
// Created by Liu Chuan (Jayce) in 2016-03-08
//

package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/docopt/docopt.go"
)

func randomChats(source string, nums int) (chs string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(source)
	for nums > 0 {
		pos := r.Intn(length)
		chs += string(source[pos])
		nums--
	}
	return
}

func rePermutation(source string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(source)
	chs := make([]byte, length)
	for i, v := range r.Perm(length) {
		chs[i] = source[v]
	}
	return string(chs)
}

func mkpasswd(digitn, specialn, lowern, uppern int) (chs string) {
	chs = randomChats("0123456789", digitn)
	chs += randomChats("~`!@#$%^&*(_-+=)[]{\\|}\"';:?/.>,<", specialn)
	chs += randomChats("abcdefghijklmnopqrstuvwxyz", lowern)
	chs += randomChats("ABCDEFGHIJKLMNOPQRSTUVWXYZ", uppern)
	chs = rePermutation(chs)
	return
}

func main() {
	Usage := `generate new password.

Usage:
	mkpasswd [options] [number]

Options:
	-l N         defines the length of the password. The default is 9.
	-d N         the minimum number of digits. The default is 2.
	-c N         the minimum number of lowercase alphabetic characters. The default is 2.
	-C N         the minimum number of uppercase alphabetic characters. The default is 2.
	-s N         the minimum number of special characters. The default is 1.
	-h --help    show help.
	-v --version
`

	arguments, _ := docopt.Parse(Usage, nil, true, "0.0.1", false)

	var (
		length   = 9
		digitn   = 2
		lowern   = 2
		uppern   = 2
		specialn = 1
	)
	for option, value := range arguments {
		if nil == value {
			continue
		}

		if _, ok := value.(string); ok {
			num, err := strconv.Atoi(value.(string))
			if nil != err {
				log.Fatalln(option, err)
			}

			switch option {
			case "-l":
				length = num
			case "-d":
				digitn = num
			case "-c":
				lowern = num
			case "-C":
				uppern = num
			case "-s":
				specialn = num
			}
		}
	}

	if (digitn + specialn + lowern + uppern) > length {
		log.Fatalf("impossible to generate %d-character password with %d numbers, %d lowercase letters, %d uppercase letters and %d special characters.",
			length, digitn, lowern, uppern, specialn)
	}

	lowern += length - (digitn + specialn + lowern + uppern)

	fmt.Println(mkpasswd(digitn, specialn, lowern, uppern))
}
