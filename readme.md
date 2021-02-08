---
   磐石SYSLOG模块
---

# 配置
```lua
    local ud = rock.syslog.server{
        protocol = "udp", -- udp , tcp , udp/tcp
        listen = "0.0.0.0:514", 
        format = "line", -- line ,json , raw
        transport = rock.file{ path = "access.log" } -- 所有满足transport模块的userdata
    }
```