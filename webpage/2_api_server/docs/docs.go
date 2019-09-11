// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-09-10 23:38:37.469166796 +0800 CST m=+0.024955264

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "这是API说明文档，开发服务器才有， 正式部署之后没有.",
        "title": "API说明",
        "termsOfService": "http://9200.gpu.raidcdn.cn:9700",
        "contact": {
            "name": "API Support",
            "url": "http://9200.gpu.raidcdn.cn:9700",
            "email": "ggxxde@163.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://9200.gpu.raidcdn.cn:9700"
        },
        "version": "1.0"
    },
    "host": "192.168.1.102:9000",
    "basePath": "/",
    "paths": {
        "/api1/authping": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "ping检查网络是否正常，需要登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 连通性（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/batchinfo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获得所有批次信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 数据/项目（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/categoryinfo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获得细胞分类信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 数据/项目（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/createdataset": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "新建数据/项目",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 数据/项目（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/dtinfo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获得所有图片及标注的统计信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 数据/项目（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/getimgnptypebymids": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "通过所选中的批次/病例/图片, 返回N/P图片的个数统计",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 数据/项目（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/imglistsofwanted": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获得所选批/次病/细胞类型的图片",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 数据/项目（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/imglistsonebyone": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "通过图片的ID获得对应的所有标注信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 数据/项目（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/job": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "python端请求一个任务（数据处理/训练/预测），python端会指定请求任务的状态和类型",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 任务（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/joblog": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获得任务数据处理/训练/预测后端产生的log",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 任务（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/jobmodel": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获得训练任务生成模型的信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 模型（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/jobpercent": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获得任务进度（数据处理/训练/预测）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 任务（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/jobresult": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获得任务状态（数据处理/训练/预测）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 任务（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "python端更新任务状态/进度（数据处理/训练/预测）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 任务（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/listdatasets": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "按数据库存储顺序依次获得数据/项目信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 数据/项目（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/medicalidinfo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获得所有病例信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 数据/项目（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/ping": {
            "get": {
                "description": "ping检查网络是否正常",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 连通性（不需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/savemodel": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "把训练任务生成模型信息存数据库",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 模型（需要认证）"
                ],
                "responses": {
                    "200": {
                        "description": "{\"ping\": \"pong\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api1/userinfo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获得当前用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 用户"
                ],
                "responses": {
                    "200": {
                        "description": "{\"data\": \"ok\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "注册新用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API1 用户"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "Email",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "手机号码",
                        "name": "Mobile",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "Username",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "Password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "确认密码",
                        "name": "ConfirmPassword",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"data\": \"ok\",\t\"status\": 200}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{ Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface {}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
