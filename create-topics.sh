#!/bin/bash
/opt/kafka/bin/kafka-topics.sh --create --topic my-topic --partitions 3 --replication-factor 1 --bootstrap-server broker:9092
/opt/kafka/bin/kafka-topics.sh --create --topic another-topic --partitions 1 --replication-factor 1 --bootstrap-server broker:9092