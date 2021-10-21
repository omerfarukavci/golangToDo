## Go To Do App


**Server: Golang  
Client: React, semantic-ui-react  
Database: Local MongoDB**


# Highlights

1. DB connection string, name and collection name moved to `.env` file as environment variable. Using `github.com/joho/godotenv` to read the environment variables.
2. [feature/cloud-native-deployment](https://github.com/abdennour/go-to-do-app/tree/feature/cloud-native-deployment) provided by [abdennour](https://github.com/abdennour). Thank you [@abdennour](https://github.com/abdennour) to dockerize it. His features supports both Docker and Kubernetes.

## Application Requirement

### golang server requirement

1. gorilla/mux package for router `go get -u github.com/gorilla/mux`
2. mongo-driver package to connect with mongoDB `go get go.mongodb.org/mongo-driver`
3. github.com/joho/godotenv to access the environment variable.

### react client

From the Application directory

`create-react-app client`

## Start the application

1. Make sure your mongoDB is started
2. Create a `.env` file inside the `go-server` and copy the keys from `.env.example` and update the DB connection string.
3. From go-server directory, open a terminal and run
   `go run main.go`
4. From client directory,  
   a. install all the dependencies using `npm install`  
   b. start client `npm start`

## :panda_face: Walk through the application

Open application at http://localhost:3000

### Index page

![](https://github.com/schadokar/go-to-do-app/blob/master/images/index.PNG)

### Create task

Enter a task and submit

### A task has the following information:
ID: Automatic incremental.
Task Title: Task title
Status: In Progress, Urgent, Done, Overdue
Description: Task description. If there are expressions such as "Acil, ACİLLLLL" in the Description section, the status is set to Urgent! 
Category: General is the default category. Adding and deleting can be done.
Progress: Pending, On Track, Some Disruption, Discontinued, Completed
Deadline: End date.
CreatedAt: Created date (default: time.now())
UpdatedAt: Date updated
Remaining Day: deadline - createdAt 
# License

MIT License

Copyright (c) 2019 Shubham Chadokar
