# 日期转时间戳

```bash
# 时间转时间戳
gocli date2ts '2023-09-08 00:00:00'

# 生成13位的时间戳
gocli date2ts '2023-09-08 00:00:00' --size=13

# 指定时区
gocli date2ts '2023-09-08 00:00:00' --zone=UTC
```