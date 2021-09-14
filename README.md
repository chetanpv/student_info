**Student information** 
The mini-project is about playing around some student information. There are two containers which do the following.

**student-service:**
This service is responsibe for creating, updating and getting student information from/to the postgres database
The rest calls exposed in this service are give below

**1. Get student information from DB:**
curl --location --request GET 'http://localhost:8080/student'

**2. Create student information from DB:**
curl --location --request POST 'http://localhost:8080/student' \
--header 'Content-Type: application/json' \
--data-raw '[
    {
        "name": "Chetan",
        "branch": "ISE",
        "email": "chetan@gmail.com",
        "PhoneNumber": 12345,
        "id": "ise12"
    },
	    {
        "name": "Bindu",
        "branch": "ISE",
        "email": "bindu@gmail.com",
        "PhoneNumber": 56789,
        "id": "ise10"
]'

**3. Get a particular student information from DB:**
curl --location --request GET 'http://localhost:8080/student/1234'

**postgres-db-service:**
This service is responsible for storing/fetching information from the database.
The credentials and other data are hardcoded for now. This can be changed in future. Useful information:
user=postgres dbname=student password=postgres host=127.0.0.1 sslmode=disable port=5433

**Do get things to work. You can do it 3 ways.**
**1. Deploy via docker images**
You can use the docker files present and run the containers using the below commands
**student-service:**
docker build -t student-service-image:v1 .
docker run -d --name student-service -p 8080:8080 cstudent-service-image:v1

**postgres-db:**
docker build -t postgres-db-image:v1 .
docker run -d --name postgres-db-service -p 5433:5433 postgres-db-image:v1

**2. Deploy via docker compose**
You can find the docker-compose.yml in the git home directory. Run the below command for the containers to start running
docker-compose up

**3. You can get everything up via kubernetes**
You can find student-service-dep.yaml file in the git home directory. Follow the below commands
docker build -t student-service-image:v1 .
docker build -t postgres-db-image:v1 .
kubectl apply -f student-service-dep.yaml



