
### 实体机部署


#### 前后端合并部署
- 前端打包dist放到后端同一项目中
- 后端设置: 
```
router.Static("/dist", "./dist")
```
- 启动接口项目
<!-- - 启动代理服务器，所有项目只需要一个脚本了: vim onekeyupdate.sh -->
air -c .\.air.conf 
或者
go build -o go_gateway
./go_gateway
