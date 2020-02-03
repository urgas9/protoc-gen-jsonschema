package testdata

const ArrayOfPrimitives = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "string"
                }
            ]
        },
        "keyWords": {
            "items": {
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "string"
                    }
                ]
            },
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "luckyBigNumbers": {
            "items": {
                "oneOf": [
                    {
                        "type": "integer"
                    },
                    {
                        "type": "string"
                    },
                    {
                        "type": "null"
                    }
                ]
            },
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "luckyNumbers": {
            "items": {
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "integer"
                    }
                ]
            },
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        }
    },
    "additionalProperties": true,
    "oneOf": [
        {
            "type": "null"
        },
        {
            "type": "object"
        }
    ]
}`
