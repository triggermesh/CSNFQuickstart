# CSNF Quickstart

This is the V1 quickstart for the [CSNF](https://github.com/onug/CSNF) project. In this version synthetic sources are used to emit 3 mock secuirty event's ( Oracle Cloudguard, Aquasec, and Microsoft Defender). These events are then transformed into the [CSNF](https://github.com/onug/CSNF) standarized format, and then sent to a self-contained Splunk enviorment.

The goal of this project is to quickly bootstrap a POC/DEV enviorment for teams and individuals intrested in the CSNF security event standardization.


## Prerequisites

  1. [Docker Desktop](https://www.docker.com/products/docker-desktop)
  2. X86 chipset (ARM is un-supported at the moment due to an issue with Splunk)

## Getting Started

  1. Pull or Download this repository.

  2. Open a terminal, from within the local repository, and run the following command:

  ```
  docker compose up
  ```

  3. View the results in in Splunk at `http://localhost:8000/en-US` by logging in with `admin` / `admin1234`


## Contributing

  All of the underlying components are open source and can be found in the `src` directory (except for the Splunk Target component. This can be found at github.com/triggermesh/triggermesh).
