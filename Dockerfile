FROM  ninetypoe-docker.jfrog.io/golang-build:builder-v3 as builder

WORKDIR /app
COPY ./ /app


RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./kafkanalysis-linux .


FROM edenhill/kafkacat:1.5.0

RUN apt-get update; \
    apt-get install -y jq grep bash


# Copy our static executable
COPY --from=builder ./app/kafkanalysis-linux /usr/bin/kafkanalysis
