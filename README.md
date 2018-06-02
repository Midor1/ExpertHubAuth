# ExpertHubAuth 鉴权服务

## 概述

这个服务为ExpertHub提供一个会话层。因为API不多，考虑使用传统的REST。

具体鉴权与权限维护可能dalao们还要讨论一下，先鸽了:)

## 部署

### 需要的环境

MySQL/MariaDB > 5.5

Redis > 2.6

### 步骤

Clone之后在工作目录下执行`sudo nohup ./experthubauth &`，也可以使用Windows的版本。

如果想要自己编译，可以在源代码目录下执行`go build main.go`

可以使用提供的`Init.sql`来完成对MySQL数据库的初始化。

运行之前请查看并编辑配置文件`config.toml`中的连接字符串与端口等内容，以确保Redis和MySQL可以被正常访问。

## API文档

### Login [/login]

用户登录，返回0表示成功登录。hashkey之后可能会约定一种散列方式(SHA512)? 我也不是很懂233

+ Request (multipart/form-data)
     + Key-Value Pairs
         + nickname (string, required)
         + hashkey (string, required)

+ Response 200 (application/json)
    + Attribute
        + status          (int)
        + token  (string, uuid)
    + Body
     
          {
              token:c4fa13ae-c17d-4da6-9e59-d660298765a2,
              status:0
          }
          
### Logout [/logout]

注销，成功返回用户ID，失败返回-1。

+ Request (multipart/form-data)
   + Key-Value Pairs
       + token (string, required)

+ Response 200 (application/json)
  + Attribute
      + status          (int)
  + Body
   
        {
            status:3
        }