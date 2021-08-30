用户管理


* 1. 登录

用户信息：
```id primary key
staff_id
name
nickname
password
gender
tel
addr
email
department
status
create_at
update_at
delete_at
```


* 2.登录成功后显示用户列表

url->用户列表页面展示:
Controller => Model(获取用户数据) => View => Router



* 3.用户认证

记录用户状态：记录在哪里？
HTTP是无状态的，下一次请求

--->cookie & session 机制

状态记录 => session:
在什么时间记录

登录成功：记录状态（session） sessionid => setcookie sessionid


状态跟踪 => (sessionid) => cookie

用户鉴权：

登录状态:
* 无sessionid
* 有sessionid sessionid无对应session信息
* 有sessionid sessionid无session登录状态信息

未登录（无session标识）
    跳转到登录界面

已登录 => 正常逻辑


beego
```
开启：配置 
    SessionOn=true/false
存储位置：内存，文件，数据库
    SessionProvider：file/mysql/redis
存储的地址：
    SessionProviderConfig
cookie中存储sessionid的名字
    SessionName
失效时间
     SessionGCMaxLifetime = 3600s
        
    操作
    存储session
        controller:SetSession key value(任意类型)
                    编码方式 默认gob 注册
    获取session
        controller:GetSession key => value interface{} 
                    运用断言转换为任意类型
    销毁session
        key1
        key2
        controller:DelSession(key)
        DestorySession()
```

1.session(登录检查)
    在任何需要登录以后才能访问的action执行之前都要进行检查
2.如果访问登录页面
    检查session已存在（用户已登录，就不再打开登录页面，而是直接跳转到首页）



1.公共地方的检查



数据操作 Table
数据：增删改查

数据定义 Table => 列，类型 => 数据 => 增删改查
面向对象 类 => 属性（属性名，类型） => 实例 => 方法调用
ORM（关系映射模型）



静态资源下载到本地：static
/static/ => static


