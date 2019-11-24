<p align="center">
    <img src="doc/image/index.jpg" alt="logo" width=300 height=200 />
</p>

# Chatroom

[![Go Report Card](https://goreportcard.com/badge/github.com/vouv/chat-room)](https://goreportcard.com/report/github.com/vouv/chat-room) [![Build Status](https://travis-ci.org/vouv/chat-room.svg?branch=master)](https://travis-ci.org/vouv/chat-room) ![license](https://img.shields.io/packagist/l/doctrine/orm.svg) [![donate](https://img.shields.io/badge/%24-donate-ff69b4.svg)](https://github.com/vouv/donate)

> 基于Golang+Vue实现的聊天室

## Tech

1. 实现了三种通信方式

- 刷新-refresh
- 长轮询-long-polling
- 长连接-websocket

2. 基于GO中Channel特性搭建聊天室模型

## 更新日志

### 2019.11.20

- 优化架构
- 优化UI

### 2019.4.30

- 更新UI
- 优化接口

### 2018.12.19

- 优化聊天室逻辑

## 效果图

### 主页

![首页](./doc/image/index.jpg)

### 聊天室

![聊天室](./doc/image/room.jpg)

## Thanks To

- [gin](https://github.com/gin-gonic/gin)
- [gorilla/websocket](https://github.com/gorilla/websocket)
- [vuejs](https://github.com/vuejs/vue)
- [element](https://github.com/ElemeFE/element)
- [axios](https://github.com/axios/axios)
- [js-cookie](https://github.com/js-cookie/js-cookie)
