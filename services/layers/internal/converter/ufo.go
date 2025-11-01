package converter

import (
	"time"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/baizhigit/go-ms-examples/layers/internal/model"
	ufoV1 "github.com/baizhigit/go-ms-examples/layers/pkg/proto/ufo/v1"
)

func UFOInfoToModel(info *ufoV1.SightingInfo) model.SightingInfo {
	var observedAt *time.Time
	if info.ObservedAt != nil {
		observedAt = lo.ToPtr(info.ObservedAt.AsTime())
	}

	var color *string
	if info.Color != nil {
		color = lo.ToPtr(info.Color.Value)
	}

	var sound *bool
	if info.Sound != nil {
		tmp := info.Sound.Value
		sound = &tmp
	}

	var durationSeconds *int32
	if info.DurationSeconds != nil {
		tmp := info.DurationSeconds.Value
		durationSeconds = &tmp
	}

	return model.SightingInfo{
		ObservedAt:      observedAt,
		Location:        info.Location,
		Description:     info.Description,
		Color:           color,
		Sound:           sound,
		DurationSeconds: durationSeconds,
	}
}

func UpdateInfoToModel(info *ufoV1.SightingUpdateInfo) model.SightingUpdateInfo {
	var observedAt *time.Time
	if info.ObservedAt != nil {
		tmp := info.ObservedAt.AsTime()
		observedAt = &tmp
	}

	var location *string
	if info.Location != nil {
		tmp := info.Location.Value
		location = &tmp
	}

	var description *string
	if info.Description != nil {
		tmp := info.Description.Value
		description = &tmp
	}

	var color *string
	if info.Color != nil {
		tmp := info.Color.Value
		color = &tmp
	}

	var sound *bool
	if info.Sound != nil {
		tmp := info.Sound.Value
		sound = &tmp
	}

	var durationSeconds *int32
	if info.DurationSeconds != nil {
		tmp := info.DurationSeconds.Value
		durationSeconds = &tmp
	}

	return model.SightingUpdateInfo{
		ObservedAt:      observedAt,
		Location:        location,
		Description:     description,
		Color:           color,
		Sound:           sound,
		DurationSeconds: durationSeconds,
	}
}

func SightingToProto(sighting model.Sighting) *ufoV1.Sighting {
	var updatedAt *timestamppb.Timestamp
	if sighting.UpdatedAt != nil {
		updatedAt = timestamppb.New(*sighting.UpdatedAt)
	}

	var deletedAt *timestamppb.Timestamp
	if sighting.DeletedAt != nil {
		deletedAt = timestamppb.New(*sighting.DeletedAt)
	}

	return &ufoV1.Sighting{
		Uuid:      sighting.Uuid,
		Info:      SightingInfoToProto(sighting.Info),
		CreatedAt: timestamppb.New(sighting.CreatedAt),
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

func SightingInfoToProto(info model.SightingInfo) *ufoV1.SightingInfo {
	var observedAt *timestamppb.Timestamp
	if info.ObservedAt != nil {
		observedAt = timestamppb.New(*info.ObservedAt)
	}

	var color *wrapperspb.StringValue
	if info.Color != nil {
		color = wrapperspb.String(*info.Color)
	}

	var sound *wrapperspb.BoolValue
	if info.Sound != nil {
		sound = wrapperspb.Bool(*info.Sound)
	}

	var durationSeconds *wrapperspb.Int32Value
	if info.DurationSeconds != nil {
		durationSeconds = wrapperspb.Int32(*info.DurationSeconds)
	}

	return &ufoV1.SightingInfo{
		ObservedAt:      observedAt,
		Location:        info.Location,
		Description:     info.Description,
		Color:           color,
		Sound:           sound,
		DurationSeconds: durationSeconds,
	}
}
