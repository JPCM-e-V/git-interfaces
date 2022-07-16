package main

type GitObjectType uint8

const (
	commit GitObjectType = iota + 1
	blob
	tree
	tag
)

type GitRef struct {
	objectId   string
	objectType GitObjectType
}

type Blob struct {
	GitRef
	content string
}

type Tree struct {
	GitRef
	children []GitRef
}

type Commit struct {
	GitRef
	author    Person
	committer Person
	parents   []GitRef
	tree      GitRef
	message   string
	gpgsig    string
}

type Tag struct {
	GitRef
	name    string
	ref     GitRef
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
