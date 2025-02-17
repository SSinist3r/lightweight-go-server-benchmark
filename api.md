[![genDocsGPT](https://img.shields.io/badge/Doc%20generated%20by-genDocsGPT-blue)](https://github.com/marco-rosner/genDocsGPT)

# API Documentation

## Table of Contents
- [Models](#models)
- [Endpoints](#endpoints)
  - [Add Person](#add-person)
  - [Search Person](#search-person)
  - [Get Person](#get-person)
  - [Count People](#count-people)

## Models
### Person
| Field      | Type      | Description                                |
|------------|-----------|--------------------------------------------|
| ID         | string    | Unique identifier for the person            |
| Name       | string    | Name of the person                          |
| Nickname   | string    | Nickname of the person                      |
| Birth      | string    | Date of birth of the person (required)      |
| CreatedAt  | time.Time | Timestamp of when the person was created    |
| UpdateAt   | time.Time | Timestamp of when the person was last updated |

## Endpoints

### Add Person
- Method: POST
- URL: /people

**Request Body:**
```json
{
  "id": "string",
  "name": "string",
  "nickname": "string",
  "birth": "string"
}
```

**Example:**
```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "id": "123",
  "name": "John Doe",
  "nickname": "johnd",
  "birth": "1990-01-01"
}' http://localhost:8080/people
```

### Search Person
- Method: GET
- URL: /people

**Query Parameters:**
- t: Search term (required)

**Example:**
```bash
curl -X GET 'http://localhost:8080/people?t=john'
```

### Get Person
- Method: GET
- URL: /people/:id

**URL Parameters:**
- id: Person ID (required)

**Example:**
```bash
curl -X GET 'http://localhost:8080/people/123'
```

### Count People
- Method: GET
- URL: /contagem-people

**Example:**
```bash
curl -X GET 'http://localhost:8080/contagem-people'
```

*This documentation file was generated using [genDocsGPT](https://github.com/marco-rosner/genDocsGPT)*