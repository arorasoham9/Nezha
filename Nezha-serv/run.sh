# docker stop $(docker ps -aq)  
# docker rm $(docker ps -aq)
# # go run main.go 
# # sleep 100
docker build -t arorasoham9/nezha-serv:latest .
docker push arorasoham9/nezha-serv:latest