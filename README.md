# Codership

Build requirements:
go 1.10
vgo latest
nodejs 8.10
chrome latest
docker latest

ToDo
 - ports 3306 nad 3307 should be open for integration tests to run (to be fixed to use random ports)
 - Integration and Unit tests should be executed separately
 - Do not trust user input. Validate the inputs.

Docker images:
published in public docker hub as:
- ufoscout/galera-mariadb:latest
- ufoscout/galera-mysql:latest

Available docker-compose file that shows how to use them
