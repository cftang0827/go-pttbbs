package ptt

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

func TestInitCurrentUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	userid1 := ptttype.UserID_t{}
	copy(userid1[:], []byte("SYSOP"))

	type args struct {
		userID *ptttype.UserID_t
	}
	tests := []struct {
		name      string
		args      args
		expected  ptttype.Uid
		expected1 *ptttype.UserecRaw
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			args:      args{&userid1},
			expected:  1,
			expected1: testUserecRaw1,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			got, got1, err := InitCurrentUser(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitCurrentUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("InitCurrentUser() got = %v, expected %v", got, tt.expected)
			}
			if !reflect.DeepEqual(got1, tt.expected1) {
				t.Errorf("InitCurrentUser() got1 = %v, expected1 %v", got1, tt.expected1)
			}
		})
		wg.Wait()
	}
}

func Test_passwdSyncUpdate(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		uid  ptttype.Uid
		user *ptttype.UserecRaw
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{uid: ptttype.Uid(1), user: testUserecRaw1},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := passwdSyncUpdate(tt.args.uid, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("passwdSyncUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		wg.Wait()
	}
}

func Test_passwdSyncQuery(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		uid ptttype.Uid
	}
	tests := []struct {
		name     string
		args     args
		expected *ptttype.UserecRaw
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{1},
			expected: testUserecRaw1,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			got, err := passwdSyncQuery(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("passwdSyncQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("passwdSyncQuery() = %v, want %v", got, tt.expected)
			}
		})
		wg.Wait()
	}
}

func TestInitCurrentUserByUid(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		uid ptttype.Uid
	}
	tests := []struct {
		name         string
		args         args
		expectedUser *ptttype.UserecRaw
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			args:         args{1},
			expectedUser: testUserecRaw1,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotUser, err := InitCurrentUserByUid(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitCurrentUserByUid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.expectedUser) {
				t.Errorf("InitCurrentUserByUid() = %v, want %v", gotUser, tt.expectedUser)
			}
		})
		wg.Wait()
	}
}
