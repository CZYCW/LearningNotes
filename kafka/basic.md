
### why should a company use kafka
- Apache Kafka is one of the most popular data streaming processing platforms in the industry
- Kafka provides a simple message queue interface on top of its append-only log-structured storage medium
- Data is distributed to multiple nodes. Kafka is highly scalable and fault-tolerant to node loss.


### use cases
- Messaging systems
- Activity Tracking
- Gather metrics from many different locations, for example, IoT devices
- Application logs analysis
- De-coupling of system dependencies
- Integration with Big Data technologies like Spark, Flink, Storm, Hadoop.
- Event-sourcing store


### topic
- Kafka topics organize related events
- The data in the topics are stored in the key-value form in binary format.

### producer
- producer send data into the topic

### consumer
- Applications that pull event data from one or more Kafka topics are known as Kafka consumers. 
- By default, Kafka consumers will only consume data that was produced after it first connected to Kafka. Which means that to read historical data in Kafka, one must specify it as an input to the command, as we will see in the practice section.
- Delivery semantics for consumers
    - At most once:
    - At least once
    - Exactly once

### broker
- A single Kafka server is called a Kafka Broker. That Kafka broker is a program that runs on the Java Virtual Machine (Java version 11+) and usually a server that is meant to be a Kafka broker will solely run the necessary program and nothing else.

- leader broker: one Kafka broker is designated by the cluster to be responsible for sending and receiving data to clients.

