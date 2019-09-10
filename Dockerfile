FROM moh90poe/kafkacat:v1.4.0-avro

RUN apt-get update; \
    apt-get install -y jq grep

COPY ./artifacts/kafkanalysis-linux /usr/local/bin/kafkanalysis
