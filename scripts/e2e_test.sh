cd tests/e2e
docker-compose up -d

# just to make sure this doesn't keep running
sleep 10 && docker-compose down