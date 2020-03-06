package repository_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/interfaces/repository"
	"github.com/ezio1119/fishapp-chat/testutil"
	uRepo "github.com/ezio1119/fishapp-chat/usecase/repository"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (uRepo.RoomRepository, sqlmock.Sqlmock, func()) {
	db, mock, _ := testutil.NewGormMock(t)
	r := repository.NewRoomRepository(db)
	return r, mock, func() {
		db.Close()
	}
}

// func TestAddMember(t *testing.T) {
// 	r, mock, dbClose := setup(t)
// 	defer dbClose()

// 	now := time.Now()
// 	tests := []struct {
// 		name       string
// 		in         *domain.Member
// 		expected   *domain.Member
// 		setSQLMock func(t *testing.T)
// 		wantErr    bool
// 		err        error
// 	}{
// 		{
// 			name: "正常",
// 			in: &domain.Member{
// 				RoomID: 1,
// 				UserID: 1,
// 			},
// 			expected: &domain.Member{
// 				ID:        1,
// 				RoomID:    1,
// 				UserID:    1,
// 				CreatedAt: now,
// 				UpdatedAt: now,
// 			},
// 			setSQLMock: func(t *testing.T) {
// 				m := &domain.Member{
// 					ID:        1,
// 					RoomID:    1,
// 					UserID:    1,
// 					CreatedAt: now,
// 					UpdatedAt: now,
// 				}
// 				query := regexp.QuoteMeta("INSERT INTO `members` (`room_id`, `user_id`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?)")
// 				rows := sqlmock.NewRows([]string{"id", "room_id", "user_id", "created_at", "updated_at"}).
// 				AddRow(m.ID, m.RoomID, m.UserID, m.CreatedAt, m.UpdatedAt)
// 				mock.ExpectQuery("").WithArgs(m.RoomID, m.UserID, m.CreatedAt, m.UpdatedAt).WillReturnRows(rows)
// 			},
// 			wantErr: false,
// 			err:     nil,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.setSQLMock(t)
// 			err := r.AddMember(context.TODO(), tt.in)
// 			assert.NoError(t, err)
// 			// assert.EqualValues(t, tt.expected, tt.in)
// 		})
// 	}
// }

func TestGetRoomByID(t *testing.T) {
	r, mock, dbClose := setup(t)
	defer dbClose()

	now := time.Now()
	tests := []struct {
		name       string
		in         int64
		expected   *domain.Room
		setSQLMock func(t *testing.T)
		wantErr    bool
		err        error
	}{
		{
			name: "正常",
			in:   1,
			expected: &domain.Room{
				ID:        1,
				PostID:    1,
				CreatedAt: now,
				UpdatedAt: now,
			},
			setSQLMock: func(t *testing.T) {
				r := &domain.Room{
					ID:        1,
					PostID:    1,
					CreatedAt: now,
					UpdatedAt: now,
				}
				query := regexp.QuoteMeta("SELECT * FROM `rooms` WHERE (id = ?) LIMIT 1")
				rows := sqlmock.NewRows([]string{"id", "post_id", "created_at", "updated_at"}).
					AddRow(r.ID, r.PostID, r.CreatedAt, r.UpdatedAt)
				mock.ExpectQuery(query).WithArgs(int64(1)).WillReturnRows(rows)
			},
			wantErr: false,
			err:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setSQLMock(t)
			r, err := r.GetRoomByID(context.TODO(), tt.in)
			assert.NoError(t, err)
			assert.EqualValues(t, tt.expected, r)
		})
	}
}
