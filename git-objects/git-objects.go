package gitobjects

type GitObjectType uint8

const (
	commit GitObjectType = iota
	blob
	tree
	tag
)

type GitObjectAccessInfo struct {
	objectId   string
	objectType GitObjectType
}

type Blob struct {
	GitObjectAccessInfo
	content string
}

type Tree struct {
	GitObjectAccessInfo
	children []GitObjectAccessInfo
}

type Commit struct {
	GitObjectAccessInfo
	author    Person
	committer Person
	parents   []GitObjectAccessInfo
	tree      GitObjectAccessInfo
	message   string
	gpgsig    string
}

type Tag struct {
	GitObjectAccessInfo
	name    string
	ref     GitObjectAccessInfo
	message string
	gpgsig  string
	tagger  Person
}

type Person struct {
	name      string
	eMail     string
	timestamp Timestamp
}

type Timestamp struct {
	time     uint
	timezone string
}

type GitRef struct {
	name   string
	target GitObjectAccessInfo
}

type ErrorStatus uint8

const (
	userError ErrorStatus = iota
	dataBaseError
)

type StatusError interface {
	error
	Status() ErrorStatus
	ConsoleOutput() string
}

type StatusErrorImpl struct {
	msg           string
	status        ErrorStatus
	consoleOutput string
}

func (e *StatusErrorImpl) Error() string {
	return e.msg
}

func (e *StatusErrorImpl) Status() ErrorStatus {
	return e.status
}

func (e *StatusErrorImpl) ConsoleOutput() string {
	return e.consoleOutput
}
