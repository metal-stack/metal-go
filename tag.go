package metalgo

import (
	"fmt"
)

// TagServicePrefix the prefix of the tag used to identify services
const TagServicePrefix = "cluster.metal-pod.io/clusterid/namespace/servicename"

// BuildServiceTag constructs the service tag for the given cluster and service
func BuildServiceTag(clusterID string, namespace, serviceName string) string {
	return fmt.Sprintf("%s/%s/%s", BuildServiceTagClusterPrefix(clusterID), namespace, serviceName)
}

// BuildServiceTagClusterPrefix constructs the prefix of the service tag that identify all services of a cluster
func BuildServiceTagClusterPrefix(clusterID string) string {
	return fmt.Sprintf("%s=%s", TagServicePrefix, clusterID)
}
