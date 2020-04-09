<!--
 * @Author: your name
 * @Date: 2019-11-16 17:49:32
 * @LastEditTime: 2020-04-09 08:18:08
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Impeller_bot/README.md
 -->

# 开源 Impeller 叶轮

基于 coolQ-GolangSDK 的消息机器人 Impeller 叶轮

> 最后更新时间为 2020 年 04 月 09 日 08:14:33

## 功能

> 因相关 SDK 断代更新,大多数功能失效,正在积极重写相关功能

- 游戏类

  - TDN 轮盘赌
    - 重写中

- 爬虫类

  - 音 mad 推送
  - 曲奇搬运稿件推送

- 功能类

  - 屑站 BV 号转 AV
  - 群聊备份和重建

## 安装环境

1. [Go 语言编译器](https://golang.google.cn/)；
2. [gcc 编译器](http://tdm-gcc.tdragon.net/)；

## 启动酷 Q 的开发者模式

请查看酷 Q 官方的[文档](https://d.cqp.me/Pro/%E5%BC%80%E5%8F%91/%E5%BF%AB%E9%80%9F%E5%85%A5%E9%97%A8)

## 修改路径

要修改的地方有几处：

1. go.mod 文件第一行，改为你自己项目的地址
2. app.go 文件 main 函数前`// cqp:`开头的注释，修改名称、版本、作者和简介
3. app.go 文件 init 函数内，修改你的 AppID
4. 若需要自动复制文件，请设置环境变量 DevDir（看看 build.bat 第 20 行）

## 编译

运行`build.bat`

最后，在酷 Q 的菜单-应用管理中，点击重载应用，你应该就能看到你的插件了。
