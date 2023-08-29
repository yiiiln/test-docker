// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/CRUD": {
            "get": {
                "description": "列出所有",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "GetAllUser",
                "operationId": "GetAllUser",
                "responses": {
                    "200": {
                        "description": "Get All Users Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "註冊",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Register",
                "operationId": "Register",
                "parameters": [
                    {
                        "description": "Register account",
                        "name": "register_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Register Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/CRUD/login": {
            "post": {
                "description": "登入",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Login",
                "operationId": "Login",
                "parameters": [
                    {
                        "description": "Login account",
                        "name": "login_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Register Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/CRUD/uA/{id}": {
            "put": {
                "description": "更新使用者 Account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Update Account By ID",
                "operationId": "Update Account By ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Account Struct",
                        "name": "update_account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.UserRequestUpdateAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update User Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/CRUD/uP/{id}": {
            "put": {
                "description": "更新使用者 Password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "Update Password By ID",
                "operationId": "Update Password By ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Password ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Password Struct",
                        "name": "update_password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.UserRequestUpdatePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update User Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/CRUD/{id}": {
            "get": {
                "description": "列出欲查詢的使用者",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "GetUser",
                "operationId": "GetUser",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get User Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "刪除使用者",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD"
                ],
                "summary": "DeleteUserByID",
                "operationId": "DeleteUserByID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Delete Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delete User Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/TestConnect": {
            "post": {
                "security": [
                    {
                        "BearerOAuth2": []
                    }
                ],
                "description": "測試在 GCP Cloud Run 是否連得上不帶 pSQL 的映像檔",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Test Connect"
                ],
                "summary": "Test Connect Cloud Run",
                "operationId": "Test Connect Cloud Run",
                "responses": {
                    "200": {
                        "description": "Connect Cloud Run Successful",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login/": {
            "post": {
                "description": "登入",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login To Display Access Token"
                ],
                "summary": "ToDisplayAccessToken",
                "operationId": "ToDisplayAccessToken",
                "responses": {
                    "200": {
                        "description": "Get Access Token Successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.UserRequest": {
            "type": "object",
            "properties": {
                "account": {
                    "description": "帳號",
                    "type": "string"
                },
                "password": {
                    "description": "密碼",
                    "type": "string"
                }
            }
        },
        "http.UserRequestUpdateAccount": {
            "type": "object",
            "properties": {
                "new_account": {
                    "description": "更新帳號",
                    "type": "string"
                }
            }
        },
        "http.UserRequestUpdatePassword": {
            "type": "object",
            "properties": {
                "new_password": {
                    "description": "更新密碼",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerOAuth2": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": ""
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "test-connect-api",
	Description:      "test-connect.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
