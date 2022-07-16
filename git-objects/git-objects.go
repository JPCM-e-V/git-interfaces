package main

type GitObjectType uint8

const (
	commit GitObjectType = iota + 1
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
