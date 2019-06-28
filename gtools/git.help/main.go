package main

import (
	"bufio"
	"gcom"
	"os"
	"strings"
)

//自己的github信息
//
//先配置好自己的git ssh 公钥
const (
	meGithub  = "github.com/seaiiok/gcom"
	repAlias  = "origin"
	userName  = "seaiiok"
	userEmail = "seaii@qq.com"
	remoteRep = "git@github.com:seaiiok/gcom.git"
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

	gCommon := gcom.New()

	

	//避免控制台显示乱码，临时采用UTF-8
	gCommon.GCmd.ExecCommand("chcp", "65001")

	//主机
	host, err := os.Hostname()
	if err != nil {
		host = "administrator"
	}
	host = strings.ToLower(host)

	//配置git
	//首次使用手动 git init
	gCommon.GCmd.ExecCommand("git", "config", "--global", "user.name", g.userName)
	gCommon.GCmd.ExecCommand("git", "config", "--global", "user.email", g.userEmail)
	gCommon.GCmd.ExecCommand("git", "remote", "add", g.repAlias, g.remoteRep)

	//循环任务
	//当前仅支持命令:
	//push  推送任务
	//exit  退出git帮助
	gCommon.GCmd.Println(13, "start git help...")
	for {
		gCommon.GCmd.Println(11, "cmd help input -push...")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		if strings.ToLower(strings.TrimSpace(input.Text())) == "push" {
			gCommon.GCmd.Println(11, "push code...")

			output := gCommon.GCmd.ExecCommand("git", "add", "-A")
			gCommon.GCmd.Println(8, output)

			gCommon.GCmd.Println(11, "please input comments...")
			input := bufio.NewScanner(os.Stdin)
			input.Scan()

			if input.Text() == "" {
				output = gCommon.GCmd.ExecCommand("git", "commit", "-m", "update by "+host)
			} else {
				output = gCommon.GCmd.ExecCommand("git", "commit", "-m", strings.ToLower(strings.TrimSpace(input.Text())))
			}

			gCommon.GCmd.Println(8, output)

			output = gCommon.GCmd.ExecCommand("git", "push", g.repAlias, "master")
			gCommon.GCmd.Println(8, output)

			//更新本地代码库
			output = gCommon.GCmd.ExecCommand("gopm", "get", "-u", g.meGithub)
			gCommon.GCmd.Println(8, output)

			output = gCommon.GCmd.ExecCommand("gopm", "get", "-g", g.meGithub)
			gCommon.GCmd.Println(8, output)

		} else if strings.ToLower(strings.TrimSpace(input.Text())) == "exit" {
			os.Exit(0)
		}
	}
}
