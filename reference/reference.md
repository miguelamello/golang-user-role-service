<style>
  * {
    font-family: Arial, Helvetica, sans-serif;
  }

  body {
    padding: 3em 5em;
    background-color: whitesmoke;
  }

  pre {
    background-color: #eee;
    padding: 10px;
    border-left: 5px solid #ccc;
    font-family: 'Courier New', Courier, monospace;
    visibility: scroll;
  }

  code {
    font-family: 'Courier New', Courier, monospace;
  }
</style>

# API Reference for User Domain Role Service

## 1) Description
This API reference provides information about the functionality of the User Domain Role Service GraphQL API. The User Domain Role Service can provide a centralized and unified place to store and manage user roles across different domains or applications within a corporate environment. The API provides functionality described below.

## 2) Version
Current API Version is: `v1`


## 3) Servers
The API is hosted on the following servers:

```
  Production server: http://orionsoft.site/udsr/v1/query
  Development server: http://localhost:8080/udsr/query
```

## 4) Endpoints
Differently from REST APIs, GraphQL APIs have only one endpoint where all the 
requests are sent using POST. The endpoint for this API is:

```
  Production endpoint: http://orionsoft.site/udsr/v1/query
  Development endpoint: http://localhost:8080/udsr/v1/query
```

### 4.1) Methods
The API provides only one method: `POST`

### 4.2) Headers
The API requires the following headers:

`Content-Type: application/json`

### 4.3) Request Body
 The request body is required and must be a valid GraphQL query. 

#### 4.3.1) Example

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

#### 4.3.1) For development:
We recommend you to use a GraphQL desktop client to query the API. Usually, the desktop client will provide a GraphQL body to query the API.

Obs: You may want to access the GraphQL playground via browser at:

```
  Production endpoint: http://orionsoft.site/udsr/v1/playground
  Development endpoint: http://localhost:8080/udsr/v1/playground
```

The playground will provide a graphical interface to query the API. While the playground is funny and good to get used to GraphQL idiom, it is not recommended to use it in production.

#### 4.3.2) For production:
We recommend you to use a GraphQL client module suitable to your programing language and production environment to query the API.

## 5) Authentication
No authentication is required to use this API, yet it may be required in the future.

## 6) Schema
The API provides the schema described below. The schema is used to query the API and to get the data you want. Continue reading to learn how to use the schema.

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

## 7) Querying
The beauty of GraphQL is that you can query the API to get the data you want. Look at the schema above and you will see that the API provides seven queries that you can use to get the data you want.

The following methods exists for querying the API:

```
  7.1) userById     //Returns a existing user by its Id.
  7.2) userByName   //Returns a existing user by its Name.
  7.3) userByEmail  //Returns a existing user by its Email.
  7.4) roles        //Returns all existing roles
  7.5) domains      //Returns all existing domains
  7.6) permissions  //Returns all existing permissions
```

### 7.1) userById
Returns a existing user by its Id.

#### 7.1.1) Example

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

## 8) Mutation
The API provides three mutations that you can use to create, update and delete 
users. The mutations are described below.

```
  8.1) createUser   //Creates a new user
  8.2) updateUser   //Updates an existing user
  8.3) deleteUser   //Deletes an existing user
```

### 8.1) createUser
This mutation creates a new user.

#### 8.1.1) Example

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

