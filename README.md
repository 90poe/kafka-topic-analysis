# kafka-topic-analysis

### Code coverage:

```
➜ go test ./... -coverprofile=cover.out && go tool cover -html=cover.out  
?       kafka-topic-analysis                            [no test files]
?       kafka-topic-analysis/data                       [no test files]
?       kafka-topic-analysis/env                        [no test files]
ok      kafka-topic-analysis/mathematicalfunctions      0.015s  coverage: 96.6% of statements
ok      kafka-topic-analysis/topics                     0.018s  coverage: 95.5% of statements
```

### How to use:
* Get the data from the topic using kafkacat docker image
> NOTE: When running the below command make sure you are connected to the vpn and all certs and necessary fields have been referenced.
> To view the `kafkacat` help menu run `docker run -t edenhill/kafkacat:1.5.0` which will explain what flags a expected and what they mean.
```
➜ docker run -v /Users/moh/kafka/dev-test-stg/:/src -t edenhill/kafkacat:1.5.0 -C -b kafka-dev-test-stage-cxijq-dev-test-stg-env.aivencloud.com:10497 -X security.protocol=SSL -X ssl.key.location=/src/service.key -X ssl.key.password=30022601ham -X ssl.certificate.location=/src/service.cert -X ssl.ca.location=/src/ca.pem -s value=avro -r https://avnadmin:m9x4v4aaj94k7pib@kafka-dev-test-stage-cxijq-dev-test-stg-env.aivencloud.com:10500 -t in_iot_vessel_sensor_oktopus_yoctopuce_gyroscope -p 14 -o beginning  -J -Z | jq . > partition14
```
> Use `-v` to mount the directory containing certs and keys to a docker volume, run the container in terminal mode using `-t`,
> the entrypoint of the image is the `kafkacat` binary so you can pass in the bootstrap broker , and cert keys by referencing them 
> from the volume you mounted to your container, define the topic and partition, use `-J` to output Newline-delimited JSON 
> and use `-Z` to output null values as `NULL` so we can capture this in the analysis, then be using `jq .` we can pipe the 
> output to a file which we can analyse 

* Once you have the data you can then analyse it using the `kafkanalysis` binary
```
➜ kafkanalysis                                                                                         
2019/09/20 11:51:56 Unknown operation: 
      --create-table           Creates a table
      --csv-filename string    The CSV file to output to, the default is output.csv (default "output.csv")
  -j, --json-filepath string   The path to the JSON output file from kafkacat
  -o, --operation string       Operation to perform:
                                 analyse - Analyse a set of results and the time gaps between the data
                                 describe - Describe the application
      --remove-duplicates      Remove the duplicate messages from the dataset
      --toCSV                  Output to a CSV File
  -v, --version                Prints the current version
```

example:
1. Get a small data sample, and verify the data is in the file
    ```
    ➜ docker run -v /Users/moh/kafka/dev-test-stg/:/src -t edenhill/kafkacat:1.5.0 -C -b kafka-dev-test-stage-cxijq-dev-test-stg-env.aivencloud.com:10497 -X security.protocol=SSL -X ssl.key.location=/src/service.key -X ssl.key.password=30022601ham -X ssl.certificate.location=/src/service.cert -X ssl.ca.location=/src/ca.pem -s value=avro -r https://avnadmin:m9x4v4aaj94k7pib@kafka-dev-test-stage-cxijq-dev-test-stg-env.aivencloud.com:10500 -t in_iot_vessel_sensor_oktopus_yoctopuce_gyroscope -p 14 -c 10000 -o -10000  -J -Z | jq . > out
    
    ~/topic-analysis at ☸️  minikube took 5s 
    ➜ cat out
    {
      "topic": "in_iot_vessel_sensor_oktopus_yoctopuce_gyroscope",
      "partition": 14,
      "offset": 225170,
      "tstype": "create",
    ...
    ...
    ```

1. What the below command does is analyse 10,000 messages for in_iot_vessel_sensor_oktopus_yoctopuce_gyroscope on Partition 14,
   the `--remove-duplicates` flag removed al the duplicate messages that are sent to the topic so they dont skew the results of the analysis, 
   the `--create-table` flag outputs these messages in a table format (which has been truncated fyi) for you to see and the `--toCSV` flag
   outputs the data to a csv file `output.csv by default`(if you dont specify a name), only 3 duplicate messages where removed from the 10,000 set of messages.
   
   At the bottom is a basic statistical analysis of values, where the Mean, Mode, Median, Max, Min and Standard Deviation, as well as some probablity of values
   over a time period using Poisson Distribution techniques. 
    ```
    ➜ kafkanalysis -o analyse -j ./out --remove-duplicates --create-table --toCSV               
    2019/09/20 12:04:01 Analysing in_iot_vessel_sensor_oktopus_yoctopuce_gyroscope in Partition 14
    2019/09/20 12:04:01 Removing Duplicates...
    2019/09/20 12:04:01 Function removed 3 messages from the Dataset
    +----------------------------------------------------+--------------------+---------------------+---------------------+---------------------+--------------------+---------------------+----------------------+
    |                     EVENT TIME                     | INTERVAL (SECONDS) |        GYRO         |    MAGNETOMETER     |       COMPASS       |   ACCELEROMETER    |        TILTX        |        TILTY         |
    +----------------------------------------------------+--------------------+---------------------+---------------------+---------------------+--------------------+---------------------+----------------------+
    | 13 Sep 19 12:24 UTC                                | N/A                |                   0 | 0.16699999570846558 |  59.900001525878906 | 0.9819999933242798 |  -57.20000076293945 |   -1.899999976158142 |
    +----------------------------------------------------+--------------------+---------------------+---------------------+---------------------+--------------------+---------------------+----------------------+
    | 13 Sep 19 12:26 UTC                                |                120 | 0.10000000149011612 | 0.15800000727176666 |  59.900001525878906 | 0.9829999804496765 |  -57.20000076293945 |   -1.899999976158142 |
    +----------------------------------------------------+--------------------+---------------------+---------------------+---------------------+--------------------+---------------------+----------------------+
    | 13 Sep 19 12:26 UTC                                |                  0 |                   0 | 0.15399999916553497 |               193.5 | 0.9739999771118164 | -2.5999999046325684 |   -6.900000095367432 |
    +----------------------------------------------------+--------------------+---------------------+---------------------+---------------------+--------------------+---------------------+----------------------+
    | 13 Sep 19 12:28 UTC                                |                120 |                   0 |  0.1469999998807907 |  59.900001525878906 | 0.9819999933242798 |  -57.20000076293945 |   -1.899999976158142 |
    +----------------------------------------------------+--------------------+---------------------+---------------------+---------------------+--------------------+---------------------+----------------------+
    | 13 Sep 19 12:28 UTC                                |                  0 | 0.20000000298023224 | 0.13099999725818634 |   192.3000030517578 | 0.9589999914169312 |  -2.799999952316284 |                 -7.5 |
    +----------------------------------------------------+--------------------+---------------------+---------------------+---------------------+--------------------+---------------------+----------------------+
    | 13 Sep 19 12:30 UTC                                |                120 | 0.20000000298023224 |  0.2529999911785126 |   191.8000030517578 | 0.9610000252723694 | -2.9000000953674316 |   -5.400000095367432 |
    +----------------------------------------------------+--------------------+---------------------+---------------------+---------------------+--------------------+---------------------+----------------------+
    | 13 Sep 19 12:30 UTC                                |                  0 |                   0 |  0.1469999998807907 |  59.900001525878906 |  0.984000027179718 |  -57.20000076293945 |   -1.899999976158142 |
    +----------------------------------------------------+--------------------+---------------------+---------------------+---------------------+--------------------+---------------------+----------------------+
    ...
   2019/09/20 12:04:03 The file has bee successfully created: output.csv
   
   	----------------------------------- Basic Stats Analysis: Time Intervals -----------------------------------
   	Time interval Mean: 60 Seconds
   	Time interval Mode: 0 Seconds
   	Time interval Median: 60 Seconds
   	Time interval Max Value: 120 Seconds
   	Time interval Min Value: 0 Seconds
   	Time interval Standard Deviation: 60 Seconds
   	
   	
   	Poisson Distribution: Time Intervals
   	The probability of a time interval being Greater than 2 minutes is 0.36787944264296 
   	The probability of a time interval being Less than 2 minutes is 0.63212055735704
   	The probability of a data interval being Between 2 & 4.5 minutes is 0.2624802171325027
   	
   	----------------------------------- Basic Stats Analysis: Gyro Values -----------------------------------
   	Gyro Mean: 0.08078424 
   	Gyro Mode: 0.10000000149011612 
   	Gyro Median: 0.10000000149011612 
   	Gyro Max Value: 1.7000000476837158 
   	Gyro Min Value: 0 
   	Gyro Standard Deviation: 0.11868484 
   	
   	
   	Poisson Distribution: Gyro Values
   	The probability of a Gyro being Greater than 0.3 is 0.9975031224074351 
   	The probability of a Gyro being Less than 0 is 0
   	The probability of a Gyro being Between 0 & 0.3 is 0.0024968775925648945
   	
   	----------------------------------- Basic Stats Analysis: Magnetometer Values -----------------------------------
   	Magnetometer Mean: 0.26321566 
   	Magnetometer Mode: 0.21299999952316284 
   	Magnetometer Median: 0.23899999260902405 
   	Magnetometer Max Value: 1.1519999504089355 
   	Magnetometer Min Value: 0.09099999815225601 
   	Magnetometer Standard Deviation: 0.10197296 
   	
   	
   	Poisson Distribution: Magnetometer Values
   	The probability of a Magnetometer being Greater than 1 is 0.9917012926719326 
   	The probability of a Magnetometer being Less than 0.5 is 0.004157998138292651
   	The probability of a Magnetometer being Between 0.5 & 1 is 0.004140709189774716
   	
   	----------------------------------- Basic Stats Analysis: Compass Values -----------------------------------
   	Compass Mean: 262.47943 
   	Compass Mode: 340.6000061035156 
   	Compass Median: 340.3999938964844 
   	Compass Max Value: 359.8999938964844 
   	Compass Min Value: 0.10000000149011612 
   	Compass Standard Deviation: 121.800705 
   	
   	
   	Poisson Distribution: Compass Values
   	The probability of a Compass being Greater than 200 is 0.1888756040967325 
   	The probability of a Compass being Less than 150 is 0.713495201707286
   	The probability of a Compass being Between 150 & 200 is 0.09762919419598148
   	
   	----------------------------------- Basic Stats Analysis: Accelerometer Values -----------------------------------
   	Accelerometer Mean: 0.967954 
   	Accelerometer Mode: 0.9750000238418579 
   	Accelerometer Median: 0.9710000157356262 
   	Accelerometer Max Value: 0.9860000014305115 
   	Accelerometer Min Value: 0.9369999766349792 
   	Accelerometer Standard Deviation: 0.008511713 
   	
   	
   	Poisson Distribution: Accelerometer Values
   	The probability of a Accelerometer being Greater than 5 is 0.9591894572690031 
   	The probability of a Accelerometer being Less than 0.2 is 0.0016652785424057237
   	The probability of a Accelerometer being Between 0.2 & 5 is 0.03914526418859121
   	
   	----------------------------------- Basic Stats Analysis: TiltX Values -----------------------------------
   	TiltX Mean: -29.796638 
   	TiltX Mode: -57 
   	TiltX Median: -55.29999923706055 
   	TiltX Max Value: -1.7999999523162842
   	TiltX Min Value: -59 
   	TiltX Standard Deviation: 27.244478 
   	
   	
   	----------------------------------- Basic Stats Analysis: TiltY Values -----------------------------------
   	TiltY Mean: -3.3177853 
   	TiltY Mode: -1.7000000476837158 
   	TiltY Median: -3.700000047683716 
   	TiltY Max Value: 3.5
   	TiltY Min Value: -24.799999237060547 
   	TiltY Standard Deviation: 1.9053484 
   
    ```
