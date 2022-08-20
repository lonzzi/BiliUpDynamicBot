// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "lonzzi",
            "url": "https://ronki.moe",
            "email": "lonzzi@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/bili/bot/check": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bili"
                ],
                "summary": "查询Bot登陆状态",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "SESSDATA",
                        "name": "SESSDATA",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.BiliAuthInfo"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/bili/dynamic/getDynamic": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bili"
                ],
                "summary": "获取动态列表(访问b站api)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "up主id",
                        "name": "mid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bili/qrcode/getLoginInfo": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bili"
                ],
                "summary": "获取二维码状态",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "oauthKey",
                        "name": "qrcode",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.BiliQrCodeInfo"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/bili/qrcode/getLoginUrl": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bili"
                ],
                "summary": "获取二维码登录链接",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/bili/reply/add": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bili"
                ],
                "summary": "根据type与oid进行回复",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "回复评论详细信息",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.CommentInfo"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/web/author/add": {
            "post": {
                "description": "需先添加up主之后才能监听动态",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "web"
                ],
                "summary": "添加up主",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "up主id和BotID",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AuthorInfo"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/web/author/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "web"
                ],
                "summary": "获取 up 主列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "BotID",
                        "name": "bot_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/web/bot/list": {
            "get": {
                "description": "根据 Token 获取 Bot 列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "web"
                ],
                "summary": "获取 Bot 列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/web/dynamic/latest": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "web"
                ],
                "summary": "获取up主最新动态",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "up主id和BotID",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AuthorInfo"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/web/dynamic/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "web"
                ],
                "summary": "获取 up 主的动态列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "BotID",
                        "name": "bot_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "up主id",
                        "name": "mid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/web/dynamic/listen": {
            "get": {
                "description": "根据设定的时间间隔监听up主动态",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "web"
                ],
                "summary": "监听up主动态",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "up主id和BotID",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AuthorInfo"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/web/dynamic/status": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "web"
                ],
                "summary": "获取传入的uid的状态",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "up主id和BotID",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AuthorInfo"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/web/dynamic/stop": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "web"
                ],
                "summary": "停止传入的uid的任务",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "up主id和BotID",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AuthorInfo"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/web/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "web"
                ],
                "summary": "站点用户登录",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "SESSDATA",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UserInfo"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/web/refreshToken": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "web"
                ],
                "summary": "刷新 AccessToken",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 刷新令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/web/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "web"
                ],
                "summary": "站点用户注册",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "information",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UserInfo"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "api.AuthorInfo": {
            "type": "object",
            "properties": {
                "bot_id": {
                    "type": "string"
                },
                "mid": {
                    "type": "string"
                }
            }
        },
        "api.BiliAuthInfo": {
            "type": "object",
            "required": [
                "SESSDATA"
            ],
            "properties": {
                "SESSDATA": {
                    "type": "string"
                }
            }
        },
        "api.BiliQrCodeInfo": {
            "type": "object",
            "required": [
                "oauthKey"
            ],
            "properties": {
                "oauthKey": {
                    "type": "string"
                }
            }
        },
        "api.CommentInfo": {
            "type": "object",
            "required": [
                "bot_id",
                "message",
                "oid",
                "type"
            ],
            "properties": {
                "bot_id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "oid": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "api.UserInfo": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "bilibot",
	Description:      "a bilibot server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
