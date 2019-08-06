# 🧧 Apple Product Guide

> Apple 产品购买指南。数据来自：Buyer's Guide by MacRumors.

## API

### 获取产品列表 `/api`

```
https://apguide.herokuapp.com/api
```

![](https://i.loli.net/2019/08/06/dZAlMJHnuEjx3vL.png)

<details>

```json
[
  "MacBook Pro",
  "iPad Mini",
  "AirPods",
  "HomePod",
  "Mac Mini",
  "MacBook",
  "iPod Touch",
  "11\" iPad Pro",
  "Apple Watch",
  "Mac Pro",
  "iPhone XR",
  "Apple TV",
  "iPad",
  "iPad Air",
  "iPhone XS",
  "12.9\" iPad Pro",
  "MacBook Air",
  "iMac",
  "iMac Pro"
]
```

</details>

### 获取指定产品详情 `/api/{product}`

> 以 12.9" iPad Pro 为例

```
https://apguide.herokuapp.com/api/12.9\"\ iPad\ Pro
```

![](https://i.loli.net/2019/08/06/X2QK8sNI9tMjx6H.png)

```json
{
    "advice": "Neutral",              // 购买建议
    "average": "542",                 // 平均更新时间（天）
    "daysSinceLastRelease": "280",    // 距离上次更新天数
    "image": "https://cdn.macrumors.com/images-new/buyers-products/ipad_pro_12_9_335.jpg", // MacRumor 配图
    "name": "12.9\" iPad Pro",        // 名称
    "releaseDate": "Oct 2018",        // 发布日期
    "status": "Mid-product Cycle"     // 产品状态
}
```

## 部署

- 安装 [Glide](https://github.com/Masterminds/glide) —— Package Management for Golang
- 安装依赖：

```
glide install
```

- 编译项目：

```shell
go build
```

- 指定服务器监听端口环境变量 `$PORT`：

```shell
export PORT=9000
```

- 运行项目：

```shell
./apple-product-guide
```

## 原理

在 `main.go` 中设置了 CRON 任务，`crawler.go` 每小时会自动执行，抓取来自 Buyer's Guide by MacRumors 最新的数据，并以 JSON 的形式存储于 `data.json`。

在 `main.go` 中同时会启动 `server.go` 服务器进程，每次请求读取 `data.json`，并返回响应的数据。

---

🧧 **Apple Buyers Guide** ©Spencer Woo. Released under the MIT License.

Authored and maintained by Spencer Woo.

[@Blog](https://spencerwoo.com/) - [ⒿJike](https://web.okjike.com/user/4DDA0425-FB41-4188-89E4-952CA15E3C5E/post) - [@GitHub](https://github.com/spencerwooo)