package services

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mocksDatastore "github.com/gorcz/ports/mocks/pkg/datastore"
	mocksIterator "github.com/gorcz/ports/mocks/pkg/model"
	"github.com/gorcz/ports/pkg/model/port"
	"github.com/stretchr/testify/require"
)

func TestService_UpsertPorts(suite *testing.T) {
	suite.Parallel()

	ctx := context.Background()
	ctrl := gomock.NewController(suite)
	defer ctrl.Finish()

	// TODO: Add other test cases. There is only happy path covered

	suite.Run("should upsert ports to datastore", func(t *testing.T) {
		firstTestPort := &port.Port{
			Code:    "FIRST",
			Details: port.Details{},
		}
		secondTestPort := &port.Port{
			Code:    "SECOND",
			Details: port.Details{},
		}

		mockDatastore := mocksDatastore.NewMockDatastore(ctrl)
		mockDatastore.EXPECT().UpsertPort(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, port *port.Port) error {
				require.Equal(t, port, firstTestPort)
				return nil
			})
		mockDatastore.EXPECT().UpsertPort(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, port *port.Port) error {
				require.Equal(t, port, secondTestPort)
				return nil
			})

		mockPortIterator := mocksIterator.NewMockIterator(ctrl)
		mockPortIterator.EXPECT().Next().DoAndReturn(func() (*port.Port, bool, error) {
			return firstTestPort, true, nil
		})
		mockPortIterator.EXPECT().Next().DoAndReturn(func() (*port.Port, bool, error) {
			return secondTestPort, true, nil
		})
		mockPortIterator.EXPECT().Next().DoAndReturn(func() (*port.Port, bool, error) {
			return nil, false, nil
		})

		testService := NewPorts(mockDatastore)
		err := testService.UpsertPorts(ctx, mockPortIterator)
		require.NoError(t, err)
	})
}
