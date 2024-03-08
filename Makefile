deploy:
	cd ./backend && make build
	docker-compose up -d --build