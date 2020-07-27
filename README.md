### About
A url shortner build with go. I build this to learn go and the fiber webframework.

Technology used for this project:

- Go as the programming language for the backend
- Fiber as webframework for the go language
- MySQL for storring the URLS
    - a json file would have probably been enough but wanted to test this for anther project

### Installation 
Make sure you add a .env file with the following content at the root of the project.
```env
MYSQL_DATABASE=db_name
MYSQL_ROOT_PASSWORD=password
MYSQL_USER=user
MYSQL_PASSWORD=password
```
To run this project make sure docker and docker-compose is installed.
Simply then run
```bash
docker-compose up --build
```
The build flag is for building and is only needed the first time.


### How it works
When you saved an url there be json returned. To get the real url again just go to
/redirect/{id}
Where id is the id saved from the database you just got.