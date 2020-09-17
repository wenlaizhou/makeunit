package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// 1. 描述
// 2. 启动命令
// 3. 工作目录
const tpl = `
[Unit]
Description=%v
Documentation=http://cyclone-robotics.com

[Service]
ExecStart=%v
Restart=always
StartLimitInterval=0
RestartSec=10
WorkingDirectory=%v

[Install]
WantedBy=multi-user.target
`

func main() {
	if len(os.Args) < 2 {
		println("::::生成系统服务::::")
		println("")
		println("参数列表:")
		println("1:服务名称, 2:执行命令(使用绝对路径表示命令位置, 使用\"\"将完整命令包括其中), 3:工作目录(绝对目录, 可为空)")
		return
	}
	name, runnable := os.Args[1], os.Args[2]
	dir := ""
	if len(os.Args) > 3 {
		dir = os.Args[3]
	}
	content := fmt.Sprintf(tpl, fmt.Sprintf("%v", name), runnable, dir)
	err := ioutil.WriteFile(fmt.Sprintf("/usr/lib/systemd/system/%v.service", name), []byte(content), os.ModePerm)
	exec.Command("systemctl", "daemon-reload").CombinedOutput()
	if err != nil {
		println(err.Error())
		return
	}
	println("服务创建完毕")
	println(fmt.Sprintf("可以直接执行 systemctl start %v 启动服务", name))
}
