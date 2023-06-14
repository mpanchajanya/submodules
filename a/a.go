package main

import (
	"fmt"

	"github.com/mpanchajanya/submodules/b"
	"github.com/mpanchajanya/submodules/test/e2e"
)

const Name = b.Name
const E2EName = e2e.Name

func main() {
	fmt.Println(Name)
}
