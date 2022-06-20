# CSNF Quickstart
This is a quickstart for the CSNF project.

## Prerequisites

  1. [Docker Desktop](https://www.docker.com/products/docker-desktop)
  2. [Azure Sentinel](https://azure.microsoft.com/en-us/services/microsoft-sentinel/) with an App Registration (https://docs.microsoft.com/en-us/power-apps/developer/data-platform/walkthrough-register-app-azure-active-directory)that has been provided proper access to the Azure Sentinel resource.
  3. [Splunk](https://www.splunk.com/) configured with an [HEC collector](https://docs.splunk.com/Documentation/Splunk/8.2.6/Data/UsetheHTTPEventCollector)

## Getting Started

  1. Pull or Download this repository.

  2. Edit the `.env` file with the required Azure Sentinel and Splunk information.

  2. Open a terminal, from within the local repository, and run the following command:

  ```
  docker compose up
  ```

  3. View the results in Sentinel and Splunk.


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

  All of the underlying components are open source and can be found in the `guts` directory (except for the Sentinal and Splunk components. They can be found at github.com/triggermesh/triggermesh).
