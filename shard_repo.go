package dagstore

import (
	"context"
	"encoding/json"

	ds "github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/namespace"
	"github.com/ipfs/go-datastore/query"
)

type ShardRepo interface {
	SaveShard(ctx context.Context, shard *PersistedShard) error
	GetShard(ctx context.Context, key string) (*PersistedShard, error)
	ListShards(ctx context.Context) ([]*PersistedShard, error)
	HasShard(ctx context.Context, key string) (bool, error)
	DeleteShard(ctx context.Context, key string) error
}

type ShardRepoImp struct {
	ds.Datastore
}

var _ ShardRepo = &ShardRepoImp{}
var _ ds.Datastore = &ShardRepoImp{}

func NewShardRepo(store ds.Datastore) ShardRepo {
	// namespace all store operations.
	store = namespace.Wrap(store, StoreNamespace)
	return &ShardRepoImp{Datastore: store}
}

func (s *ShardRepoImp) SaveShard(ctx context.Context, shard *PersistedShard) error {
	key := ds.NewKey(shard.Key)
	data, err := json.Marshal(shard)
	if err != nil {
		return err
	}

	return s.Datastore.Put(ctx, key, data)
}

func (s *ShardRepoImp) GetShard(ctx context.Context, key string) (*PersistedShard, error) {
	dsKey := ds.NewKey(key)
	data, err := s.Datastore.Get(ctx, dsKey)
	if err != nil {
		return nil, err
	}
	var shard PersistedShard
	err = json.Unmarshal(data, &shard)
	if err != nil {
		return nil, err
	}

	return &shard, nil
}

func (s *ShardRepoImp) ListShards(ctx context.Context) ([]*PersistedShard, error) {
	res, err := s.Datastore.Query(ctx, query.Query{})
	if err != nil {
		return nil, err
	}

	var shards []*PersistedShard
	for {
		result, ok := res.NextSync()
		if !ok {
			break
		}

		var shard PersistedShard
		if err := json.Unmarshal(result.Value, &shard); err != nil {
			return nil, err
		}

		shards = append(shards, &shard)
	}

	return shards, nil
}

func (s *ShardRepoImp) HasShard(ctx context.Context, key string) (bool, error) {
	dsKey := ds.NewKey(key)

	return s.Datastore.Has(ctx, dsKey)
}

func (s *ShardRepoImp) DeleteShard(ctx context.Context, key string) error {
	dsKey := ds.NewKey(key)

	return s.Datastore.Delete(ctx, dsKey)
}
