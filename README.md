<!-- *




go run github.com/99designs/gqlgen generate
go run server.go

docker compose run --service-ports web bash





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






docker container create [OPTIONS] IMAGE [COMMAND] [ARG...]

docker exec pid_id bashcommandincontainer

docker run -d 
-p 5001:5000  
--name containername 
--v /tmp/container:/tmp/container 

example:
docker run --name -v website $PWD/website:/usr/share/nginx/html -p 8080:80




docker rm -f nameofcontainer

--interactive --tty

things to read later
---------------------
[server mesh vs message queue vs](https://arcentry.com/blog/api-gateway-vs-service-mesh-vs-message-queue/) 


* -->