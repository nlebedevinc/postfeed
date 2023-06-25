Run connect and create Kafka topic:

> connect to Kafka instance
docker exec -it kafka bash

> create Kafka topic
kafka-topics.sh --zookeeper zookeeper:2181 --create --topic twitter.newTweets --replication-factor 1 --partitions 10
