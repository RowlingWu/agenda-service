![build status](https://travis-ci.org/RowlingWu/agenda-service.svg?branch=master)

# 使用说明

## 客户端使用说明

**注意**：使用客户端前，请先运行服务端：`cd /agenda-service/service`, `go run main.go`

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

服务端使用步骤：

1. 打开终端，使用sqlite3创建agenda-service.db，表User结构见entities.User

2.  运行main.go进行监听

3.  打开终端用curl命令进行测试，具体使用的命令如下

	列出所有用户：

	`curl -v http://localhost:8080/v1/users`

	注册新用户：

	`curl -d "Name=xxx&Passwd=xxx&Email=xxx&Tel=xxx" http://localhost:8080/v1/users`

	用户登录：

	`curl -d "Name=xxx&Passwd=xxx" http://localhost:8080/v1/user/login`

	用户登出：

	`curl -d "Name=xxx" http://localhost:8080/v1/user/logout`

	删除当前用户（这里结合cli命令进行测试），删除当前用户后默认登出。

# 测试结果

## 使用镜像

（此处填写镜像使用结果）

## cli mock测试

（此处填写cli mock测试结果）

## service测试

#### test

![这里写图片描述](http://img.blog.csdn.net/20171218150921479?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

####  在已有用户arcus、fiveking的基础上注册新用户abc，并获取用户列表：

![这里写图片描述](http://img.blog.csdn.net/20171218145048474?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

####  登录用户arcus:

![这里写图片描述](http://img.blog.csdn.net/20171218145511127?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

####  查询当前login列表，登出，并再次查询：

![这里写图片描述](http://img.blog.csdn.net/20171218145715706?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### 登录用户arcus，删除当前用户（使用cli），并查询：

![这里写图片描述](http://img.blog.csdn.net/20171218150242144?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171218150155244?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

#### 进行以上操作时的监听窗口：（405 输错命令）

![这里写图片描述](http://img.blog.csdn.net/20171218150351279?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbGVwcmVjaGF1bl8=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

## 使用命令行（综合测试）

 - 初始状态，没有任何用户，也没有登录的用户
 
    ![初始状态](http://img.blog.csdn.net/20171218163008735?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
 
    直接访问服务端，返回的用户列表为空

    ![curl列出所有用户](http://img.blog.csdn.net/20171218163339737?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

	没有用户登录

	![这里写图片描述](http://img.blog.csdn.net/20171218163559021?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 - 注册一个新用户
 
	![注册&查询](http://img.blog.csdn.net/20171218165518718?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

	服务端后台信息：

	![服务端后台信息-注册](http://img.blog.csdn.net/20171218172358019?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

	访问服务端验证，返回了刚才注册的用户信息

	![注册&查所有用户](http://img.blog.csdn.net/20171218165801334?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 - 登录，并再次查询登录状态

	![登录&查询登录状态](http://img.blog.csdn.net/20171218170049702?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

	服务端后台信息：

	![服务端后台信息-登录](http://img.blog.csdn.net/20171218171537708?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

	访问服务器，验证登录状态

	![curl登陆后访问登录状态](http://img.blog.csdn.net/20171218170306081?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 - 登出，并直接访问服务端验证

	![这里写图片描述](http://img.blog.csdn.net/20171218170531079?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

	服务端后台信息：

	![服务端后台信息-登出](http://img.blog.csdn.net/20171218171353655?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

	服务端显示没有登录的用户：

	![这里写图片描述](http://img.blog.csdn.net/20171218170715830?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


 - 删除该用户（先登录，再删除）

	![删除用户](http://img.blog.csdn.net/20171218170950552?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

	服务端后台信息：

	![服务端后台信息-登录，删除用户](http://img.blog.csdn.net/20171218171150025?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)



