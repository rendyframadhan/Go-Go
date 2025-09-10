package service

import (
	"fmt"
	"learning-golang/constant"
)

func DescribeString(input string) string {
	output := fmt.Sprintf(
		"%s%s\n%s%d",
		constant.STRING, input,
		constant.TOTAL_CHAR, len(input),
	)
	return output
}
