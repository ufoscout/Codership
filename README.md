# Codership - Galera Cluster

## The task

Write a server program in Go that
- installs a 3-node MySQL/Galera cluster in an environment of your choice: AWS, Docker, VMs, single host, etc.
- starts the cluster up
- waits for HTTP client connections
- displays a list of running nodes on the home page.

## The implementation

The application is implemented as SPA (Single Page Application) with:
- backend: written in Go, it exposes the required features as REST endpoints;
- frontend: written in TypeScript, NodeJs and Angular 6. It is a Single Page Application that uses the backend endpoints;
- the cluster nodes are deplyoyed as docker containers on the local machine.

## The Docker images
The backend uses custom Docker images whose source code is available in the 'docker' folder. 
These images are pushed to the central docker hub and available at:
- https://hub.docker.com/r/ufoscout/galera-mariadb/
- https://hub.docker.com/r/ufoscout/galera-mysql/

These images were created explicitly for this coding task. They are downloaded automatically by the backend when a cluster is deployed for the first time.
They are also downloaded during the execution of the backend tests.

WARNING: the MySql images are much slower than the MariaDB ones. The problem is caused by the fact that the nodes synchronization is performed more than once per node (caused by a two phases initialization). To be fixed.


## The projet structure
```
Codership
  |- backend         -> Go code for the backend
  |     |- build.sh  -> build script for the backend
  |- docker
  |     |- mariadb   -> Dockerfile for the MariaDB Galera Cluster node
  |     |- mysql     -> Dockerfile for the MySQL Galera Cluster node
  |- frontend        -> NodeJs/Angular/Typescript code for the frontend
  |     |- build.sh  -> build script for the frontend
  |- build.sh        -> this script builds the whole project calling the backend and frontend scripts
```

## How to build
### Build requirements
The following tools are required to build the project:

For building the backend:
- go 1.10 
- vgo latest (see https://github.com/golang/vgo )
- Docker latest (required at runtime to deploy the clusters and to execute the integration tests)

For building the frontend:
- nodejs 8.10
- Chrome latest (not required to build, required only to execute the frontend tests)

### Runtime requirements
Once the project is built, it will produce a native executable that has no specific requirements. Even NodeJS is not required at runtime but only to build the frontend.
The only real runtime requirement is to have Docker properly installed.

WARNING: the project was developed and tested in a Linux environment (Ubuntu 18.04 64 bit). As it requires native Docker, it could not work as expected with non native Docker installations (e.g. Docker machine) or on non-linux OSes.

### Perform the build on Ubuntu linux
To build the project execute the "build.sh" script on the main folder. It will:
- execute the backend tests
- build the backend
- execute the frontend tests
- build the frontend

### Perform the build on other OSes
WARNING: this process was not tested!
Steps:
- enter the "backend" folder
- open a terminal
- launch the build: `vgo build`
- enter the "frontend" folder
- open a terminal
- launch the build: `npm run build`

## Start the application
When the build is completed, open a terminal in the 'backend' folder and start the 'backend' executable file ('backend.exe' in Windows - tested only on Linux though).
If everything went fine, the application will be available at http://localhost:8080

## Warning
The first time a cluster is created, or the first time the backend tests are executed, the process could take several minutes as it has to download the MySql and MariaDB custom images from Docker Hub.

## The UI
The UI is available at the URL http://localhost:8080 (tested only with Firefox). It displays the list of already created clusters if any.

Click on the "Create new cluster" button to create a new cluster.

The button will open a modal window in which the following options are available:
- *Deployment Type* : it could be "Docker", "Kubernates" or "Ansible". Only "Docker" is implemented, however, it shows how the backend code is independent from the real implementation which is selected at runtime.
- *Database Type* : it could be "MySql" or "MariaDB", this is the database type that will be used to create the cluster.
- *Cluster Name* : It should be a unique name with no white spaces (input validation is not implemented)
- *Cluster Size* : the number of nodes of the cluster. Select "3" to create a cluster with three nodes.
- *First node Port* : This specifies the host port used by the first node of the cluster; the other nodes will uses progressive port numbers starting from this one. For example, if cluster size is 4 and first-node-port is 3306, then the cluster nodes will use the following ports:
  - first node: port 3306
  - second node: port 3307
  - third node: port 3308
  - fourth node: port 3309

## ToDo
The application is only a small POC, even if the backend has a good test coverage, it is not intended to be used in production environments.
ToDos and improvements:
 - The UI should show the progress of the operations (currently the first time a cluster is created it takes long time to download the docker images and it seems that the application is stuck) 
 - State is not persisted. A database should be used to persist the clusters states and configurations. Currently they are only saved in the browser local storage.
 - The backend Integration and Unit tests should be executed separately
 - The backend integration tests use the two hardcoded ports 12306 and 12307. If those ports are not available the tests fail, random free ports should be used instead.
 - The user input should be validated both on the frontend and the backend.
 - The internationalization is not complete (there is some hardcoded English text) and not correct (Finnish translations taken from google translator)
- The frontend coverage is not optimal and there are no e2e tests