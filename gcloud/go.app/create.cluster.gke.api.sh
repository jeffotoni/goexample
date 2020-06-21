POST https://container.googleapis.com/v1beta1/projects/projeto-eng1/zones/southamerica-east1-c/clusters
{
  "cluster": {
    "name": "go-app",
    "masterAuth": {
      "clientCertificateConfig": {}
    },
    "network": "projects/projeto-eng1/global/networks/default",
    "addonsConfig": {
      "httpLoadBalancing": {},
      "horizontalPodAutoscaling": {},
      "kubernetesDashboard": {
        "disabled": true
      },
      "istioConfig": {
        "disabled": true
      },
      "dnsCacheConfig": {}
    },
    "subnetwork": "projects/projeto-eng1/regions/southamerica-east1/subnetworks/default",
    "nodePools": [
      {
        "name": "default-pool",
        "config": {
          "machineType": "custom-1-1024",
          "diskSizeGb": 10,
          "oauthScopes": [
            "https://www.googleapis.com/auth/devstorage.read_only",
            "https://www.googleapis.com/auth/logging.write",
            "https://www.googleapis.com/auth/monitoring",
            "https://www.googleapis.com/auth/servicecontrol",
            "https://www.googleapis.com/auth/service.management.readonly",
            "https://www.googleapis.com/auth/trace.append"
          ],
          "metadata": {
            "disable-legacy-endpoints": "true"
          },
          "imageType": "COS",
          "diskType": "pd-standard",
          "shieldedInstanceConfig": {
            "enableIntegrityMonitoring": true
          }
        },
        "initialNodeCount": 3,
        "autoscaling": {
          "enabled": true,
          "maxNodeCount": 3
        },
        "management": {
          "autoUpgrade": true,
          "autoRepair": true
        },
        "upgradeSettings": {
          "maxSurge": 1
        }
      }
    ],
    "networkPolicy": {},
    "ipAllocationPolicy": {
      "useIpAliases": true
    },
    "masterAuthorizedNetworksConfig": {},
    "defaultMaxPodsConstraint": {
      "maxPodsPerNode": "110"
    },
    "authenticatorGroupsConfig": {},
    "privateClusterConfig": {},
    "databaseEncryption": {
      "state": "DECRYPTED"
    },
    "shieldedNodes": {},
    "releaseChannel": {
      "channel": "RAPID"
    },
    "clusterTelemetry": {
      "type": "ENABLED"
    },
    "location": "southamerica-east1-c"
  }
}