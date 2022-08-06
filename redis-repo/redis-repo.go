package redisrepo

import (
	gitobjects "github.com/JPCM-e-V/git-interfaces/git-objects"
)

type StatusError gitobjects.StatusError

func Init() StatusError {
	return nil
}

func LsRefs(repo string) ([]gitobjects.GitRef, StatusError) {
	return make([]gitobjects.GitRef, 0), nil
}

func GetObject(id string) (string, StatusError) {
	return "tree ff20de5879003250108bb9051e2dbe4e3969d85c\n" +
		"parent 4f6f2847525dbec327170548138475a8c32d3283\n" +
		"author InformaTiger <informatiger@pemoma.de> 1657381568 +0200\n" +
		"committer InformaTiger <informatiger@pemoma.de> 1657381568 +0200\n" +
		"\n" +
		"started entities\n", nil
}

func AddObject(id string, content string) StatusError {
	return nil
}

func SetRef(gitobjects.GitRef) StatusError {
	return nil
}

func DeleteRef(id string) StatusError {
	return nil
}
