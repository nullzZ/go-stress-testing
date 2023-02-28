// Package golink 连接
package golink

import (
	"bytes"
	"compress/gzip"
	"context"
	"github.com/link1st/go-stress-testing/requests"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/link1st/go-stress-testing/model"
	"github.com/link1st/go-stress-testing/server/client"
)

// HTTP 请求
func HTTP(ctx context.Context, chanID uint64, ch chan<- *model.RequestResults, totalNumber uint64, wg *sync.WaitGroup,
	request *model.Request) {
	defer func() {
		wg.Done()
	}()
	HTTP2(ctx, chanID, ch, totalNumber, request)
	return
}

func HTTP2(ctx context.Context, chanID uint64, ch chan<- *model.RequestResults, totalNumber uint64,
	request *model.Request) {
	// fmt.Printf("启动协程 编号:%05d \n", chanID)
	for i := uint64(0); i < totalNumber; i++ {
		if ctx.Err() != nil {
			//fmt.Printf("ctx.Err err: %v \n", ctx.Err())
			break
		}
		//body2 := requests.NewTestC()
		buf := bytes.Buffer{}
		buf.WriteString(strconv.Itoa(int(chanID)))
		buf.WriteString("_")
		buf.WriteString(strconv.Itoa(int(i)))
		body2 := requests.NewTestC2(buf.String())
		request.Body = body2
		//todo!!!!!!!!!!!!!
		list := getRequestList(request)

		isSucceed, errCode, requestTime, contentLength := sendList(chanID, list)
		requestResults := &model.RequestResults{
			Time:          requestTime,
			IsSucceed:     isSucceed,
			ErrCode:       errCode,
			ReceivedBytes: contentLength,
		}
		requestResults.SetID(chanID, i)
		ch <- requestResults
	}

	return
}

// sendList 多个接口分步压测
func sendList(chanID uint64, requestList []*model.Request) (isSucceed bool, errCode int, requestTime uint64,
	contentLength int64) {
	errCode = model.HTTPOk
	for _, request := range requestList {
		succeed, code, u, length := send(chanID, request)
		isSucceed = succeed
		errCode = code
		requestTime = requestTime + u
		contentLength = contentLength + length
		if succeed == false {
			break
		}
	}
	return
}

// send 发送一次请求
func send(chanID uint64, request *model.Request) (bool, int, uint64, int64) {
	var (
		// startTime = time.Now()
		isSucceed     = false
		errCode       = model.HTTPOk
		contentLength = int64(0)
		err           error
		resp          *http.Response
		requestTime   uint64
	)
	newRequest := getRequest(request)

	resp, requestTime, err = client.HTTPRequest(chanID, newRequest)

	if err != nil {
		errCode = model.RequestErr // 请求错误
	} else {
		// 此处原方式获取的数据长度可能是 -1，换成如下方式获取可获取到正确的长度
		contentLength, err = getBodyLength(resp)
		if err != nil {
			contentLength = resp.ContentLength
		}
		// 验证请求是否成功
		errCode, isSucceed = newRequest.GetVerifyHTTP()(newRequest, resp)
	}
	return isSucceed, errCode, requestTime, contentLength
}

// getBodyLength 获取响应数据长度
func getBodyLength(response *http.Response) (length int64, err error) {
	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		defer func() {
			_ = reader.Close()
		}()
	default:
		reader = response.Body
	}
	body, err := ioutil.ReadAll(reader)
	return int64(len(body)), err
}
