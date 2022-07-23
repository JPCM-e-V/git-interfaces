package redisrepo

import gitobjects "github.com/JPCM-e-V/git-interfaces/git-objects"

func Init() {}

func LsRefs(repo string) []gitobjects.GitRef {
	return make([]gitobjects.GitRef, 0)
}

func GetObject(id string) string {
	return "tree ff20de5879003250108bb9051e2dbe4e3969d85c\n" +
		"parent 4f6f2847525dbec327170548138475a8c32d3283\n" +
		"author InformaTiger <informatiger@pemoma.de> 1657381568 +0200\n" +
		"committer InformaTiger <informatiger@pemoma.de> 1657381568 +0200\n" +
		"\n" +
		"started entities\n"
}

func AddObject(id string, content string) {}

func SetRef(gitobjects.GitRef) {}

func DeleteRef(id string) {}
