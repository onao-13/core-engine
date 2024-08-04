package failure

import "errors"

var (
	ErrPersonAlreadyExist = errors.New("person already exist")
	ErrPersonNotFound     = errors.New("person not found")
	ErrEnsAlreadyExist    = errors.New("ens already exist")

	ErrNoEnsFile = errors.New("no ens file")

	ErrParseUseEnsSyntax = errors.New("parse use ens syntax error. format <ens-name>.<key> = <value>")

	ErrSelectSyntaxError = errors.New("select syntax error. format: select <var-name> = \"select1\" | \"select2\"")
)
