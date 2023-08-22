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

# User Domain Role Service GraphQL API Reference

Version 1


## 1) Introduction
This API reference provides information about the functionality of the User Domain Role Service GraphQL API. The User Domain Role Service can provide a centralized and unified place to store and manage user roles across different domains or applications within a corporate environment. The API provides functionality described below.


## 2) Authentication
No authentication is required to use this API, yet it may be required in the future.


## 3) Routes
The following routes are available:

```
  Query endpoint:       http://udrs.orionsoft.site/query
  Playground endpoint:  http://udrs.orionsoft.site/playground
  Reference endpoint:   http://udrs.orionsoft.site/
```

### 3.1) Query endpoint
The query endpoint is used to query the API. The query endpoint is the only endpoint that you will use in production. When querying the API, keep in mind the following:

```
  1) The query endpoint is a POST endpoint.
  2) The query endpoint requires a valid GraphQL query in the request body.
  3) The query endpoint requires the following headers: Content-Type: application/json
```

### 3.2) Playground endpoint
The playground endpoint is used to query the API from a web browser. The playground endpoint is a GET endpoint that provides a graphical interface to query the API. The playground endpoint is not recommended to be used in production. When querying the API, keep in mind the following:

```
  1) The playground endpoint is a GET endpoint.
  2) The playground endpoint requires a valid GraphQL query in the request body.
```

### 3.3) Reference endpoint
The reference endpoint is used to get information about the API. All the information you need to use the API is provided in this document. 


## 4) Schema
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

## 5) Querying
The beauty of GraphQL is that you can query the API to get the data you want. Look at the schema above and you will see that the API provides seven queries that you can use to get the data you want.

The following methods exists for querying the API:

```
  5.1) userById     //Returns a existing user by its Id.
  5.2) userByName   //Returns a existing user by its Name.
  5.3) userByEmail  //Returns a existing user by its Email.
  5.4) roles        //Returns all existing roles
  5.5) domains      //Returns all existing domains
  5.6) permissions  //Returns all existing permissions
```

### 5.1) userById
Returns a existing user by its Id.

#### 5.1.1) Example

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

## 6) Mutation
The API provides three mutations that you can use to create, update and delete 
users. The mutations are described below.

```
  6.1) createUser   //Creates a new user
  6.2) updateUser   //Updates an existing user
  6.3) deleteUser   //Deletes an existing user
```

### 6.1) createUser
This mutation creates a new user.

#### 6.1.1) Example

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


## 7) Rate Limiting
The API imposes a rate limit of 100 requests per minute per user. If the rate limit is exceeded, a `429 Too Many Requests` response will be returned.


## 8) Change Log
v1 (2023-07-16): &nbsp;Initial release of the API.

## 9) Support and Contact Information
For any questions or issues, please contact Miguel Mello at miguelangelomello@gmail.com.

