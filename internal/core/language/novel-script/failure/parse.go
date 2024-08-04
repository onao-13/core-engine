package failure

import "errors"

var (
	ErrNoChapterName = errors.New("no chapter name")
	ErrNoActName     = errors.New("no act name")

	ErrNoPerson      = errors.New("no person")
	ErrNoPersonName  = errors.New("no person name")
	ErrNoPersonAsset = errors.New("no person asset")

	ErrNoPersonReplica = errors.New("no person replica")

	ErrUnsupportedAssetFormat = errors.New("unsupported asset format")

	ErrParseBackgroundName = errors.New("no background name")

	ErrParseMusicName = errors.New("no music name")

	ErrParseUseType    = errors.New("no use type")
	ErrUseTypeNotFound = errors.New("use type not found")

	ErrNoEnsName      = errors.New("no ens name")
	ErrParseEnsSyntax = errors.New("parse ens syntax error. format <name> \"<file>\"")
	ErrNoEnsFile      = errors.New("no ens file")

	ErrNoSetName         = errors.New("no set name")
	ErrEnsNotLoaded      = errors.New("ens not loaded")
	ErrParseUseEnsSyntax = errors.New("parse use ens syntax error. format <ens-name>.<key> = <value>")
	ErrSetTypeNotFound   = errors.New("set type not found")
)
