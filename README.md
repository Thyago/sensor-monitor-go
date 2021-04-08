# Sensor Monitor

Author:
[Thyago Clemente](https://www.linkedin.com/in/thyagotc/)

The goal of this project is to provide line managers a website to present a line chart which shows in a time range of 6 hours when the temperature of a Parin sensor has been above 84.3 degrees.



## Instructions


### Running on Docker
The project contains a docker folder containing a [docker-compose.yml](docker/docker-compose.yml) ready to run the project, which includes:
- A MySQL database container (db);
- A container to run the API (sensormonitor) and the background task that checks for sensors;
- A container to execute tests (sensormonitor_test);
- An Apache container to load the simple html interface (sensormonitorview) containing a bar chart.

To run the API together with the database and interface, execute:
```sh
docker-compose --profile dev up
```

To run only the tests together with the database, execute:
```sh
docker-compose --profile test up
```

### Manual Setup

First setup the required environment variables:
| ENV VARIABLE                 | DESCRIPTION                           | DEFAULT |
| ---------------------------- | :------------------------------------ | :-----: |
| DB_USER                      | MySQL Username                        |         |
| DB_PASSWORD                  | MySQL Password                        |         |
| DB_HOST                      | MySQL Host                            |         |
| DB_PORT                      | MySQL Port                            | 3306    |
| DB_NAME                      | MySQL Database Name                   |         |
| SERVER_PORT                  | API Server Default Port               | 8080    |
| PARIN_API_KEY                | The Parin API Key                     |         |
| PARIN_SENSOR_CHECK_FREQUENCY | The frequency to run the sensor check |         |


Then run it from the root directory:
```sh
go run main.go
```

To run the unit tests, execute 
```sh
go test --tags=unit -v ./...
```

To run the integration tests, execute 
```sh
go test --tags=integration -v ./...
```


## API Documentation

An OpenAPI (Swagger) specification is not yet available for the project.

Check [controllers/server.go](server.go) to check the endpoints list.

Current version uses `/v1` prefix for all endpoints.



## Third Party License Attributions
See [ATTRIBUTIONS.md](ATTRIBUTIONS.md)

## License
The MIT License (MIT)

Copyright (c) 2021 Thyago Clemente

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.