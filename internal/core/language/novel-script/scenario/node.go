package scenario

// Region node infos

type PersonReplicaInfo struct {
	Person  string `json:"person"`
	Asset   string `json:"asset"`
	Replica string `json:"replica"`
}

type EnvironmentInfo struct {
	Background string `json:"background,omitempty"`
	Music      string `json:"music,omitempty"`
}

// End node infos

// Region node actions

type Action struct {
	ChangeEns *ActionChangeEns `json:"changeEns,omitempty"`
	Select    *ActionSelect    `json:"select,omitempty"`
	Condition *ActionCondition `json:"condition,omitempty"`
}

type ActionChangeEns struct {
	NewValue string `json:"newValue"`
	Key      string `json:"type"`
	File     string `json:"ensFile"`
}

type ActionSelect struct {
	Variable string   `json:"variable"`
	Values   []string `json:"values"`
}

type ActionCondition struct {
	Variable string `json:"variable"`
	Value    string `json:"value"`
}

type ActionGoto struct {
	File string `json:"file"`
}

// End node actions

type Node struct {
	PersonInfo      *PersonReplicaInfo `json:"personInfo,omitempty"`
	EnvironmentInfo *EnvironmentInfo   `json:"environmentInfo,omitempty"`
	Action          *Action            `json:"action,omitempty"`
	Goto            *ActionGoto        `json:"goto,omitempty"`
}
