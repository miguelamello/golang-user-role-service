# API Reference for User Domain Role Service

## Description
This API reference provides information about the functionality of the 
User Domain Role Service API. The User Domain Role Service can provide a centralized and unified place to store and manage user roles across different domains or applications within a corporate environment. The API provides functionality described below.

## Version
No versioning is used for this API. This is a evolving API and the changes will 
be reflected in the reference. We garantee that the changes will be backward 
compatible. 

## Servers
The API is hosted on the following servers:

Production server: `http://orionsoft.site/udsr`<br>
Development server: `http://localhost:8080/udsr/graphql`

## Endpoints
Differently from REST APIs, GraphQL APIs have only one endpoint where all the 
requests are sent. The endpoint for this API is:

Remote endoint: `http://orionsoft.site/udsr`<br>
Local endpoint: `http://localhost:8080/udsr/graphql`

### Methods
The API provides only one method: `POST`

### Headers
The API requires the following headers:

`Content-Type: application/json`

### Request Body
 The request body is required and must be a valid GraphQL query. 

 - For development:
 We recommend you to use a GraphQL desktop client to query the API. 
 Usually, the client will provide a GraphQL body to query the API.

 Obs: When using in development you can access the GraphQL playground at 
  `http://localhost:8080/` via browser. The playground will provide a graphical 
  interface to query the API.

 - For production:
 We recommend you to use a GraphQL client module suitable to your 
 programing language and production environment to query the API.

## Authentication
No authentication is required to use this API.

## Schema
The API provides the schema described below. The schema is used to query the API 
and to get the data you want. Continue reading to learn how to use the schema.

```
  type User {
    id: ID!
    name: String!
    email: String!
    pwdhash: String!
    role: Role!
    domain: Domain!
    permissions: [Permission!]!
  }

  type Role {
    id: ID!
    name: String!
    description: String!
    permissions: [Permission!]!
  }

  type Domain {
    id: ID!
    name: String!
    description: String!
  }

  type Permission {
    id: ID!
    name: String!
    description: String!
  }

  type Query {
    userById(id: ID!): User
    userByName(name: String!): User
    userByEmail(email: String): User
    roles: [Role!]!
    domains: [Domain!]!
    permissions: [Permission!]!
  }

  input NewUser {
    name: String!
    email: String!
  }

  input UpdateUser {
    name: String!
    email: String!
    pwdhash: String!
    role: ID!
    domain: ID!
    permissions: [ID!]!
  }

  type Mutation {
    createUser(input: NewUser!): User!
    updateUser(id: ID!, input: UpdateUser!): User!
    deleteUser(id: ID!): User!
  }
```

## Querying
The beauty of GraphQL is that you can query the API to get the data you want. 
Look at the schema above and you will see that the API provides seven queries 
that you can use to get the data you want.

The following methods exists for querying the API:

1) [userById](#1-userById)
2) [userByName](#2-userByName)
3) [userByEmail](#3-userByEmail)
4) [roles](#4-roles)
5) [domains](#5-domains)
6) [permissions](#6-permissions)

### 1) userById
This query returns a existing user.

#### Example

The request will be:

```
  query {
    userById(id:"b8b0c615-aa84-4e5c-bbf7-a53c181acd89") {
      id
      name
      email
    }
  }
```
The response will be:

```
  {
    "data": {
      "userById": {
        "id": "b8b0c615-aa84-4e5c-bbf7-a53c181acd89",
        "name": "Miguel Angelo Mello",
        "email": "miguelangelomello@gmail.com"
      }
    }
  }
```

## Mutation
The API provides three mutations that you can use to create, update and delete 
users. The mutations are described below.

1) [createUser](#1-createUser)
2) [updateUser](#2-updateUser)
3) [deleteUser](#3-deleteUser)

### 1) createUser
This mutation creates a new user.

#### Example

The request will be:

```
  mutation {
    createUser(input: {
      name: "Miguel Angelo Mello",
      email: "miguelangelomello@gmail.com"
    }) {
      id
      name
      email
    }
  }
```
The response will be:

```
  {
    "data": {
      "createUser": {
        "id": "381db2e8-920b-4f4d-96be-89ca8bfaa185",
        "name": "Miguel Angelo Mello",
        "email": "miguelangelomello@gmail.com"
      }
    }
  }
```

to be continued...

