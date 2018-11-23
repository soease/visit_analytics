# visit_analytics
网站访问量统计工具,现已改为sqlite3版。

## 截图
### 总访问量
![screenshot](./screenshot.png)
### 显示日(月)访问量
![screenshot2](./screenshot2.png)

## 使用
### 下载安装
``` sh
git clone https://github.com/nladuo/visit_analytics.git
cd visit_analytics
make prepare 		# 下载库文件
```

### 创建数据库
自动创建

### 配置文件
打开config,yaml, 配置mysql与访问密码。
``` yaml
manage:
  username: nladuo
  password: nladuo
db:
  dbname: visit_analytics.db
deploy_host: localhost:3000
run_at: :3000
```

### 安装运行
``` sh
make && ./visit_analytics config.yaml
```

### 部署
``` sh
./deploy.sh
```

### 查看访问量
打开[http://localhost:3000/manage](http://localhost:3000/manage)(将localhost:3000替换成你部署的服务器地址)

## 使用案例
见[test.html](./www/test.html)，添加`<script type="text/javascript" src="http://localhost:3000/analytics.js"></script>`到要统计的网页中。(将localhost:3000替换成你部署的服务器地址)

## LICENSE
MIT
