package ufo

import (
	"context"
	"errors"
	"time"

	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/baizhigit/go-ms-examples/di/ufo/internal/model"
)

func (r *repository) Delete(ctx context.Context, uuid string) error {
	// Проверяем существование документа
	var existing bson.M
	err := r.collection.FindOne(ctx, bson.M{"_id": uuid}).Decode(&existing)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.ErrSightingNotFound
		}
		return err
	}

	// Мягкое удаление - устанавливаем deleted_at
	updateDoc := bson.M{
		"$set": bson.M{
			"deleted_at": lo.ToPtr(time.Now()),
		},
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": uuid}, updateDoc)
	if err != nil {
		return err
	}

	return nil
}
