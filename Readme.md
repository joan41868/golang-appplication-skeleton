## Go Application Skeleton

### Short description:
The project's aim is to give a base application boilerplate for new go projects.
The template is suitable for many types of web projects.

### This is a sample skeleton for go applications. It includes:
- Model layer - for defining persistable application entities
- Repository layer - for persisting said entities. Not implemented, but rather suggested via interface
- HTTP Server layer - for exposing the API
- Application layer - for grouping components together


### Major TODOs:
- [ ] Service Layer
- [ ] Security Layer
- [ ] gRPC Server

### What other packages are present here:
- `gorilla/mux` - the fancy router
- `gorm` - ORM for working with relational databases
- `godotenv` - for loading .env files