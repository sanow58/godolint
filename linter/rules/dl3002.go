package rules

import (
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

// validateDL3002 is "Last user should not be root."
func validateDL3002(node *parser.Node) (rst []ValidateResult, err error) {
	isLastRootUser := false
	var lastRootUserPos int
	for _, child := range node.Children {
		if child.Value == USER {
			if child.Next.Value == "root" || child.Next.Value == "0" {
				isLastRootUser = true
				lastRootUserPos = child.StartLine
			} else {
				isLastRootUser = false
				lastRootUserPos = 0
			}
		}
	}
	if isLastRootUser {
		rst = append(rst, ValidateResult{line: lastRootUserPos, addMsg: ""})
	}
	return rst, nil
}
