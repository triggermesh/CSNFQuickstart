# CSNF Quickstart

This is the V1 quickstart for the [CSNF](https://github.com/onug/CSNF) project. In this version synthetic sources are used to emit 3 mock secuirty event's ( Oracle Cloudguard, Aquasec, and Microsoft Defender). These events are then transformed into the [CSNF](https://github.com/onug/CSNF) standarized format, and then sent to a self-contained Splunk enviorment.

The goal of this project is to quickly bootstrap a POC/DEV enviorment for teams and individuals intrested in the CSNF security event standardization.


## Prerequisites

  1. [Docker Desktop](https://www.docker.com/products/docker-desktop)

## Getting Started

  1. Pull or Download this repository.

  2. Open a terminal, from within the local repository, and run the following command:

  ```
  docker compose up
  ```

  3. View the results in in Splunk at `http://localhost:8000/en-US` by logging in with `admin` / `admin1234`


## Bringing Your Own Data

  This project also provides a way to bring your own data into the system via HTTP POST requests to `http://localhost:8080`. If you would like to use this feature, you will need to update the `WEBHOOK_EVENT_TYPE` in the `.env` file to reflect the type of data you are sending. The options are: `azure-defender-transformation`, `oracle-cloudguard-transformation`, or `aquasec-transformation`. The data you send will be transformed into the CSNF format and sent to both Splunk and Sentinel.

### Example usage

  For this example we will be using the `aquasec_new.json` sample data located in the `./samples` directory.

  1. Open the `.env` file and update the `WEBHOOK_EVENT_TYPE` to `aquasec-transformation`.

  2. Open a terminal, from within the local repository, and run the following command:

  ```
  docker compose up
  ```

  3. Open a second terminal and run the following command:
  ```
  curl -X POST -H "Content-Type: application/json" -d @./samples/aquasec_new.json http://localhost:8080
  ```

  4. View the results in Sentinel and Splunk.

## Contributing

  All of the underlying components are open source and can be found in the `src` directory (except for the Sentinal and Splunk components. They can be found at github.com/triggermesh/triggermesh).
