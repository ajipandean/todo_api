# run postgres:13-alpine container
sudo docker run -d -p 5432:5432 --name postgres_todo_api --env-file ./postgresql.conf postgres:13-alpine