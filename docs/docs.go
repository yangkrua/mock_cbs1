// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/albums": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "GetAlbums",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetAlbums"
                ],
                "summary": "GetAlbums",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/api/rlcs/follow": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "RlcsFollow",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RlcsFollow"
                ],
                "summary": "RlcsFollow",
                "parameters": [
                    {
                        "description": "UpdateWrapUpCode Data",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateWrapUpCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/rlcs-autodial-service/api/auth/login": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "RlcsLogin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RlcsLogin"
                ],
                "summary": "RlcsLogin",
                "parameters": [
                    {
                        "description": "Login Data",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RlcsLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Account": {
            "type": "object",
            "properties": {
                "acctNo": {
                    "type": "string"
                },
                "no": {
                    "type": "string"
                },
                "ptpAmount": {
                    "type": "string"
                },
                "ptpDate": {
                    "type": "string"
                }
            }
        },
        "models.RlcsLogin": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "integer"
                }
            }
        },
        "models.UpdateWrapUpCodeRequest": {
            "type": "object",
            "properties": {
                "accountFollows": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Account"
                    }
                },
                "activityCode_Sub": {
                    "type": "string"
                },
                "activity_Code": {
                    "type": "string"
                },
                "activity_Code2": {
                    "type": "string"
                },
                "activity_Desc": {
                    "type": "string"
                },
                "activity_Sub": {
                    "type": "string"
                },
                "activity_Sub2": {
                    "type": "string"
                },
                "cifNo": {
                    "type": "string"
                },
                "collector": {
                    "type": "string"
                },
                "collectorName": {
                    "type": "string"
                },
                "follow_date": {
                    "type": "string"
                },
                "receiveCollector": {
                    "type": "string"
                },
                "telNo": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:9343",
	BasePath:         "",
	Schemes:          []string{"http", "https"},
	Title:            "MOCK RLCS Service",
	Description:      "Testing Swagger APIs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
