package main

import (
	"bufio"
	"ii"
	"os"
	"strings"
)

//自己的github信息
//
//先配置好自己的git ssh 公钥
const (
	meGithub  = "github.com/seaiiok/ii"
	repAlias  = "origin"
	userName  = "seaiiok"
	userEmail = "seaii@qq.com"
	remoteRep = "git@github.com:seaiiok/ii.git"
)

type Git struct {
	meGithub  string
	repAlias  string
	userName  string
	userEmail string
	remoteRep string
}

func main() {

	git := &Git{
		meGithub:  meGithub,
		repAlias:  repAlias,
		userName:  userName,
		userEmail: userEmail,
		remoteRep: remoteRep,
	}
	// 执行git帮助
	git.GitLoopHelp()
}

// GitLoopHelp 执行git相关命令
//
// git add -A
//
// git commit -m ""
//
// git push remoteRep master
//
// 如需pull,请手动
//
func (g *Git) GitLoopHelp() {

	ig := ii.New()

	//避免控制台显示乱码，临时采用UTF-8
	ig.CmdV1.ExecCommand("chcp", "65001")

	//主机
	host, err := os.Hostname()
	if err != nil {
		host = "administrator"
	}
	host = strings.ToLower(host)

	//配置git
	//首次使用手动 git init
	ig.CmdV1.ExecCommand("git", "config", "--global", "user.name", g.userName)
	ig.CmdV1.ExecCommand("git", "config", "--global", "user.email", g.userEmail)
	ig.CmdV1.ExecCommand("git", "remote", "add", g.repAlias, g.remoteRep)

	//循环任务
	//当前仅支持命令:
	//push  推送任务
	//exit  退出git帮助
	ig.CmdV1.Println(13, "start git help...")
	for {
		ig.CmdV1.Println(11, "cmd help input -push...")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		if strings.ToLower(strings.TrimSpace(input.Text())) == "push" {
			ig.CmdV1.Println(11, "push code...")

			output := ig.CmdV1.ExecCommand("git", "add", "-A")
			ig.CmdV1.Println(8, output)

			ig.CmdV1.Println(11, "please input comments...")
			input := bufio.NewScanner(os.Stdin)
			input.Scan()

			if input.Text() == "" {
				output = ig.CmdV1.ExecCommand("git", "commit", "-m", "update by "+host)
			} else {
				output = ig.CmdV1.ExecCommand("git", "commit", "-m", strings.ToLower(strings.TrimSpace(input.Text())))
			}

			ig.CmdV1.Println(8, output)

			output = ig.CmdV1.ExecCommand("git", "push", g.repAlias, "master")
			ig.CmdV1.Println(8, output)

			//更新本地代码库
			output = ig.CmdV1.ExecCommand("gopm", "get", "-u", g.meGithub)
			ig.CmdV1.Println(8, output)

			output = ig.CmdV1.ExecCommand("gopm", "get", "-g", g.meGithub)
			ig.CmdV1.Println(8, output)

		} else if strings.ToLower(strings.TrimSpace(input.Text())) == "exit" {
			os.Exit(0)
		}
	}
}