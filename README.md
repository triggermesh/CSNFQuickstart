# CSNF Quickstart
This is a quickstart for the CSNF project.

# Prerequisites

  1. [Docker Desktop](https://www.docker.com/products/docker-desktop)
  2. [Azure Sentinel](https://azure.microsoft.com/en-us/services/microsoft-sentinel/) with an App Registration (https://docs.microsoft.com/en-us/power-apps/developer/data-platform/walkthrough-register-app-azure-active-directory)that has been provided proper access to the Azure Sentinel resource.
  3. [Splunk](https://www.splunk.com/) configured with an [HEC collector](https://docs.splunk.com/Documentation/Splunk/8.2.6/Data/UsetheHTTPEventCollector)

# Getting Started

  1. Pull or Download this repository.

  2. Edit the `.env` file with the required Azure Sentinel and Splunk information.

  2. Open a terminal, from within the local repository, and run the following command:

  ```
  docker compsoe up
  ```

  3. View the results in Sentinel and Splunk.
