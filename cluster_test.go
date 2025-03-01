package proxmox

import (
	"testing"

	"github.com/luthermonson/go-proxmox/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCluster(t *testing.T) {
	mocks.On(mockConfig)
	defer mocks.Off()
	client := mockClient()

	cluster, err := client.Cluster()
	assert.Nil(t, err)
	assert.Equal(t, 4, cluster.Version)
	assert.Equal(t, "cluster", cluster.ID)
	for _, n := range cluster.Nodes {
		assert.Contains(t, n.ID, "node/node")
		assert.Equal(t, "node", n.Type)
	}
}

func TestNextID(t *testing.T) {
	mocks.On(mockConfig)
	defer mocks.Off()
	client := mockClient()

	cluster, err := client.Cluster()
	assert.Nil(t, err)
	nextid, err := cluster.NextID()
	assert.Nil(t, err)
	assert.Equal(t, 100, nextid)
}

func TestCluster_Resources(t *testing.T) {
	mocks.On(mockConfig)
	defer mocks.Off()
	client := mockClient()

	cluster, err := client.Cluster()
	assert.Nil(t, err)

	// json unmarshall tests
	rs, err := cluster.Resources()
	assert.Equal(t, 20, len(rs))

	// type param test
	rs, err = cluster.Resources("node")
	assert.Equal(t, 1, len(rs))
}
