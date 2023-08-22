# HTML生成PDF

```bash
# html文件转pdf
gocli html2pdf -f index.html
gocli html2pdf -f ${HOME}/index.html

# 网页链接转pdf
gocli html2pdf -u https://github.com/trending

# 指定输出的pdf文件名
gocli html2pdf -u https://github.com/trending -o github.pdf
```