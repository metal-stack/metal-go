package metalgo

import "testing"

func TestTagIsMemberOfCluster(t *testing.T) {
	tests := []struct {
		name      string
		tag       string
		clusterID string
		want      bool
	}{
		{
			name:      "is not a cluster tag",
			tag:       "a.tag",
			clusterID: "123",
			want:      false,
		},
		{
			name:      "is a cluster tag but no clusterid specified",
			tag:       "cluster.metal-stack.io/clusterid",
			clusterID: "123",
			want:      false,
		},
		{
			name:      "is a cluster tag but clusterid specified at wrong place",
			tag:       "cluster.metal-stack.io/clusterid/namespace/servicename=kube-system/metallb",
			clusterID: "123",
			want:      false,
		},
		{
			name:      "is a cluster tag but clusterid specified at right place",
			tag:       "cluster.metal-stack.io/clusterid/namespace/servicename=123/kube-system/metallb",
			clusterID: "123",
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TagIsMemberOfCluster(tt.tag, tt.clusterID); got != tt.want {
				t.Errorf("TagIsMemberOfCluster() = %v, want %v", got, tt.want)
			}
		})
	}
}
