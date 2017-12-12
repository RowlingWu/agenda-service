最后更新时间：12.12 22:20

# 分工

项目共分3部分：客户端，服务端，网页端

### 客户端

实现cli目录的代码,通过**使用服务端的接口**,实现在容器内注册和查询等功能.

实现dockerfile, 构造镜像.

### 网页端

使用API blueprint,完成网页开发.

### 服务端

用sqlite数据库.

实现service目录的代码.

其中对于**客户端**，需要提供接口：

```
FindAllUsers() *[]User
CreateUser(user *User) error // 若名字、密码为空，或有其他错误，则要返回error信息，否则返回nil
DeleteUser(id int) error // 若用户ID不存在，或其他错误，返回error
```

对于**网页端**要求，(暂不明），但是由于网页端无法判断用户是否已登录，因此需要服务端来判断。

**判断是否已登录方法：**查看指定路径（在/entity/entity.go里）的curUser.txt文件，里面只存用户ID。若curUser.txt里面存放了用户ID，则用户已登录；若curUser.txt里为空，则未登录。

根据用户ID来进行增删改查。


# 注意事项

项目全过程使用`/entity/entity.go`中`User`结构

请仔细阅读`User`注释

必须阅读agenda项目需求:ex-service-agenda.html

pmlpml博客教程：http://blog.csdn.net/pmlpml/article/details/78727210


 
