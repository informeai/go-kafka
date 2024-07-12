clear:
	#kill containers
	- docker kill consumer producer kafka webui
	#remove containers
	- docker container rm consumer producer kafka webui
  #remove network
	- docker network rm netkafka
run:
	#create network
	docker network create netkafka
	#container kafka
	docker run -d --name kafka --network netkafka -e KAFKA_CFG_NODE_ID=0 -e KAFKA_CFG_PROCESS_ROLES=controller,broker -e KAFKA_CFG_LISTENERS=PLAINTEXT://:9092,CONTROLLER://:9093,EXTERNAL://:9094 -e KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,EXTERNAL:PLAINTEXT,PLAINTEXT:PLAINTEXT -e KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=0@kafka:9093 -e KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER bitnami/kafka:latest
	#container producer
	docker build -f Dockerfile.producer -t producer-image . && docker run -d --name producer --network netkafka producer-image
	#container consumer
	docker build -f Dockerfile.consumer -t consumer-image . && docker run -d --name consumer --network netkafka consumer-image
	#web ui kafka
	docker run -d --name webui --network netkafka -p 9000:9000 -e KAFKA_BROKERCONNECT=kafka:9092,kafka:9092 obsidiandynamics/kafdrop
