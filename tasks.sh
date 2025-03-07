#!/bin/bash

API_URL=http://localhost:8080
DB_URL=postgres://skillsrock:secret@localhost:5432/tasks

case "$1" in
    start)
		docker compose up -d && goose postgres -dir sql/schema $DB_URL up
        ;;
    stop)
		docker compose down
        ;;
    bench)
        ab -n 10000 -c 1000 $API_URL/tasks
        ;;
    clean)
		docker compose down && docker image rm tasks-server && docker volume rm skillsrock-tasks_postgres_data
        ;;
    
    *)
        echo "Usage: $0 {start|stop|bench|clean"
        exit 1
esac
