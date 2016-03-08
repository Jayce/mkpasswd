//
// Created by Liu Chuan (Jayce) in 2016-03-08
//

package mkpasswd

import (
	"fmt"
	"testing"
)

func TestMkpasswd(t *testing.T) {
	fmt.Println(randomCharacters(10, true))
	fmt.Println(MakePassword(2, 2, 2, 2))
}
