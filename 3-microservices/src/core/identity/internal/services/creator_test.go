package services

import (
	"errors"
	"identity/internal/databases"
	"identity/internal/httpbody/request"
	mocks "identity/mocks/databases/creator"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	tests := []struct {
		name                 string
		creatorId            string
		requestUpdateCreator request.UpdateCreator
		givenError           error
		expectError          error
		expectedResponse     interface{}
	}{
		{
			name:      "update creator success",
			creatorId: "1",
			requestUpdateCreator: request.UpdateCreator{
				Type: "test",
			},
		},
		{
			name:      "not found creator",
			creatorId: "2",
			requestUpdateCreator: request.UpdateCreator{
				Type: "test",
			},
			givenError:  errors.New("not found record"),
			expectError: errors.New("not found record"),
		},
		{
			name:      "internal error",
			creatorId: "3",
			requestUpdateCreator: request.UpdateCreator{
				Type: "test",
			},
			givenError:  errors.New("timeout"),
			expectError: errors.New("timeout"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// given
			mockCreator := new(mocks.Repository)
			mockCreator.On(
				"Update",
				mock.Anything, // context can be mocked
				mock.Anything,
			).Return(
				tc.givenError,
			)

			dbStore := databases.DBStore{
				Creator: mockCreator,
			}

			// when
			c := &Creator{dbStore}
			err := c.Update(tc.creatorId, tc.requestUpdateCreator)

			// then
			require.Equal(t, tc.expectError, err)
		})
	}
}
