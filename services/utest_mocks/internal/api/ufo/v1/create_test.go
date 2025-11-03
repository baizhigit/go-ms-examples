package v1

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/baizhigit/go-ms-examples/utest_mocks/internal/converter"
	ufoV1 "github.com/baizhigit/go-ms-examples/utest_mocks/pkg/proto/ufo/v1"
)

func (s *APISuite) TestCreateSuccess() {
	var (
		location        = gofakeit.City()
		description     = gofakeit.Paragraph(3, 5, 5, " ")
		observedAt      = time.Now()
		color           = gofakeit.Color()
		sound           = gofakeit.Bool()
		durationSeconds = int32(60)

		expectedUUID = gofakeit.UUID()

		protoInfo = &ufoV1.SightingInfo{
			ObservedAt:      timestamppb.New(observedAt),
			Location:        location,
			Description:     description,
			Color:           wrapperspb.String(color),
			Sound:           wrapperspb.Bool(sound),
			DurationSeconds: wrapperspb.Int32(durationSeconds),
		}

		req = &ufoV1.CreateRequest{
			Info: protoInfo,
		}

		expectedModelInfo = converter.UFOInfoToModel(protoInfo)
	)

	s.ufoService.On("Create", s.ctx, expectedModelInfo).Return(expectedUUID, nil)

	res, err := s.api.Create(s.ctx, req)
	s.Require().NoError(err)
	s.Require().NotNil(res)
	s.Require().Equal(expectedUUID, res.GetUuid())
}

func (s *APISuite) TestCreateServiceError() {
	var (
		serviceErr      = gofakeit.Error()
		location        = gofakeit.City()
		description     = gofakeit.Paragraph(3, 5, 5, " ")
		observedAt      = time.Now()
		color           = gofakeit.Color()
		sound           = true
		durationSeconds = int32(60)

		protoInfo = &ufoV1.SightingInfo{
			ObservedAt:      timestamppb.New(observedAt),
			Location:        location,
			Description:     description,
			Color:           wrapperspb.String(color),
			Sound:           wrapperspb.Bool(sound),
			DurationSeconds: wrapperspb.Int32(durationSeconds),
		}

		req = &ufoV1.CreateRequest{
			Info: protoInfo,
		}

		expectedModelInfo = converter.UFOInfoToModel(protoInfo)
	)

	s.ufoService.On("Create", s.ctx, expectedModelInfo).Return("", serviceErr)

	res, err := s.api.Create(s.ctx, req)
	s.Require().Error(err)
	s.Require().ErrorIs(err, serviceErr)
	s.Require().Nil(res)
}
