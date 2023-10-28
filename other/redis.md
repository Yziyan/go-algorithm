

# Redis Lua 脚本的测试

## 快速使用 Docker 搭建一个环境

```shell
# 创建 redis 容器
docker run -d --name my-redis -p 6379:6379 redis:latest

# 交互式进入 redis-cli 工具
exec -it my-redis redis-cli
```

