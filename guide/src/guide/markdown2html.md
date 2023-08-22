# Markdown生成HTML

```bash
# markdown文件生成html
gocli markdown2html -f README.md
gocli markdown2html -f ${HOME}/README.md

# 根据markdown网页链接生成html
gocli markdown2html -u https://raw.githubusercontent.com/wychl/gocli/main/README.md

# 指定输出的文件名
gocli markdown2html -u https://raw.githubusercontent.com/wychl/gocli/main/README.md -o index.html

# unix管道方式生成html
cat README.md | gocli markdown2html
```