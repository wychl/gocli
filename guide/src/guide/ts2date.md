# 时间戳转日期

```bash
# 时间戳转日期
gocli ts2date 1694102400

# 13位的时间戳转日期
gocli ts2date 1694102400000

# 指定时区
gocli ts2date 1694102400000 --zone=UTC

# 以管道方式执行命令
echo "1694102400" | gocli ts2date
```