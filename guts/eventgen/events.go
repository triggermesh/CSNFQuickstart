package main

var AzureEventChain = []Event{
	{
		EventType: "com.microsoft.azure.defender.alert",
		EventData: `{
			"id": "/subscriptions/97e01fd4-3326-41f4-b9e3-3cfd6809e10f/resourceGroups/Sample-RG/providers/Microsoft.Security/locations/centralus/alerts/2517538088322968242_7951468c-3909-4b52-a442-c1f4b92d5162",
			"name": "2517538088322968242_7951468c-3909-4b52-a442-c1f4b92d5162",
			"type": "Microsoft.Security/Locations/alerts",
			"properties": {
			  "vendorName": "Microsoft",
			  "alertDisplayName": "[SAMPLE ALERT] Unusual amount of data extracted from a storage account",
			  "alertName": "SIMULATED_Storage.Blob_DataExfiltration.AmountOfDataAnomaly",
			  "detectedTimeUtc": "2022-03-28T18:26:07.7031757Z",
			  "description": "THIS IS A SAMPLE ALERT: Someone has extracted an unusual amount of data from your Azure Storage account 'Sample-Storage'.",
			  "remediationSteps": "• Limit access to your storage account, following the 'least privilege' principle: https://go.microsoft.com/fwlink/?linkid=2075737.\r\n• Revoke all storage access tokens that may be compromised and ensure that your access tokens are only shared with authorized users.\r\n• Ensure that storage access tokens are stored in a secured location such as Azure Key Vault. Avoid storing or sharing storage access tokens in source code, documentation, and email.",
			  "actionTaken": "Undefined",
			  "reportedSeverity": "High",
			  "compromisedEntity": "Sample-Storage",
			  "associatedResource": "/SUBSCRIPTIONS/97e01fd4-3326-41f4-b9e3-3cfd6809e10f/RESOURCEGROUPS/Sample-RG/providers/Microsoft.Storage/storageAccounts/Sample-Storage",
			  "subscriptionId": "97e01fd4-3326-41f4-b9e3-3cfd6809e10f",
			  "instanceId": "7951468c-3909-4b52-a442-c1f4b92d5162",
			  "extendedProperties": {
				"resourceType": "Storage",
				"investigation steps": "{\"displayValue\":\"View related storage activity using Storage Analytics Logging. See how to configure Storage Analytics logging and more information.\",\"kind\":\"Link\",\"value\":\"https:\\/\\/go.microsoft.com\\/fwlink\\/?linkid=2075734\"}",
				"potential causes": "This alert indicates that an unusually large amount of data has been extracted compared to recent activity on this Storage container.\r\nPotential causes:\r\n• An attacker has extracted a large amount of data from a Storage container (for example: data exfiltration/breach, unauthorized transfer of data).\r\n• A legitimate user or application has extracted an unusual amount of data from a Storage container (for example: maintenance activity).",
				"client IP address": "00.00.00.00",
				"client location": "Azure Data Center: East Us",
				"authentication type": "Anonymous",
				"operations types": "GetBlob",
				"service type": "Azure Blobs",
				"user agent": "dummyAgent",
				"container": "eicarTestStorageContainer",
				"extracted data": "140 MB",
				"test: Slice start time": "03/28/2022 18:26:07",
				"test: Pipeline name": "1.0.4656.1_storagetd-brs-a3",
				"extracted blobs": "500",
				"killChainIntent": "Exfiltration"
			  },
			  "state": "Active",
			  "reportedTimeUtc": "2022-03-28T18:26:47.1036441Z",
			  "workspaceArmId": "/subscriptions/97e01fd4-3326-41f4-b9e3-3cfd6809e10f/resourcegroups/csnf/providers/microsoft.operationalinsights/workspaces/csnfsentinel",
			  "confidenceReasons": [],
			  "canBeInvestigated": true,
			  "isIncident": false,
			  "entities": [
				{
				  "$id": "centralus_1",
				  "address": "00.00.00.00",
				  "location": {
					"countryName": "United States",
					"city": "Washington"
				  },
				  "type": "ip"
				}
			  ]
			}
		  }`,
	},
	{
		EventType: "com.microsoft.azure.defender.alert",
		EventData: `{
			"id": "/subscriptions/97e01fd4-3326-41f4-b9e3-3cfd6809e10f/resourceGroups/Sample-RG/providers/Microsoft.Security/locations/centralus/alerts/2517538088722812591_ee939333-c75e-461d-83c2-712ba9abfadb",
			"name": "2517538088722812591_ee939333-c75e-461d-83c2-712ba9abfadb",
			"type": "Microsoft.Security/Locations/alerts",
			"properties": {
			  "vendorName": "Microsoft",
			  "alertDisplayName": "[SAMPLE ALERT] Suspected successful brute force attack",
			  "alertName": "SIMULATED_SQL.VM_BruteForce",
			  "detectedTimeUtc": "2022-03-28T18:25:27.7187408Z",
			  "description": "THIS IS A SAMPLE ALERT: A successful login occurred after an apparent brute force attack on your resource",
			  "remediationSteps": "Go to the firewall settings in order to lock down the firewall as tightly as possible.",
			  "actionTaken": "Undefined",
			  "reportedSeverity": "High",
			  "compromisedEntity": "Sample-VM",
			  "associatedResource": "/SUBSCRIPTIONS/97e01fd4-3326-41f4-b9e3-3cfd6809e10f/RESOURCEGROUPS/Sample-RG/providers/Microsoft.Compute/virtualMachines/Sample-VM",
			  "subscriptionId": "97e01fd4-3326-41f4-b9e3-3cfd6809e10f",
			  "instanceId": "ee939333-c75e-461d-83c2-712ba9abfadb",
			  "extendedProperties": {
				"resourceType": "SQL Server 2019",
				"potential causes": "Brute force attack; penetration testing.",
				"client principal name": "Sample-account",
				"alert Id": "00000000-0000-0000-0000-000000000000",
				"client IP address": "00.00.00.00",
				"client IP location": "san antonio, united states",
				"client application": "Sample-app",
				"successful logins": "1",
				"oms workspace ID": "00000000-0000-0000-0000-000000000000",
				"failed logins": "0",
				"oms agent ID": "00000000-0000-0000-0000-000000000000",
				"enrichment_tas_threat__reports": "{\"Kind\":\"MultiLink\",\"DisplayValueToUrlDictionary\":{\"Report: Brute Force\":\"https://interflowwebportalext.trafficmanager.net/reports/DisplayReport?callerIdentity=ddd5443d-e6f4-441c-b52b-5278d2f21dfa&reportCreateDateTime=2022-03-28T17%3a59%3a38&reportName=MSTI-TS-Brute-Force.pdf&tenantId=52aab34c-2534-4485-bf2f-b4e7e5c42e44&urlCreateDateTime=2022-03-28T17%3a59%3a38&token=ubgWSjyr0wFBBiPZyz2jJ5lmrPmwXy2jLxDNy7RVz7k=\"}}",
				"killChainIntent": "PreAttack"
			  },
			  "state": "Active",
			  "reportedTimeUtc": "2022-03-28T18:26:47.490371Z",
			  "confidenceReasons": [],
			  "canBeInvestigated": true,
			  "isIncident": false,
			  "entities": [
				{
				  "$id": "centralus_1",
				  "hostName": "Sample-VM",
				  "azureID": "/SUBSCRIPTIONS/97e01fd4-3326-41f4-b9e3-3cfd6809e10f/RESOURCEGROUPS/Sample-RG/providers/Microsoft.Compute/virtualMachines/Sample-VM",
				  "omsAgentID": "00000000-0000-0000-0000-000000000000",
				  "type": "host"
				},
				{
				  "$id": "centralus_2",
				  "address": "00.00.00.00",
				  "location": {
					"countryCode": "sample",
					"countryName": "united states",
					"state": "texas",
					"city": "san antonio",
					"longitude": 0,
					"latitude": 0,
					"asn": 0,
					"carrier": "sample",
					"organization": "sample-organization",
					"organizationType": "sample-organization",
					"cloudProvider": "Azure",
					"systemService": "sample"
				  },
				  "type": "ip"
				},
				{
				  "$id": "centralus_3",
				  "sourceAddress": {
					"$ref": "centralus_2"
				  },
				  "protocol": "Tcp",
				  "type": "network-connection"
				},
				{
				  "$id": "centralus_4",
				  "name": "Sample-SA",
				  "host": {
					"$ref": "centralus_1"
				  },
				  "type": "account"
				}
			  ]
			}
		  }`,
	},
}

var OracleEventChain = []Event{
	{
		EventType: "io.triggermesh.oracle.cloudguard.alert",
		EventData: `{
			"eventType" : "com.oraclecloud.cloudguard.problemdetected",
			"cloudEventsVersion" : "0.1",
			"eventTypeVersion" : "2.0",
			"source" : "CloudGuardResponderEngine",
			"eventTime" : "2022-03-01T17:24:59Z",
			"contentType" : "application/json",
			"data" : {
			  "compartmentId" : "ocid1.compartment.oc1..1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e",
			  "compartmentName" : "Comp-Name",
			  "resourceName" : "Scanned host has vulnerabilities",
			  "resourceId" : "ocid1.cloudguardproblem.oc1.iad.1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e",
			  "additionalDetails" : {
				"tenantId" : "ocid1.tenancy.oc1..1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e",
				"status" : "OPEN",
				"reason" : "Existing Problem updated by CloudGuard",
				"problemName" : "SCANNED_HOST_VULNERABILITY",
				"riskLevel" : "CRITICAL",
				"problemType" : "CONFIG_CHANGE",
				"resourceName" : "b60b95e8-e229-4398-b3bf-25d1fe51b4f0",
				"resourceId" : "ocid1.instance.oc1.iad.1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e",
				"resourceType" : "HostAgentScan",
				"targetId" : "ocid1.cloudguardtarget.oc1.iad.1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e",
				"labels" : "VSS",
				"firstDetected" : "2022-01-30T11:14:29.130Z",
				"lastDetected" : "2022-03-01T17:24:07.129Z",
				"region" : "us-ashburn-1",
				"problemAdditionalDetails" : {
				  "Number of Low CVEs" : "2",
				  "High CVEs" : "[CVE-2018-20796, CVE-2018-1000001, CVE-2020-1751, CVE-2017-1000366, CVE-2018-12886, CVE-2018-11237, CVE-2020-1752, CVE-2016-8637, CVE-2019-9192, CVE-2015-8982, CVE-2021-3326, CVE-2016-3706, CVE-2020-7014, CVE-2020-7013, CVE-2015-8983, CVE-2020-7012, CVE-2015-7547, CVE-2020-6096, CVE-2019-19246, CVE-2018-19591, CVE-2016-6323, CVE-2021-37322, CVE-2016-3075, CVE-2019-1010023, CVE-2019-6488, CVE-2019-15847, CVE-2021-38604, CVE-2016-5417, CVE-2018-10897, CVE-2016-1234, CVE-2015-5180, CVE-2020-7009, CVE-2019-16163, CVE-2009-5155]",
				  "Scan Result Id" : "ocid1.vsshostscanresult.oc1..aaaaaaaa4bhyrwosapbnctty2xpftp6escxuxml3qppegofhjduzgre3nqiq",
				  "Critical CVEs" : "[CVE-2021-45046, CVE-2019-9169, CVE-2022-23218, CVE-2015-8778, CVE-2015-8779, CVE-2022-23219, CVE-2018-6485, CVE-2017-15670, CVE-2021-35942, CVE-2019-1010022, CVE-2014-9984, CVE-2018-11236, CVE-2014-9761, CVE-2015-8776, CVE-2017-15804, CVE-2021-44228]",
				  "Number of Critical CVEs" : "16",
				  "Number of High CVEs" : "34",
				  "Low CVEs" : "[CVE-2020-7020, CVE-2021-22136]"
				},
				"problemDescription" : "Prerequisite: Create a Host Scan Recipe and a Host Scan Target in the Scanning service. The Scanning service scans compute hosts to identify known cybersecurity vulnerabilities related to applications, libraries, operating systems, and services. This detector triggers a problem when the Scanning service has reported that an instance has one or more CRITICAL (or lower severity, based on the Input Settings within the detector config) vulnerabilities.",
				"problemRecommendation" : "Patch the reported CVE's detected on host by performing actions recommended for each CVE."
			  }
			}
		}`,
	}, {
		EventType: "io.triggermesh.oracle.cloudguard.alert",
		EventData: `{
			"eventType" : "com.oraclecloud.cloudguard.problemdetected",
			"cloudEventsVersion" : "0.1",
			"eventTypeVersion" : "2.0",
			"source" : "CloudGuardResponderEngine",
			"eventTime" : "2022-03-01T17:24:59Z",
			"contentType" : "application/json",
			"data" : {
			  "compartmentId" : "ocid1.compartment.oc1..1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e",
			  "compartmentName" : "Comp-Name",
			  "resourceName" : "Scanned host has vulnerabilities",
			  "resourceId" : "ocid1.cloudguardproblem.oc1.iad.1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e",
			  "additionalDetails" : {
				"tenantId" : "ocid1.tenancy.oc1..1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e",
				"status" : "OPEN",
				"reason" : "Existing Problem updated by CloudGuard",
				"problemName" : "SCANNED_HOST_VULNERABILITY",
				"riskLevel" : "CRITICAL",
				"problemType" : "CONFIG_CHANGE",
				"resourceName" : "aDocker",
				"resourceId" : "ocid1.instance.oc1.iad.1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e",
				"resourceType" : "HostAgentScan",
				"targetId" : "ocid1.cloudguardtarget.oc1.iad.1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e1q2w3e",
				"labels" : "VSS",
				"firstDetected" : "2022-01-30T11:14:29.130Z",
				"lastDetected" : "2022-03-01T17:24:07.129Z",
				"region" : "us-ashburn-1",
				"problemAdditionalDetails" : {
				  "Number of Low CVEs" : "2",
				  "High CVEs" : "[CVE-2018-20796, CVE-2018-1000001, CVE-2020-1751, CVE-2017-1000366, CVE-2018-12886, CVE-2018-11237, CVE-2020-1752, CVE-2016-8637, CVE-2019-9192, CVE-2015-8982, CVE-2021-3326, CVE-2016-3706, CVE-2020-7014, CVE-2020-7013, CVE-2015-8983, CVE-2020-7012, CVE-2015-7547, CVE-2020-6096, CVE-2019-19246, CVE-2018-19591, CVE-2016-6323, CVE-2021-37322, CVE-2016-3075, CVE-2019-1010023, CVE-2019-6488, CVE-2019-15847, CVE-2021-38604, CVE-2016-5417, CVE-2018-10897, CVE-2016-1234, CVE-2015-5180, CVE-2020-7009, CVE-2019-16163, CVE-2009-5155]",
				  "Scan Result Id" : "ocid1.vsshostscanresult.oc1..aaaaaaaa4bhyrwosapbnctty2xpftp6escxuxml3qppegofhjduzgre3nqiq",
				  "Critical CVEs" : "[CVE-2021-45046, CVE-2019-9169, CVE-2022-23218, CVE-2015-8778, CVE-2015-8779, CVE-2022-23219, CVE-2018-6485, CVE-2017-15670, CVE-2021-35942, CVE-2019-1010022, CVE-2014-9984, CVE-2018-11236, CVE-2014-9761, CVE-2015-8776, CVE-2017-15804, CVE-2021-44228]",
				  "Number of Critical CVEs" : "16",
				  "Number of High CVEs" : "34",
				  "Low CVEs" : "[CVE-2020-7020, CVE-2021-22136]"
				},
				"problemDescription" : "Prerequisite: Create a Host Scan Recipe and a Host Scan Target in the Scanning service. The Scanning service scans compute hosts to identify known cybersecurity vulnerabilities related to applications, libraries, operating systems, and services. This detector triggers a problem when the Scanning service has reported that an instance has one or more CRITICAL (or lower severity, based on the Input Settings within the detector config) vulnerabilities.",
				"problemRecommendation" : "Patch the reported CVE's detected on host by performing actions recommended for each CVE."
			  }
			}
		}`,
	},
}

var AquasecEventChain = []Event{
	{
		EventType: "com.microsoft.aquasec.alert",
		EventData: `{
			"id": 442,
			"time": 1627587060,
			"date": 0,
			"type": "Container.Engine",
			"user": "",
			"action": "exec",
			"image": "httpd:latest",
			"imagehash": "sha256:73b8cfec11558fe86f565b4357f6d6c8560f4c49a5f15ae970a24da86c9adc93",
			"imageid": "",
			"container": "apache",
			"containerid": "5a4da19ff2703ad2f48db4d55e563a37828dccc0f11fb6c00e60a271ab3c37cb",
			"host": "aks-default-15484652-vmss000001.tpe5bzjk4yoevknn00ux3h31kb.ax.internal.cloudapp.net",
			"hostid": "5f38a4a3-8047-4b63-adf5-5608f2a9f6eb",
			"category": "container",
			"result": 2,
			"data": "{\"host\": \"aks-default-15484652-vmss000001.tpe5bzjk4yoevknn00ux3h31kb.ax.internal.cloudapp.net\", \"rule\": \"test-block-exec\", \"time\": 1627587060, \"image\": \"httpd:latest\", \"level\": \"block\", \"vm_id\": \"909679a6-f7a8-4d1e-ab49-ebce7eaef47d\", \"action\": \"exec\", \"hostid\": \"5f38a4a3-8047-4b63-adf5-5608f2a9f6eb\", \"hostip\": \"10.240.0.5\", \"reason\": \"Unauthorized container exec\", \"result\": 2, \"tactic\": \"Execution\", \"control\": \"Block Container Exec\", \"imageid\": \"73b8cfec11558fe86f565b4357f6d6c8560f4c49a5f15ae970a24da86c9adc93\", \"podname\": \"apache\", \"podtype\": \"container\", \"vm_name\": \"aks-default-15484652-vmss_1\", \"category\": \"container\", \"resource\": \"bash\", \"vm_group\": \"MC_rnd-aks2729-aks-rg_aks2729_westeurope\", \"container\": \"apache\", \"hostgroup\": \"aquactl-default-enforcer-group\", \"rule_type\": \"runtime.policy\", \"technique\": \"Command and Script Interpreter\", \"repository\": \"httpd\", \"containerid\": \"5a4da19ff2703ad2f48db4d55e563a37828dccc0f11fb6c00e60a271ab3c37cb\", \"k8s_cluster\": \"aqua-secure\", \"vm_location\": \"westeurope\", \"podnamespace\": \"test\"}",
			"account_id": 0
		}`,
	},
}
