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
        "/summary": {
            "get": {
                "description": "weather에 대한 summary를 반환환다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "GetWeatherSummary",
                "parameters": [
                    {
                        "maximum": 90,
                        "minimum": -90,
                        "type": "number",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "maximum": 180,
                        "minimum": -180,
                        "type": "number",
                        "name": "lon",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.WeatherResult"
                        }
                    },
                    "400": {
                        "description": "Invalid parameter requested"
                    },
                    "408": {
                        "description": "request timeout"
                    },
                    "500": {
                        "description": "Internal Server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Weather": {
            "type": "object",
            "properties": {
                "gretting": {
                    "type": "string"
                },
                "heads_up": {
                    "type": "string"
                },
                "temperture": {
                    "type": "string"
                }
            }
        },
        "model.WeatherResult": {
            "type": "object",
            "properties": {
                "summary": {
                    "$ref": "#/definitions/model.Weather"
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