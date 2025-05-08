go run github.com/99designs/gqlgen generate
go run server.go





query users {
  users {
    username
  } 
}


mutation createUser {
  createUser(input:
    {
   username: "T3R",
    email: "test@test.com",
    firstName: "Trev",
    lastName: "Test",}
  
  
  ) {
username
email
firstName
lastName
  }
}