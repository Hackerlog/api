// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-07-18 18:51:43.795638454 -0500 CDT m=+0.060286478

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is the Hackerlog API",
        "title": "Hackerlog API",
        "contact": {
            "name": "Deric Cain",
            "email": "deric.cain@gmail.com"
        },
        "license": {},
        "version": "v0.1"
    },
    "basePath": "/v1",
    "paths": {
        "/core/version": {
            "get": {
                "description": "This endpoint takes a few parameters and with those parameters, it looks to see if",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "Returns a link of the latest version of the Core app",
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
                            "type": "object",
                            "$ref": "#/definitions/main.VersionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/mailing-list": {
            "post": {
                "description": "This adds a user to the mailing list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mailing-list"
                ],
                "summary": "Adds a user to the mailing list",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.GenericResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
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
        },
        "/users/login": {
            "post": {
                "description": "Authenticates a user and returns a JWT on successful login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authenticates a user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.Auth"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Auth": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "type": "object",
                    "$ref": "#/definitions/main.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "main.GenericResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "main.Unit": {
            "type": "object",
            "properties": {
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
                    "type": "integer"
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
                    "type": "integer"
                }
            }
        },
        "main.User": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "editor_token": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "units": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Unit"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "main.VersionResponse": {
            "type": "object",
            "properties": {
                "download": {
                    "type": "string",
                    "example": "https://github.com/Hackerlog/core/releases/download/v0.5/core_0.5_windows_amd64.zip"
                },
                "latest": {
                    "type": "boolean",
                    "example": false
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
