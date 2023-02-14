package redsumer

import (
	"context"

	"github.com/enerBit/redsumer/pkg/client"
	"github.com/enerBit/redsumer/pkg/consumer"
	"github.com/redis/go-redis/v9"
)

type redConsumer struct {
	args   RedConsumerArgs
	client *redis.Client
}

type RedConsumerArgs struct {
	Group        string
	Stream       string
	ConsumerName string
	RedisHost    string
	RedisPort    int
	Db           int
}

func NewRedisConsumer(args RedConsumerArgs) (redConsumer, error) {

	client, err := client.NewRedisClient(args.RedisHost, args.RedisPort, args.Db)

	if err != nil {
		return redConsumer{}, err
	}

	return redConsumer{
		args:   args,
		client: client,
	}, nil
}

func (c redConsumer) Consume(ctx context.Context) ([]redis.XMessage, error) {

	messages, err := consumer.Consume(ctx, c.client, c.args.Group, c.args.ConsumerName, c.args.Stream)
	return messages, err
}

func (c redConsumer) WaitForStream(ctx context.Context, tries []int) error {

	err := consumer.WaitForStream(ctx, c.client, c.args.Stream, tries)
	return err
}

// Get the raw redis client from the librari redis-go
func (c redConsumer) RawRedis() *redis.Client {
	return c.client
}

func (c redConsumer) GetArgs() RedConsumerArgs {
	return c.args
}
