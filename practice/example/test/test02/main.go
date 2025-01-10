package main

import "fmt"

type ManagedCluster struct {
	ClusterID string
	AgentPool *[]AgentPoolProperties
}

type AgentPoolProperties struct {
	ID     string
	Taints string
}

func main() {
	taints := []string{"--abc", "--cde", "--klj"}
	AgentPools := []AgentPoolProperties{
		{
			ID: "11",
		},
		{
			ID: "22",
		},
		{
			ID: "33",
		},
	}
	mc := &ManagedCluster{
		ClusterID: "1",
		AgentPool: &AgentPools,
	}
	i := 0
	for index := range *mc.AgentPool {
		ap := &(*mc.AgentPool)[index] // Get a pointer to the element
		if ap.ID != "" && ap.Taints == "" {
			ap.Taints = taints[i]
			i++
		}
	}

	fmt.Println(*mc.AgentPool)
}
