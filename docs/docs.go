// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/tasks": {
            "get": {
                "description": "Retrieve All Tasks",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Retrieve All Tasks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/model.Task"
                                }
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update Task",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Update Task",
                "parameters": [
                    {
                        "description": "Task to be updated",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "post": {
                "description": "Create a Task",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Create a Task",
                "parameters": [
                    {
                        "description": "Task to be created",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/tasks/{taskId}": {
            "get": {
                "description": "Retrieve task by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Retrieve task by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Task"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Task by Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Delete Task by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "taskId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateTaskRequest": {
            "type": "object",
            "required": [
                "due_date",
                "status",
                "title",
                "user_id"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "due_date": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "Pending",
                        "In",
                        "Progress",
                        "Completed"
                    ]
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Task": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "due_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.UpdateTaskRequest": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "Pending",
                        "In",
                        "Progress",
                        "Completed"
                    ]
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{"http"},
	Title:            "Task Management API",
	Description:      "Manage your tasks",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
