package failure

import "errors"

var (
	ErrSyntaxBackground = errors.New("syntax background error")

	ErrSyntaxMusic = errors.New("syntax music error")

	ErrSyntaxSet = errors.New("syntax set error")

	ErrSyntaxUse = errors.New("syntax use error")

	ErrSelectSyntax = errors.New("select syntax error")

	ErrSyntaxGoto = errors.New("syntax goto error")
)
