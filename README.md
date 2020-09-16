# makeunit

创建linux systemd 系统服务工具

使用方式：

```
wget
```

参数介绍:
```
./makeunit

::::生成系统服务::::

参数列表:
1:服务名称, 2:执行命令(使用绝对路径表示命令位置, 使用""将完整命令包括其中), 3:工作目录(绝对目录, 可为空)

```

示例:
```
./makeunit devops "/usr/bin/java -jar devops.jar" /data/server/devops
```
 
