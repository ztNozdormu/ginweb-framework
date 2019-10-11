# 基于GIN框架的基础开发框架,ginweb-framework
## 项目结构
  application  main.go 启动文件
  conf 配置文件
  controller 控制层
  dao 数据交互层
  dto 数据传输层【前端数据传到后台后映射为对应的结构体，校验成功后；转换为数据交互层model】
  middleware 自定义中间件层
  public   公共的工具函数
  router   路由层 路由相关的文件
  tmpl     存放模板文件[前后端分离开发一般不需要]
  log      日志文件
