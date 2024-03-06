This repository contains code which fetchs data from sport news. 

1. Run: docker-compose up -d
2. Start applications: go run main.go

On-call Example App

This is a Go application that allows fetch news from given URL:
- Vorker that parses news from xml file and writes it to the database runs every two minutes. 
- To run the application you just need to pull up the docker compose file. 
- To view all news articles or single news articles, you need to call the appropriate endpoints.