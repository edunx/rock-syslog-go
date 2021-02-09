---
   磐石SYSLOG模块
---

# 配置
```lua
    local kafka = rock.kafka{}
    local file = rock.file{}
    local grpc = rock.grpc{}
    
    local ud = rock.syslog.server{
        protocol = "udp", -- udp , tcp , udp/tcp
        listen = "0.0.0.0:514", 
        format = "line", -- line ,json , raw
        transport = {kafka , file , grpc} -- 所有满足transport模块的userdata
    }

    rock.notify() --监控系统退出命
```