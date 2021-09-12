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

用户管理

任务管理

资产管理
机房管理
工单管理
告警管理
统计图表

1. 编辑
   a.打开编辑页面
    Get id => 查找数据 => 页面渲染
   b.提交
    Post id/xxx => 更新数据 => （成功）跳转到列表页面/（失败）提示
   
Flash:处理成功后将消息存储 cookie；如果想要显示时从存储中获取消息并显示 从cookie中读取/删除
      页面使用后端模板：从当前Controller某个URL跳转到另一个URL
      存储: flash := beego.Newflash   
            flash.Set(key, value)
            flash.Store(&Controller)
      获取：   beego.ReadFromRequest(&c.Controller)
              flash.Data
              c.Data[key] = value


CSRF:网络攻击-扩展请求伪造
     
     a.配置
     开启CSRF防护，配置Token => Key
     过期时间
     
     b.打开页面的时候生活曾token，从controller生成，传递到页面

     c.提交时必须提交token，beego自动验证(Post, Delete, Put)，csrf_token =>


cobra
    beego orm --db --force --verbose

    main web
         db init

管理员：操作用户管理
        UserController
        Prepare => 判断用户是否是管理员 => 不是跳转到无权限页面/是管理员则继续访问

普通用户：任务管理

限制普通用户操作UserController
        发起请求 => 请求操作成功

限制未登录用户操作UserController
        session => User 有 放行
                        无 跳转到登录页面

role {
    1 : [controller1.Action1, controller2.Action2],
    2 : [controller1.Action1, controller2.Action2]
}

Authorization => session => user => role => actions
            GetControllerAndAction
    在 actions => 有权限
    不在 => 无权限

sudo rm -rf 文件

Prometheus
    Node => 查询 删除
    Agent   => API register => 不存在 天剑
                             存在 更新
            => 属性
                uuid
                hostname
                addr https://host:port/
                #username
                #password
                created_at
                updated_at
                deleted_at

    Job => 增删改查
        => 属性
                任务标识 [a-zA-Z][0-9a-zA-Z_]
                备注
                Node node_id
                created_at
                updated_at
                deleted_at

    Target => 增删改查
                名称
                备注
                Addr
                Job
                created_at
                updated_at
                deleted_at

数据库关系
用户 <=> 密码
user：id name
password：id user_id password

api => json 

controller => json
beego:
    CopyRequestBody=true
    c.Input.RequestBody json => json.Unmarshal => Form `json`

beego:xsrf检查 => 关闭xsrf检查
    认证： 登录 => set-cookie: sessionid, cookie: sessionid
            Token: 固定的随机字符串 V
                    JWT Token（认证）
                    id key => id params(timestamp) 签名Token

            header: Authrozation: Token xxx
                                  Bearer xxx



Agent:
    1.注册
    2.获取配置 => 生成Prometheus配置
