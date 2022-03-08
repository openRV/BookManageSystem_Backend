defualt: run

.PHONY: run stress

run:
	go run main.go

stress:
	ab -n 1000 -c 1000 -p 'post.txt' -T 'application/x-www-form-urlencoded'  'http://localhost:8080/login'

