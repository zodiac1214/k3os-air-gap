package build

import (
	"context"
	"fmt"
)

/*
interact with hashicorp's packer lib to create machine images
*/

func Packer(ctx context.Context, param BuildParameters) {
	fmt.Println("build packer")
}
