<div align="center">

<img src="https://i.loli.net/2019/08/07/Kudc3DxaIQ9E8Sj.png" width="160px" alt="icon" />

<h1> 🧧 Apple Product Guide </h1>

[![](https://flat.badgen.net/badge/icon/Product%20Guide?icon=apple&label=Apple)](https://buyersguide.macrumors.com)
[![](https://jaywcjlove.github.io/sb/lang/chinese.svg)](README-zh.md)
[![](https://flat.badgen.net/badge/icon/Heroku?icon=https://simpleicons.now.sh/heroku/fff&label=Deployed%20to&color=795A9D)](https://www.heroku.com/)
[![](https://flat.badgen.net/badge/Lists/iPhone,Mac,Others?list=|&color=red)](https://apguide.herokuapp.com/api)

*Apple Product Guide API. Data acquired from: Buyer's Guide by MacRumors.*

**This API should not be used commercially. Details: [🧸 Disclaimer](#-免责-disclaimer).**

</div>

## 🧬 API

### Request Product List `/api`

```
https://apguide.herokuapp.com/api
```

![](https://i.loli.net/2019/08/13/UqawjJPNXiBeYOo.png)

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

### Request Product Details `/api/{product}`

> Requesting information on the *12.9" iPad Pro* for example.

```
https://apguide.herokuapp.com/api/12.9\"\ iPad\ Pro
```

![](https://i.loli.net/2019/08/13/ePmYpzj8bgS716J.png)

```jsonc
{
    "advice": "Neutral",              // Buying advice by MacRumors
    "average": "542",                 // Average update time (in days)
    "daysSinceLastRelease": "280",    // Days since last update
    "image": "https://cdn.macrumors.com/images-new/buyers-products/ipad_pro_12_9_335.jpg", // MacRumors' device thumbnail
    "name": "12.9\" iPad Pro",        // Device name
    "releaseDate": "Oct 2018",        // Release date
    "status": "Mid-product Cycle"     // Device status
}
```

## 📦 Deploying

- Install [Glide](https://github.com/Masterminds/glide) —— Package Management for Golang
- Install dependencies:

```
glide install
```

- Build project:

```shell
go build
```

- Set environment variable `$PORT` for local server to listen to:

```shell
export PORT=9000
```

- Run project locally：

```shell
./apple-product-guide
```

## 🧮 How is this done?

`main.go` defines a CRON task, which will execute `crawler.go` once every hour. `crawler.go` will automatically fetch the latest data from MacRumors' website, and store the data locally in JSON format.

`main.go` will also execute the API server in another thread defined in `server.go`. Upon every request, `server.go` will read from `data.json` and respond with the corresponding data.

## 🧸 Disclaimer

The "Apple Product Guide" API is not affiliated with MacRumors in any way.

The "Apple Product Guide" API is made for personal use and for personal use only.

---

🧧 **Apple Buyers Guide** ©Spencer Woo. Released under the MIT License.

Authored and maintained by Spencer Woo.

[@Blog](https://spencerwoo.com/) - [ⒿJike](https://web.okjike.com/user/4DDA0425-FB41-4188-89E4-952CA15E3C5E/post) - [@GitHub](https://github.com/spencerwooo)
