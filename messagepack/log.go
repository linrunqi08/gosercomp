//go:generate msgp

package messagepack

type LogGroup struct {
	Logs        []*Log    `msg:"log"`
	LogTag      []*LogTag `msg:"logtag"`
	Category    string    `msg:"age"`
	Topic       string    `msg:"topic"`  // this field is ignored
	Source      string    `msg:"source"` // this field is also ignored
	MachineUUID string    `msg:"machineUUID"`
}

type LogTag struct {
	Key   string
	Value string
}

type Log struct {
	Time     uint32
	Contents []*LogContent
}

type LogContent struct {
	Key   string
	Value string
}
