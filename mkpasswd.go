//
// Created by Liu Chuan (Jayce) in 2016-03-08
//

package mkpasswd

import (
	"math/rand"
	"time"
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

func randomSpecialChats(nums int) (chs string) {
	return randomChats("~`!@#$%^&*(_-+=)[]{\\|}\"';:?/.>,<", nums)
}

func randomDigitsChats(nums int) (chs string) {
	return randomChats("0123456789", nums)
}

func randomCharacters(nums int, lowercase bool) (chs string) {
	chs = randomChats("ABCDEFGHIJKLMNOPQRSTUVWXYZ", nums)
	if lowercase {
		newStr := make([]rune, len(chs))
		for i, c := range chs {
			newStr[i] = c + 32
		}
		return string(newStr)
	}
	return
}

func MakePassword(digitn, specialn, charn, Charn int) (chs string) {
	chs = randomDigitsChats(digitn)
	chs += randomSpecialChats(specialn)
	chs += randomCharacters(charn, true)
	chs += randomCharacters(Charn, false)
	chs = randomChats(chs, len(chs))
	return
}
