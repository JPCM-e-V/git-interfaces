package main

type GitObjectType int

const (
	commit GitObjectType = iota + 1
	blob
	tree
	tag
)

type GitObject struct {
	objectId string
}

type Blob struct {
	GitObject
	content string
}

type Tree struct {
	GitObject
	childTrees []Tree
	childBlobs []Blob
}

type Commit struct {
	GitObject
	author    Person
	committer Person
	parents   []Commit
	tree      Tree
	message   string
	gpgsig    string
}

type Tag struct {
	GitObject
	name      string
	refType   GitObjectType
	refObject *GitObject
	message   string
	gpgsig    string
	tagger    Person
}

type Person struct {
	name          string
	eMail         string
	unknownNumber string
	timezone      string
}
