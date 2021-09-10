1. 基于Prometheus的开发应用 API, 配置, PromQL
    服务发现 => CMDB Prometheus Target 配置
    图形化管理 => 可视化
    告警规则管理
    告警 alertmanager => CMDB
        email
        短信
   
    多个Prometheus
   
    配置文件 yml， json
    API => http client
   
    file_sd_configs:
        a. job管理 对job的增删改查
            id prometheus jobname scheme metrics basic_auth
        b. prometheus 增删改查
            id name key(ip)
        c. target 管理
            target->job 
   
    Prometheus => config

    Agent

api 获取数据
    展示数据，图形化: js + 主机资源监控

Exporter
    Agent => Prometheus client => 统计 采集目标进行通信（socket）

告警 => Db存储
    通知 短信 => 腾讯云
   

2. Prometheus
    服务发现