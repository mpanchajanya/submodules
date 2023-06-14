package main

import (
	"fmt"

	"github.com/mpanchajanya/submodules/b"
	subtest "github.com/mpanchajanya/submodules/e2e"
)

const Name = b.Name

// const E2EName = e2e.Name

func main() {
	fmt.Println(Name, subtest.E2EName)
}
