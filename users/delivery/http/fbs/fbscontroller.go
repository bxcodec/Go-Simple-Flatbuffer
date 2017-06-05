package fbs

import (
	"github.com/bxcodec/Go-Simple-Flatbuffer/users"
	flatbuffers "github.com/google/flatbuffers/go"
)

type UserFbsHandler struct {
}

func (u *UserFbsHandler) MakeListUser(b *flatbuffers.Builder, listUser []*users.UserObj) []byte {

	b.Reset()
	ptrs := make([]flatbuffers.UOffsetT, len(listUser))

	for k, v := range listUser {
		name := b.CreateString(v.Name)
		users.UserStart(b)
		users.UserAddName(b, name)
		users.UserAddId(b, v.ID)
		userPosition := users.UserEnd(b)
		ptrs[k] = userPosition
	}
	users.UserContainerStartListOfUsersVector(b, len(listUser))
	for i := len(listUser) - 1; i >= 0; i-- {
		b.PrependUOffsetT(ptrs[i])
	}
	vptr := b.EndVector(len(listUser))
	users.UserContainerStart(b)
	users.UserContainerAddListOfUsers(b, vptr)
	container := users.UserContainerEnd(b)
	b.Finish(container)
	return b.Bytes[b.Head():]
}

func (u *UserFbsHandler) MakeUser(b *flatbuffers.Builder, userobj *users.UserObj) []byte {
	// re-use the already-allocated Builder:
	b.Reset()

	// create the name object and get its offset:
	name_position := b.CreateString(userobj.Name)

	// write the User object:
	users.UserStart(b)
	users.UserAddName(b, name_position)
	users.UserAddId(b, userobj.ID)
	user_position := users.UserEnd(b)

	// finish the write operations by our User the root object:
	b.Finish(user_position)

	// return the byte slice containing encoded data:
	return b.Bytes[b.Head():]
}

func (u *UserFbsHandler) ReadUser(buf []byte) *users.UserObj {
	// initialize a User reader from the given buffer:
	user := users.GetRootAsUser(buf, 0)

	// point the name variable to the bytes containing the encoded name:
	name := user.Name()

	// copy the user's id (since this is just a int64):
	id := user.Id()

	userobj := &users.UserObj{id, string(name)}

	return userobj
}
func (u *UserFbsHandler) ReadUserList(buff []byte) []*users.UserObj {
	container := users.GetRootAsUserContainer(buff, 0)

	res := make([]*users.UserObj, container.ListOfUsersLength())
	for i := 0; i < len(res); i++ {
		u := &users.User{}
		container.ListOfUsers(u, i)
		ubj := &users.UserObj{
			ID:   u.Id(),
			Name: string(u.Name()),
		}
		res[i] = ubj
	}
	return res
}
