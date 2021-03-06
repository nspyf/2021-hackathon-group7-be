# 2021-hackathon-backend

# 用户基础API文档

host: https://nspyf.top

port：11000

## 使用前须知

用户名除长度要求外，只能包含大小写字母，数字

密码除长度要求外，只能包含大小写字母，数字，普通可见符号(不含空格)

/auth下的路由请求头均要携带Token：

| Key | Value | 说明 |
| ---  | ---  | ---  |
|  Token  |    |  用户token |

## 方法描述：用户注册

URL地址：/register

请求方法：POST

请求体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  username  |  string  |  用户名,required,min=2,max=16 |
|  password  |  string  |  密码,required,min=8,max=32  |

请求体示例

```
{
    "username":"abc",
    "password":"12345678"
}
```

响应体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |

响应体示例

```
{
    "code": 0,
    "message": "成功"
}
```

## 方法描述：用户登录

URL地址：/login

请求方法：POST

请求体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  user  |  string  |  用户名或邮箱,required |
|  password  |  string  |  密码,required  |

请求体示例

```
{
    "user":"abc",
    "password":"12345678"
}
```

响应体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |
|  data  |  JSON  |  具体看示例  |

响应体示例

```
{
    "code": 0,
    "data": {
        "id": 1,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTUzNTU3NDUsImp0aSI6IjAiLCJzdWIiOiIxIn0.Y038jQ__Dhfz0sFxegB8CcMEAJgt2Svum_0DdFeUiLg",
        "username": "nspyf"
    },
    "message": "成功"
}
```

## 方法描述：验证token

URL地址：/auth/token

请求方法：GET

请求体：无

响应体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |

响应体示例

```
{
    "code": 0,
    "message": "成功"
}
```

## 方法描述：更改密码

URL地址：/auth/password

请求方法：POST

请求体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  password  |  string  |  旧密码,required |
|  new_password  |  string  |  新密码,required,min=8,max=32 |

请求体示例

```
{
    "password":"12345678",
    "new_password":"123456789"
}
```

响应体：

| 字段 | 类型 | 说明 |
| ---  | ---  | ---  |
|  code  |  int  |  状态码  |
|  message  |  string  |  状态码描述  |

响应体示例

```
{
    "code": 0,
    "message": "成功"
}
```