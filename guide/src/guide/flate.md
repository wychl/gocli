# 压缩

```bash
# 从参数读取值
gocli flate compress hello

# 从命令行管道取值
echo hello | gocli flate compress

# 根据提供的路径压缩文件内容
cat  README.md | gocli flate compress

# 根据提供的url压缩文件内容
curl https://raw.githubusercontent.com/wychl/gocli/main/README.md | gocli flate compress
```

# 解压

```bash
# 从参数读取值
gocli flate decompress ykjNyckHBAAA//8=

# 从命令行管道取值
echo ykjNyckHBAAA//8= | gocli flate decompress
```