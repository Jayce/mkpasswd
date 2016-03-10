//
// Created by Liu Chuan (Jayce) in 2016-03-08
//

package main

import (
	"fmt"
	"testing"
)

func distribution(source string, length int, count int) map[rune]int {
	dis := make(map[rune]int)
	for i := 0; i < count; i++ {
		chs := rePermutation(source)
		for _, c := range chs[:length] {
			dis[c]++
		}
	}
	return dis
}

func TestRandomChats(t *testing.T) {
	testchs := "0123456789"
	testchs += "~`!@#$%^&*(_-+=)[]{\\|}\"';:?/.>,<"
	testchs += "abcdefghijklmnopqrstuvwxyz"
	testchs += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	dis := distribution(testchs, len(testchs), 10000)

	totalchs := float64(10000 * len(testchs))
	for c, x := range dis {
		fmt.Printf("%s: %0.2f%%\n", string(c), float64(x)/totalchs*100)
	}
}
