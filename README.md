<!--
 * @Author: your name
 * @Date: 2019-11-16 17:49:32
 * @LastEditTime: 2019-12-07 16:38:47
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Impeller_bot/README.md
 -->
 
# 开源Impeller叶轮

基于coolQ-GolangSDK的自律机器人Impeller叶轮

> 最后更新时间为2019年12月07日16:38:19

## 规划的功能

> 大部分功能尚未实装,正在逐步补充中.

- 游戏类
  - 俄罗斯轮盘赌
    - 当前已经完成原版的基本功能并加入了处刑(禁言或出群留学)
  - 基于emjoy的五子棋和扫雷
  
- 爬虫类
  - 胡言乱语生成器
  - 随机野兽先辈新说
  - 音mad推送

- 功能类
  - 基于敏感词的裁定系统(重点开发功能)
    - bot监控用户在下qq群发言,根据敏感词进行裁定
    - 针对敏感词进行禁言或出群留学的操作
    - 屑站(bilibili)和腾讯QQ之间的消息互通.
    - 个性化的表现模式
      - 心理测量者的"支配者"
      - 大屁股企鹅的"动物裁判"
      - 刁蛮要素或者银梦要素的特殊表达方式


## 安装环境

1. [Go语言编译器](https://golang.google.cn/)；
2. [gcc编译器](http://tdm-gcc.tdragon.net/)；  

## 启动酷Q的开发者模式

请查看酷Q官方的[文档](https://d.cqp.me/Pro/%E5%BC%80%E5%8F%91/%E5%BF%AB%E9%80%9F%E5%85%A5%E9%97%A8)

## 修改路径

要修改的地方有几处：
1. go.mod文件第一行，改为你自己项目的地址
2. app.go文件main函数前`// cqp:`开头的注释，修改名称、版本、作者和简介
3. app.go文件init函数内，修改你的AppID
4. 若需要自动复制文件，请设置环境变量DevDir（看看build.bat第20行）

## 编译

运行`build.bat`

最后，在酷Q的菜单-应用管理中，点击重载应用，你应该就能看到你的插件了。