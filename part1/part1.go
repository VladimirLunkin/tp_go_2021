package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/VladimirLunkin/tp_go_2021/part1/uniq"
)

func main() {
	buf, err := io.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}

	str := string(buf)

	fmt.Println(uniq.Uniq(strings.Split(str, "\n"), uniq.Options{}))

}
