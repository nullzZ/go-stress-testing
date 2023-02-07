package requests

import (
	"bytes"
	"encoding/binary"
	"github.com/link1st/go-stress-testing/message"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"strconv"
	"time"
)

func NewAuthC() string {
	accId := strconv.Itoa(int(time.Now().UnixMilli())) + strconv.Itoa(rand.Intn(1000000))
	authC := &message.Auth_C{
		ToServerId: 0,
		AccountId:  accId,
	}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, message.MSG_TYPE_E_Auth_C)
	bb, _ := proto.Marshal(authC)
	buf.Write(bb)
	return buf.String()
}

func NewTestC() string {
	accId := strconv.Itoa(int(time.Now().UnixMilli())) + strconv.Itoa(rand.Intn(1000000))
	authC := &message.Test_C{
		AccountId: accId,
	}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, message.MSG_TYPE_E_Test_C)
	bb, _ := proto.Marshal(authC)
	buf.Write(bb)
	return buf.String()
}
