# rock-syslog-go
syslog 模块

#  rock.syslog.*
- 函数: rock.syslog.server( table )
- 参数: 配置参数服务的必要参数 
```lua
    local kafka = rock.kafka{}
    local file = rock.file{}
    local grpc = rock.grpc{}
     
    local ud = rock.syslog.server{
        protocol  = "udp", -- udp , tcp , udp/tcp
        listen    = "0.0.0.0:514", 
        format    = "line", -- line ,json , raw
        transport = {kafka , file , grpc} -- 所有满足transport模块的userdata
    }

    rock.system.notify() --监控系统退出命
```