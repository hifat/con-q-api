// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplatev1 = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login request",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authDomain.ReqLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authDomain.ResToken"
                        }
                    },
                    "401": {
                        "description": "Invalid Credentials",
                        "schema": {
                            "$ref": "#/definitions/errorDomain.Response"
                        }
                    },
                    "422": {
                        "description": "Form validation error",
                        "schema": {
                            "$ref": "#/definitions/errorDomain.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errorDomain.Response"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Register request",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authDomain.ReqRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authDomain.ReqRegister"
                        }
                    },
                    "409": {
                        "description": "Duplicate record",
                        "schema": {
                            "$ref": "#/definitions/errorDomain.Response"
                        }
                    },
                    "422": {
                        "description": "Form validation error",
                        "schema": {
                            "$ref": "#/definitions/errorDomain.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errorDomain.Response"
                        }
                    }
                }
            }
        },
        "/healtz": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Healt Check",
                "parameters": [
                    {
                        "description": "Register request",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authDomain.ReqRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
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
        "authDomain.ReqLogin": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 70,
                    "minLength": 8,
                    "example": "Cq123456_"
                },
                "username": {
                    "type": "string",
                    "maxLength": 100,
                    "example": "conq"
                }
            }
        },
        "authDomain.ReqRegister": {
            "type": "object",
            "required": [
                "name",
                "password",
                "username"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "example": "Corn Dog"
                },
                "password": {
                    "type": "string",
                    "maxLength": 70,
                    "minLength": 8,
                    "example": "Cq123456_"
                },
                "username": {
                    "type": "string",
                    "maxLength": 100,
                    "example": "conq"
                }
            }
        },
        "authDomain.ResToken": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string",
                    "example": "eyJhbGciO..."
                },
                "refreshToken": {
                    "type": "string",
                    "example": "eyJhbGciO..."
                }
            }
        },
        "errorDomain.Error": {
            "type": "object",
            "properties": {
                "attribute": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "errorDomain.Response": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/errorDomain.Error"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "security": [
        {
            "bearer": []
        }
    ]
}`

// SwaggerInfov1 holds exported Swagger Info so clients can modify it
var SwaggerInfov1 = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "ConQ API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "v1",
	SwaggerTemplate:  docTemplatev1,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfov1.InstanceName(), SwaggerInfov1)
}
