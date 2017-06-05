package fbs_test

import (
	"encoding/json"
	"testing"

	"github.com/bxcodec/Go-Simple-Flatbuffer/users"

	httpdlv "github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http"
	"github.com/bxcodec/Go-Simple-Flatbuffer/users/delivery/http/fbs"
	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/stretchr/testify/assert"
)

const BenchCall = 100000

func BenchmarkCreateUserJSON(b *testing.B) {
	b.N = BenchCall
	for i := 0; i < b.N; i++ {
		var user = &users.UserObj{
			ID:   64,
			Name: "Iman",
		}

		byt, _ := json.Marshal(user)
		// jsonUser := string(byt)
		jsonUser := byt
		var userD users.UserObj
		// err := json.Unmarshal([]byte(jsonUser), &userD)
		err := json.Unmarshal(jsonUser, &userD)
		assert.NoError(b, err)
		assert.Equal(b, "Iman", userD.Name)
	}

}

func BenchmarkCreateUserFBS(b *testing.B) {
	b.N = BenchCall

	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(0)
		handler := &fbs.UserFbsHandler{}
		var user = &users.UserObj{
			ID:   64,
			Name: "Iman",
		}

		res := handler.MakeUser(builder, user)
		decoded := handler.ReadUser(res)
		assert.Equal(b, "Iman", decoded.Name)
	}

}

func BenchmarkCreateUserListFBS(b *testing.B) {
	b.N = BenchCall
	for j := 0; j < b.N; j++ {
		list := make([]*users.UserObj, httpdlv.DATA_SIZE)

		for i := 0; i < httpdlv.DATA_SIZE; i++ {
			var user = &users.UserObj{
				ID:   64,
				Name: "Iman",
			}
			list[i] = user
		}

		builder := flatbuffers.NewBuilder(0)
		handler := &fbs.UserFbsHandler{}
		resBuf := handler.MakeListUser(builder, list)
		decoded := handler.ReadUserList(resBuf)

		assert.Len(b, decoded, httpdlv.DATA_SIZE)
		assert.Equal(b, "Iman", decoded[0].Name)
	}

}

func BenchmarkCreateUserListJSON(b *testing.B) {
	b.N = BenchCall
	for j := 0; j < b.N; j++ {
		list := make([]*users.UserObj, httpdlv.DATA_SIZE)

		for i := 0; i < httpdlv.DATA_SIZE; i++ {
			var user = &users.UserObj{
				ID:   64,
				Name: "Iman",
			}
			list[i] = user
		}

		byt, _ := json.Marshal(list)
		// jsonList := string(byt)
		jsonList := byt
		var userLstD []*users.UserObj
		// err := json.Unmarshal([]byte(jsonList), &userLstD)
		err := json.Unmarshal(jsonList, &userLstD)
		assert.NoError(b, err)
		assert.Len(b, userLstD, httpdlv.DATA_SIZE)
		assert.Equal(b, "Iman", userLstD[0].Name)
	}

}
