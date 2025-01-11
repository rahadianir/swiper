package swiper

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"
	"testing"

	"github.com/rahadianir/swiper/internal/common"
	"github.com/rahadianir/swiper/internal/models"
	mock_swiper "github.com/rahadianir/swiper/internal/swiper/mock"
	"go.uber.org/mock/gomock"
)

var (
	expectedUser = models.User{
		ID:         1,
		Name:       "fulanah",
		Username:   "fulanahxyz",
		Age:        29,
		Gender:     "female",
		Location:   "bekasi",
		IsPremium:  false,
		IsVerified: false,
	}
	exampleCacheData = models.ActivityCache{
		UserID: 23,
		Pass:   []int{5, 6, 7},
		Likes:  []int{11, 12, 13},
	}
	exampleUser = models.User{
		ID:         23,
		Name:       "ian",
		Username:   "ianabc",
		Age:        27,
		Gender:     "male",
		Location:   "jakarta",
		IsPremium:  false,
		IsVerified: false,
	}
)

func TestSwiperLogic_GetTargetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_swiper.NewMockUserRepositoryInterface(ctrl)
	mockCacheStore := mock_swiper.NewMockCacheInterface(ctrl)
	mockSwiperRepo := mock_swiper.NewMockSwiperRepositoryInterface(ctrl)

	deps := common.Dependencies{}

	type fields struct {
		Dependencies *common.Dependencies
		UserRepo     UserRepositoryInterface
		CacheStore   CacheInterface
		SwiperRepo   SwiperRepositoryInterface
	}
	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
		setup   func()
	}{
		{
			name: "success get target profile",
			fields: fields{
				Dependencies: &deps,
				UserRepo:     mockUserRepo,
				CacheStore:   mockCacheStore,
				SwiperRepo:   mockSwiperRepo,
			},
			args: args{
				ctx:    context.Background(),
				userID: 23,
			},
			want:    expectedUser,
			wantErr: false,
			setup: func() {
				cacheBytes, err := json.Marshal(exampleCacheData)
				if err != nil {
					t.Fatalf("%v", err)
				}
				mockCacheStore.EXPECT().Get(gomock.Any(), strconv.Itoa(23)).Return(string(cacheBytes), nil)

				mockSwiperRepo.EXPECT().GetUserLikedUserIDs(gomock.Any(), 23, models.LikedUserParams{
					IsMatched: true,
				}).Return([]int{8, 9, 10}, nil)

				excludedlist := []int{23, 5, 6, 7, 11, 12, 13, 8, 9, 10}

				mockUserRepo.EXPECT().GetRandomUser(gomock.Any(), excludedlist).Return(expectedUser, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			logic := &SwiperLogic{
				Dependencies: tt.fields.Dependencies,
				UserRepo:     tt.fields.UserRepo,
				CacheStore:   tt.fields.CacheStore,
				SwiperRepo:   tt.fields.SwiperRepo,
			}
			got, err := logic.GetTargetProfile(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SwiperLogic.GetTargetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwiperLogic.GetTargetProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwiperLogic_SwipeRight(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_swiper.NewMockUserRepositoryInterface(ctrl)
	mockCacheStore := mock_swiper.NewMockCacheInterface(ctrl)
	mockSwiperRepo := mock_swiper.NewMockSwiperRepositoryInterface(ctrl)

	deps := common.Dependencies{}

	type fields struct {
		Dependencies *common.Dependencies
		UserRepo     UserRepositoryInterface
		CacheStore   CacheInterface
		SwiperRepo   SwiperRepositoryInterface
	}
	type args struct {
		ctx      context.Context
		userID   int
		targetId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
		setup   func()
	}{
		{
			name: "success swipe right and matched!",
			fields: fields{
				Dependencies: &deps,
				UserRepo:     mockUserRepo,
				CacheStore:   mockCacheStore,
				SwiperRepo:   mockSwiperRepo,
			},
			args: args{
				ctx:      context.Background(),
				userID:   23,
				targetId: 1,
			},
			want:    true,
			wantErr: false,
			setup: func() {
				cacheBytes, err := json.Marshal(exampleCacheData)
				if err != nil {
					t.Fatalf("%v", err)
				}
				mockCacheStore.EXPECT().Get(gomock.Any(), strconv.Itoa(23)).Return(string(cacheBytes), nil)

				mockUserRepo.EXPECT().GetUserByUserID(gomock.Any(), 23).Return(exampleUser, nil)
				mockSwiperRepo.EXPECT().StoreUserLike(gomock.Any(), 23, 1).Return(nil)

				newExampleCacheData := exampleCacheData
				newExampleCacheData.Likes = append(newExampleCacheData.Likes, 1)
				mockCacheStore.EXPECT().Update(gomock.Any(), strconv.Itoa(23), newExampleCacheData).Return(nil)

				mockSwiperRepo.EXPECT().GetUserLikedUserIDs(gomock.Any(), 1, models.LikedUserParams{}).Return([]int{23}, nil)
				mockSwiperRepo.EXPECT().UpdateMatchStatus(gomock.Any(), 23, 1, true).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			logic := &SwiperLogic{
				Dependencies: tt.fields.Dependencies,
				UserRepo:     tt.fields.UserRepo,
				CacheStore:   tt.fields.CacheStore,
				SwiperRepo:   tt.fields.SwiperRepo,
			}
			got, err := logic.SwipeRight(tt.args.ctx, tt.args.userID, tt.args.targetId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SwiperLogic.SwipeRight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SwiperLogic.SwipeRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwiperLogic_SwipeLeft(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mock_swiper.NewMockUserRepositoryInterface(ctrl)
	mockCacheStore := mock_swiper.NewMockCacheInterface(ctrl)
	mockSwiperRepo := mock_swiper.NewMockSwiperRepositoryInterface(ctrl)

	deps := common.Dependencies{}

	type fields struct {
		Dependencies *common.Dependencies
		UserRepo     UserRepositoryInterface
		CacheStore   CacheInterface
		SwiperRepo   SwiperRepositoryInterface
	}
	type args struct {
		ctx      context.Context
		userID   int
		targetId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		setup   func()
	}{
		{
			name: "success swipe left",
			fields: fields{
				Dependencies: &deps,
				UserRepo:     mockUserRepo,
				CacheStore:   mockCacheStore,
				SwiperRepo:   mockSwiperRepo,
			},
			args: args{
				ctx:      context.Background(),
				userID:   23,
				targetId: 1,
			},
			wantErr: false,
			setup: func() {
				cacheBytes, err := json.Marshal(exampleCacheData)
				if err != nil {
					t.Fatalf("%v", err)
				}
				mockCacheStore.EXPECT().Get(gomock.Any(), strconv.Itoa(23)).Return(string(cacheBytes), nil)

				mockUserRepo.EXPECT().GetUserByUserID(gomock.Any(), 23).Return(exampleUser, nil)

				newExampleCacheData := exampleCacheData
				newExampleCacheData.Pass = append(newExampleCacheData.Pass, 1)
				mockCacheStore.EXPECT().Update(gomock.Any(), strconv.Itoa(23), newExampleCacheData).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			logic := &SwiperLogic{
				Dependencies: tt.fields.Dependencies,
				UserRepo:     tt.fields.UserRepo,
				CacheStore:   tt.fields.CacheStore,
				SwiperRepo:   tt.fields.SwiperRepo,
			}
			if err := logic.SwipeLeft(tt.args.ctx, tt.args.userID, tt.args.targetId); (err != nil) != tt.wantErr {
				t.Errorf("SwiperLogic.SwipeLeft() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
