package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository[Model any, AggregateResult any] struct {
	collection *mongo.Collection
}

func NewRepository[Model any, AggregateResult any](collection *mongo.Collection) *Repository[Model, AggregateResult] {
	return &Repository[Model, AggregateResult]{collection: collection}
}

func (r *Repository[Model, AggregateResult]) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*Model, error) {
	var result Model
	err := r.collection.FindOne(ctx, filter, opts...).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *Repository[Model, AggregateResult]) InsertOne(ctx context.Context, document *Model, opts ...*options.InsertOneOptions) (interface{}, error) {
	insertOneResult, err := r.collection.InsertOne(ctx, document, opts...)
	if err != nil {
		return nil, err
	}

	return insertOneResult.InsertedID, nil
}

func (r *Repository[Model, AggregateResult]) InsertMany(ctx context.Context, document *[]Model, opts ...*options.InsertManyOptions) ([]interface{}, error) {
	docsInterface := make([]interface{}, len(*document))
	for i := range *document {
		docsInterface[i] = (*document)[i]
	}

	insertManyResult, err := r.collection.InsertMany(ctx, docsInterface, opts...)
	if err != nil {
		return nil, err
	}

	return insertManyResult.InsertedIDs, nil
}

func (r *Repository[Model, AggregateResult]) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) error {
	_, err := r.collection.DeleteOne(ctx, filter, opts...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository[Model, AggregateResult]) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) error {
	_, err := r.collection.DeleteMany(ctx, filter, opts...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository[Model, AggregateResult]) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) error {
	_, err := r.collection.UpdateOne(ctx, filter, update, opts...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository[Model, AggregateResult]) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	cursor, err := r.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

func (r *Repository[Model, AggregateResult]) FindAll(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]Model, error) {
	cursor, err := r.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []Model
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository[Model, AggregateResult]) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	cursor, err := r.collection.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	return cursor, err
}

func (r *Repository[Model, AggregateResult]) AggregateAll(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) ([]AggregateResult, error) {
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []AggregateResult
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository[Model, AggregateResult]) CountDocuments(ctx context.Context, filter *Model, opts ...*options.CountOptions) (int64, error) {
	count, err := r.collection.CountDocuments(ctx, filter, opts...)
	if err != nil {
		return 0, err
	}

	return count, err
}

func (r *Repository[Model, AggregateResult]) Drop(ctx context.Context) error {
	err := r.collection.Drop(ctx)
	if err != nil {
		return err
	}

	return nil
}
