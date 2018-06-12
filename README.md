# ExpertHubAuth 鉴权服务

## 概述

这个服务为ExpertHub提供一个会话层。因为API不多，考虑使用反向REST，所有URL都是动词:)



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
        
### Register [/register]

注册普通用户，返回用户的UID。

+ Request (multipart/form-data)
     + Key-Value Pairs
         + nickname (string, required)
         + hashkey (string, required)

+ Response 200 (application/json)
  + Attribute
      + status          (int)
  + Body
   
        {
            status:3
        }
        
### Check [/check]

用token取得UID，不存在则返回-1。

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

### Grant [/grant]

普通用户登记一个专家权限，nickname可以作为该用户这次登记的专家头衔。

系统会向提供的邮箱内发送验证码，可以调用Validate的API验证，否则AccountStatus字段会始终为-1(停用)状态。

+ Request (multipart/form-data)
     + Key-Value Pairs
         + nickname (string, required)
         + token (string, required)
         + email (string, required)

+ Response 200 (application/json)
  + Attribute
      + status          (int)
  + Body
   
        {
            status:0
        }
        
### Validate [/validate]

验证邮箱验证码，成功返回专家的ID(EID)，失败返回-1。

+ Request (multipart/form-data)
     + Key-Value Pairs
         + captcha (string, required)
         + token (string, required)

+ Response 200 (application/json)
  + Attribute
      + status          (int)
  + Body
   
        {
            status:0
        }

### Map [/map]

从用户的token取得他登记的全部专家。

+ Request (multipart/form-data)
     + Key-Value Pairs
         + token (string, required)

+ Response 200 (application/json)
  + Attribute
      + eid          (int[])
  + Body
   
        {
            "eid": [
                1,
                2
            ]
        }

### CRUD [/crud]

取得或编辑专家信息。如果不发送token则为查询，为幂等操作；如果发送了token则为编辑操作。

+ Request (multipart/form-data)
     + Key-Value Pairs
         + eid           (int, required)
         + token         (string, optional)
         + Nickname      (string, optional)
         + Gender        (string, optional)
         + Email         (string, optional)
         + Tel           (string, optional)
         + Subgroup      (string, optional)
         + Category      (string, optional) 
         + Avatar        (string, optional)

+ Response 200 (application/json)
  + Attribute 1
      + info          (Info)
  + Body 1
   
        {
            "eid": 1,
            "uid": 4,
            "nickname": "AuthExpert",
            "gender": "Male",
            "email": "",
            "tel": "",
            "subgroup": "",
            "category": "",
            "avatar": "Avatar",
            "status": 0
        }
  + Attribute 2
       + status      (int)
  + Body 2
     
        {
            "status": 0
        }      
