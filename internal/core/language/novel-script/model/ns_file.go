package model

type Action struct {
	ChangeEnvironment *ChangeEnvironment
	Replica           *Replica
	ChangeEnsValue    *ChangeEnsValue
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

type NovelScript struct {
	Act      string
	Chapter  string
	Persons  map[string]Person
	Actions  map[int64]Action
	EnsFiles map[string]string
}
