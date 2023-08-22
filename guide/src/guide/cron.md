# 解析cron表达式

```bash
# 解析5位cron表达式
gocli cron '0 */2 * * *'

# 解析6位cron表达式
gocli cron '*/5 * * * *'

# 指定时区
gocli cron '*/5 * * * *' --zone=UTC
```