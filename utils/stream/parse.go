package stream

import (
	"fmt"
	"os"
)

func ParseStream(argv []string) (input *os.File, output *os.File, err error) {
	if len(argv) >= 1 && argv[0] != "" && argv[0] != "-" {
		input, err = os.OpenFile(argv[0], os.O_RDONLY, 0600)
		if err != nil {
			println(fmt.Sprintf("can not open input file '%s': %s", argv[0], err))
			return nil, nil, err
		}
	} else {
		input = os.Stdin
	}

	if len(argv) >= 2 && argv[1] != "" && argv[1] != "-" {
		output, err = os.OpenFile(argv[1], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
		if err != nil {
			println(fmt.Sprintf("can not open output file '%s': %s", argv[1], err))
			return nil, nil, err
		}
	} else {
		output = os.Stdout
	}

	return input, output, err
}
