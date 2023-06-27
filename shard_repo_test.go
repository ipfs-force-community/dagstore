package dagstore

import (
	"context"
	"testing"

	"github.com/ipfs/go-datastore"
	ds "github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
	"github.com/stretchr/testify/assert"
)

func TestShardRepo(t *testing.T) {
	var shards = []*PersistedShard{
		{Key: "Key", URL: "URL", TransientPath: "TransientPath", State: ShardStateInitializing, Lazy: true, Error: ""},
		{Key: "Key2", URL: "URL2", TransientPath: "TransientPath2", State: ShardStateAvailable, Lazy: false, Error: "Error"},
	}
	ctx := context.Background()
	store := datastore.NewLogDatastore(dssync.MutexWrap(datastore.NewMapDatastore()), "trace")
	repo := NewShardRepo(store)

	_, ok := repo.(ds.Datastore)
	assert.True(t, ok)

	// test save shard
	for _, shard := range shards {
		assert.NoError(t, repo.SaveShard(ctx, shard))
	}

	// test get shard
	for _, shard := range shards {
		res, err := repo.GetShard(ctx, shard.Key)
		assert.NoError(t, err)
		assert.Equal(t, shard, res)
	}
	res, err := repo.GetShard(ctx, "")
	assert.Error(t, err)
	assert.Nil(t, res)

	// test list shard
	list, err := repo.ListShards(ctx)
	assert.NoError(t, err)
	assert.Contains(t, list, shards[0])
	assert.Contains(t, list, shards[1])

	// test has shard
	for _, shard := range shards {
		has, err := repo.HasShard(ctx, shard.Key)
		assert.NoError(t, err)
		assert.True(t, has)
	}
	has, err := repo.HasShard(ctx, "")
	assert.NoError(t, err)
	assert.False(t, has)

	// test delete shard
	for _, shard := range shards {
		assert.NoError(t, repo.DeleteShard(ctx, shard.Key))

		has, err := repo.HasShard(ctx, shard.Key)
		assert.NoError(t, err)
		assert.False(t, has)
	}
}
