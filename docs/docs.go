// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-07-05 19:35:56.161972469 -0500 CDT m=+0.057149244

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is the Hackerlog API for collecting editor stats",
        "title": "Hackerlog API",
        "contact": {
            "name": "Deric Cain",
            "url": "https://dericcain.com",
            "email": "deric.cain@gmail.com"
        },
        "license": {},
        "version": "1.0"
    },
    "basePath": "/v1",
    "paths": {
        "/units": {
            "get": {
                "description": "This gets all of the units of work for a specific user. The user is identified by the",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "units"
                ],
                "summary": "Gets units of work for a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "X-Hackerlog-EditorToken",
                        "name": "X-Hackerlog-EditorToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Unit"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Unit": {
            "type": "object",
            "properties": {
                "computer_type": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "editor_type": {
                    "type": "string"
                },
                "file_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "loc_deleted": {
                    "type": "integer"
                },
                "loc_written": {
                    "type": "integer"
                },
                "os": {
                    "type": "string"
                },
                "project_name": {
                    "type": "string"
                },
                "started_at": {
                    "type": "string"
                },
                "stopped_at": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
