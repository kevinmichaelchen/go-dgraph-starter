# user-api
This project functions as a store for users.

The data model and API surface area is very simple.

We use gRPC, Postgres, and sqlboiler as an ORM.

## Getting started
### Running migrations
```
make migrate
```

### Auto-generating sqlboiler models
```
make boil
```