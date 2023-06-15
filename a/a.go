package main

import (
	"fmt"

	"github.com/mpanchajanya/submodules/b"
	subtest "github.com/mpanchajanya/submodules/e2e"
	e2eframework "github.com/mpanchajanya/submodules/test/e2e/framework"
	"github.com/mpanchajanya/submodules/x/y"
)

const Name = b.Name

// const E2EName = e2e.Name

func main() {
	fmt.Println(Name, subtest.E2EName, y.YName, e2eframework.E2EName)
}
