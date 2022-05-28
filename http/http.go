package main

import (
	"fmt"
	"log"
	"net/http"

	gitutils "github.com/JPCM-e-V/git-interfaces/gitutils"
	redisrepo "github.com/JPCM-e-V/git-interfaces/redis-repo"
)

const reponame string = "test"

func PrintRequest(r *http.Request) {
	fmt.Printf("%s %s %s", r.Method, r.URL, r.Proto)
	if r.ContentLength > 0 {
		fmt.Printf(" Content: %d bytes of %s", r.ContentLength, r.Header.Get("Content-Type"))
	}
	fmt.Println()
}

func GitUploadPackInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-git-upload-pack-advertisement")
	// WriteGitProtocol(w, []string{"# service=git-upload-pack"})
	gitutils.WriteGitProtocol(w, []string{"version 2", "ls-refs", "fetch"})
}

func GitUploadPack(w http.ResponseWriter, r *http.Request) {
	lines, err := gitutils.ReadGitProtocol(r.Body)
	if err == nil {
		var command string
		for _, line := range lines {
			if len(line) > 9 && line[:9] == "pcommand=" {
				command = line[9:]
			}
		}
		if command == "ls-refs" {
			// gitutils.WriteGitProtocol(w, []string{"8ed3ded8cb3ecff8345165ad40dbd36f421bfb2a HEAD"})
			if refs, err := redisrepo.LsRefs(reponame); err == nil {
				gitutils.WriteGitProtocol(w, refs)
			} else {
				w.WriteHeader(500)
				gitutils.WriteGitProtocol(w, []string{"ERR InternalServerError: " + err.Error()})
				fmt.Print(err.Error())
			}
		} else if command == "fetch" {
			fmt.Println(lines)
		}
	} else {
		w.WriteHeader(400)
		gitutils.WriteGitProtocol(w, []string{"ERR Bad Request: " + err.Error()})
	}
}

func GitReceivePackInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(1)
	w.Header().Set("Content-Type", "application/x-git-receive-pack-advertisement")
	gitutils.WriteGitProtocol(w, []string{"# service=git-receive-pack"})
	gitutils.WriteGitProtocol(w, []string{"0000000000000000000000000000000000000000 capabilities^{}\x00report-status"})

	// gitutils.WriteGitProtocol(w, []string{"version 2", "ls-refs", "fetch"})
}

func GitReceivePack(w http.ResponseWriter, r *http.Request) {
	lines, err := gitutils.ReadGitProtocol(r.Body)
	fmt.Println(lines, err)

	// buf := new(bytes.Buffer)
	// gitutils.WriteGitProtocol(buf, []string{"unpack ok", "ok refs/heads/main"})
	// fmt.Fprintf()
	// gitutils.WriteGitProtocol(w, []string{"\x01" + buf.String()})
	// fmt.fprint
	// fmt.Fprintf(w, "%s\x01%s0000", gitutils.PktLine("unpack ok\n"), gitutils.PktLine("ok refs/heads/main\n"))
	// fmt.Fprint(w, "\x30\x30\x30\x65\x75\x6e\x70\x61\x63\x6b\x20\x6f\x6b\x0a\x30\x30\x31\x37\x6f\x6b\x20\x72\x65\x66\x73\x2f\x68\x65\x61\x64\x73\x2f\x6d\x61\x69\x6e\x0a\x30\x30\x30\x30")
	gitutils.WriteGitProtocol(w, []string{"unpack ok", "ok refs/heads/main"})
}

type GitHandler struct {
	gitUploadPackInfoHandler  http.Handler
	gitUploadPackHandler      http.Handler
	gitReceivePackInfoHandler http.Handler
	gitReceivePackHandler     http.Handler
}

func (g *GitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	PrintRequest(r)
	if r.Method == "GET" && r.URL.Path == "/info/refs" {
		fmt.Println(r.URL.Query().Get("service"))
		if r.URL.Query().Get("service") == "git-upload-pack" {
			g.gitUploadPackInfoHandler.ServeHTTP(w, r)
			return
		}
		if r.URL.Query().Get("service") == "git-receive-pack" {
			g.gitReceivePackInfoHandler.ServeHTTP(w, r)
			return
		}
	} else if r.Method == "POST" && r.URL.Path == "/git-upload-pack" {
		g.gitUploadPackHandler.ServeHTTP(w, r)
		return
	} else if r.Method == "POST" && r.URL.Path == "/git-receive-pack" {
		g.gitReceivePackHandler.ServeHTTP(w, r)
		return
	}
	w.WriteHeader(404)
	gitutils.WriteGitProtocol(w, []string{"ERR Not Found"})
}

func main() {
	redisrepo.Init()
	var s *http.Server = &http.Server{
		Addr: ":8080",
		Handler: &GitHandler{
			gitUploadPackInfoHandler:  http.HandlerFunc(GitUploadPackInfo),
			gitUploadPackHandler:      http.HandlerFunc(GitUploadPack),
			gitReceivePackInfoHandler: http.HandlerFunc(GitReceivePackInfo),
			gitReceivePackHandler:     http.HandlerFunc(GitReceivePack),
		},
	}
	fmt.Println("Running on http://localhost:8080")
	log.Fatal(s.ListenAndServe())
}

// func main() {
// 	fmt.Printf("%q", PktLine("version 2"))
// }
