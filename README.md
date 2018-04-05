# micro-pubsub
Testing pubsub with RabbitMQ and NATS. By default we use NATS, see below how to use RabbitMQ.

## NATS
``docker-compose up nats``

``docker-compose run service``

## RabbitMQ
Uncomment the RabbitMQ address and comment out the nats address. Also change the broker to ``rabbitmq``.

``docker-compose up rabbitmq``

``docker-compose run service``
