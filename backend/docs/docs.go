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
        "/admin/recalculate": {
            "post": {
                "description": "Start recalculating matches",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Recalculate matches",
                "responses": {
                    "200": {
                        "description": "recalculation done",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/mmr/matches": {
            "get": {
                "description": "Get all matches",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Matches"
                ],
                "summary": "Get matches",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/view.MatchDetails"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Submit a match for MMR calculation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Matches"
                ],
                "summary": "Submit a match",
                "parameters": [
                    {
                        "description": "Match object",
                        "name": "match",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.Match"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "match submitted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/stats/leaderboard": {
            "get": {
                "description": "Get leaderboard stats including wins, loses, and MMR of users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Leaderboard"
                ],
                "summary": "Get leaderboard stats",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/repos.LeaderboardEntry"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "repos.LeaderboardEntry": {
            "type": "object",
            "properties": {
                "loses": {
                    "type": "integer"
                },
                "mmr": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "wins": {
                    "type": "integer"
                }
            }
        },
        "view.Match": {
            "type": "object",
            "required": [
                "team1",
                "team2"
            ],
            "properties": {
                "team1": {
                    "$ref": "#/definitions/view.MatchTeam"
                },
                "team2": {
                    "$ref": "#/definitions/view.MatchTeam"
                }
            }
        },
        "view.MatchDetails": {
            "type": "object",
            "required": [
                "date",
                "team1",
                "team2"
            ],
            "properties": {
                "date": {
                    "type": "string"
                },
                "mmrCalculations": {
                    "$ref": "#/definitions/view.MatchMMRCalculationDetails"
                },
                "team1": {
                    "$ref": "#/definitions/view.MatchTeam"
                },
                "team2": {
                    "$ref": "#/definitions/view.MatchTeam"
                }
            }
        },
        "view.MatchMMRCalculationDetails": {
            "type": "object",
            "required": [
                "team1",
                "team2"
            ],
            "properties": {
                "team1": {
                    "$ref": "#/definitions/view.MatchMMRCalculationTeam"
                },
                "team2": {
                    "$ref": "#/definitions/view.MatchMMRCalculationTeam"
                }
            }
        },
        "view.MatchMMRCalculationTeam": {
            "type": "object",
            "required": [
                "player1MMRDelta",
                "player2MMRDelta"
            ],
            "properties": {
                "player1MMRDelta": {
                    "type": "integer"
                },
                "player2MMRDelta": {
                    "type": "integer"
                }
            }
        },
        "view.MatchTeam": {
            "type": "object",
            "required": [
                "member1",
                "member2",
                "score"
            ],
            "properties": {
                "member1": {
                    "type": "string"
                },
                "member2": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
