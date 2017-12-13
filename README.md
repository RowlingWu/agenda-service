最后更新时间：12.13  12:50


# 分工

项目共分3部分：客户端，服务端，网页端

### 客户端

实现cli目录的代码,通过**使用服务端的接口**,实现在容器内注册和查询等功能.

实现dockerfile, 构造镜像.

实现登录、登出、注册、删除当前用户、查询所有用户5个功能.

### 网页端

使用API blueprint，使用服务端接口.

要能**登录、登出、注册、删除当前用户、查询所有用户**.

### 服务端

用sqlite3数据库.

实现service目录的代码，提供各种Handler.


# 注意事项

项目全过程使用`entity/entity.go`::`User`结构

请仔细阅读`User`注释

必须阅读agenda项目需求: ex-service-agenda.html

pmlpml博客教程： http://blog.csdn.net/pmlpml/article/details/78727210
