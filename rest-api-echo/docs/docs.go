// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Api Support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comment": {
            "post": {
                "description": "Create Comment Record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Create Comment Record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "postId",
                        "name": "postId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "body",
                        "name": "body",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "Delete Comment Record By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Delete Comment Record By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/json/comments": {
            "get": {
                "description": "get all Comments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Show all Comments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/posts.Posts"
                            }
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/json/comments/{id}": {
            "get": {
                "description": "get  Comment record By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Show Comment Record By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/comments.Comments"
                            }
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/json/posts": {
            "get": {
                "description": "get all Posts records",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Show all Post in json",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/posts.Posts"
                            }
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/json/posts/{id}": {
            "get": {
                "description": "get  Post record By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Show Post Record By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/posts.Posts"
                            }
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/post": {
            "post": {
                "description": "Create Post Record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Create Post Record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userId",
                        "name": "userId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "body",
                        "name": "body",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "Delete Post Record By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Delete Post Record By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/xml/comments": {
            "get": {
                "description": "get all Comments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Show all Comments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/posts.Posts"
                            }
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/xml/comments/{id}": {
            "get": {
                "description": "get  Comment record By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Show Comment Record By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/comments.Comments"
                            }
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/xml/posts": {
            "get": {
                "description": "get all Posts records",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Show all Post in json",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/posts.Posts"
                            }
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/xml/posts/{id}": {
            "get": {
                "description": "get  Post record By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json",
                    "text/xml"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Show Post Record By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/posts.Posts"
                            }
                        }
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "comments.Comments": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "post_id": {
                    "type": "integer"
                }
            }
        },
        "posts.Comments": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "post_id": {
                    "type": "integer"
                }
            }
        },
        "posts.Posts": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "comment": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/posts.Comments"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
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
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample Rest Api.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
