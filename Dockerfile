FROM  ninetypoe-docker.jfrog.io/golang-build:builder-v3 as builder

WORKDIR /app
COPY ./ /app


RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./kafkanalysis-linux .


FROM moh90poe/kafkacat:v1.4.0-avro

RUN apt-get update; \
    apt-get install -y jq grep


# Copy our static executable
COPY --from=builder ./artifacts/kafkanalysis-linux /kafkanalysis
COPY --from=builder ./VERSION /
