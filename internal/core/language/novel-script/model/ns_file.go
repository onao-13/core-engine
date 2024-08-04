package model

type Action struct {
	ChangeEnvironment *ChangeEnvironment
	Replica           *Replica
	ChangeEnsValue    *ChangeEnsValue
	Select            *Select
	Condition         *Condition
	Goto              *Goto
}

type ChangeEnvironment struct {
	BackgroundAsset string
	MusicFile       string
}

type Replica struct {
	PersonName string
	Replica    string
}

type Person struct {
	Name  string
	Asset string
}

type ChangeEnsValue struct {
	Name  string
	Key   string
	Value string
}

type Select struct {
	Variable string
	Values   []string
}

type Condition struct {
	Conditions map[string]string
	Then       []Action
}

type Goto struct {
	File string
}

type NovelScript struct {
	Act      string
	Chapter  string
	Persons  map[string]Person
	Actions  map[int64]Action
	EnsFiles map[string]string
}
