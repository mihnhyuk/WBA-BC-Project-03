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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/all": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call GetLiquidityList, return ok by json.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/pool": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call AddLiquidity, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "private key",
                        "name": "pk",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Token address and amount",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.AddLiquidityInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call CreateLiquidity, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "private key",
                        "name": "pk",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Token address",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateLiquidityInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call RemoveLiquidity, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "private key",
                        "name": "pk",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Token address and amount",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.RemoveLiquidityInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/pool/balance": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call BalanceOf, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "pool name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "my address",
                        "name": "address",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        },
        "/swap": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call Swap, return ok by json.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "private key",
                        "name": "pk",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Token symbol and amount",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SwapInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Controller"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.AddLiquidityInput": {
            "type": "object",
            "properties": {
                "amount_a": {
                    "type": "string"
                },
                "amount_b": {
                    "type": "string"
                },
                "sym_a": {
                    "type": "string"
                },
                "sym_b": {
                    "type": "string"
                }
            }
        },
        "controller.Controller": {
            "type": "object"
        },
        "controller.CreateLiquidityInput": {
            "type": "object",
            "properties": {
                "token1": {
                    "type": "string"
                },
                "token2": {
                    "type": "string"
                }
            }
        },
        "controller.RemoveLiquidityInput": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controller.SwapInput": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "input": {
                    "type": "string"
                },
                "output": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
