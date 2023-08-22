# gocli

## 使用

```bash
# 安装最新版
go install github.com/wychl/gocli@latest
```

## 支持的命令

- 复制 `echo hello world | gocli copy`
- 粘贴 `gocli paste`
- 压缩 `echo hello world | gocli flate compress`
- 解压缩 `echo hello world | gocli flate decompress`

## 开发

### 准备工作

```bash
cp .cobra.yaml ${HOME}/.cobra.yaml
```

### 添加新命令

```bash
cobra-cli add flate
cobra-cli add compress -p 'flateCmd'
cobra-cli add decompress -p 'flateCmd'
```


### 待实现

### 参考链接

- https://github.com/spf13/cobra-cli/blob/main/README.md
- 使用文档 https://github.com/spf13/cobra/blob/main/user_guide.md#using-the-cobra-library