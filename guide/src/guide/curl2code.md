# curl生成代码

***注意： 此命令需要联网***
***网址：https://curlconverter.com/***

```bash
# curl命令生成代码
gocli curl2code

# 指定生成的语言
gocli curl2code -l=rust

# 指定curl命令代码
gocli curl2code -l=rust -c="curl 'http://fiddle.jshell.net/echo/html/' \
    -H 'Origin: http://fiddle.jshell.net' \
    -H 'Accept-Encoding: gzip, deflate' \
    -H 'Accept-Language: en-US,en;q=0.8' \
    -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36' \
    -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' \
    -H 'Accept: */*' \
    -H 'Referer: http://fiddle.jshell.net/_display/' \
    -H 'X-Requested-With: XMLHttpRequest' \
    -H 'Connection: keep-alive' \
    --data 'msg1=wow&msg2=such&msg3=data'"

# 管道指定curl命令代码
echo "curl 'http://fiddle.jshell.net/echo/html/' \
    -H 'Origin: http://fiddle.jshell.net' \
    -H 'Accept-Encoding: gzip, deflate' \
    -H 'Accept-Language: en-US,en;q=0.8' \
    -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36' \
    -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' \
    -H 'Accept: */*' \
    -H 'Referer: http://fiddle.jshell.net/_display/' \
    -H 'X-Requested-With: XMLHttpRequest' \
    -H 'Connection: keep-alive' \
    --data 'msg1=wow&msg2=such&msg3=data' --compressed" | gocli curl2code -l=rust

```

## 支持的语言列表：

- ansible
- cfml
- clojure
- csharp
- dart
- elixir
- go
- har
- http
- httpie
- java, java-httpurlconnection, java-jsoup, java-okhttp
- javascript, javascript-jquery, javascript-xhr
- json
- kotlin
- matlab
- node, node-http, node-axios, node-got, node-ky, node-request, node-superagent
- ocaml
- php, php-guzzle, php-requests
- powershell, powershell-webrequest
- python (the default)
- r
- ruby
- rust
- swift
- wget