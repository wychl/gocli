# 复制

```bash
# 从参数复制
gocli copy hello

# 从命令行管道复制
echo hello | gocli copy

# 根据提供的路径复制文件内容
gocli copy -f README.md

# 根据提供的url复制文件内容
gocli copy -u https://raw.githubusercontent.com/wychl/gocli/main/README.md
```

# 粘贴

```bash
gocli paste
```