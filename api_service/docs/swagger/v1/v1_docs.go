// Package v1 GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package v1

import "github.com/swaggo/swag"

const docTemplatev1 = `{
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
        "/info": {
            "get": {
                "description": "Получает информацию о приложении",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "info"
                ],
                "summary": "Получить информацию о приложении",
                "operationId": "info",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/sign-up": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Регистрация пользователя",
                "parameters": [
                    {
                        "description": "Запрос на регистрацию пользователя",
                        "name": "signUpRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "409": {
                        "description": "Conflict"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.SignUpRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfov1 holds exported Swagger Info so clients can modify it
var SwaggerInfov1 = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "API сервиса заметок",
	Description:      "Сервис заметок.",
	InfoInstanceName: "v1",
	SwaggerTemplate:  docTemplatev1,
}

func init() {
	swag.Register(SwaggerInfov1.InstanceName(), SwaggerInfov1)
}