![build status](https://travis-ci.org/RowlingWu/agenda-service.svg?branch=master)

# 使用说明

## 客户端使用说明

1. 列出所有用户：`cli qu`

2. 查询登录状态（未登录或已登录）：`cli status`

3. 注册：`cli register -u [username] -p [password] -e [Email] -t [Telephone]`

例如：`cli register -u myuser -p 123456 -e my@qq.com -t 13600001111`

4. 登录：`cli login -u [username] -p [password]`

例如：`cli login -u myuser -p 123456`

5. 删除当前用户：`cli del`

6. 登出：`cli logout`

7. 帮助：`cli help`

## 服务端使用说明

（此处填写服务端使用说明）


# 测试结果

## 使用镜像

（此处填写镜像使用结果）

## 使用命令行

（此处填写命令行使用结果）

## cli mock测试

（此处填写cli mock测试结果）

## service测试

（此处填写service测试结果）