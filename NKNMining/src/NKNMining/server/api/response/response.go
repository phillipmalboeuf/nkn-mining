package apiServerResponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
	"NKNMining/common"
)

type ResponseData struct {
	Data interface{}
	Timestamp string
	ApiVersion string
}

type Response struct {
	ctx   *gin.Context
	mutex *sync.Mutex
}

func New(ctx *gin.Context) *Response {
	return &Response{
		ctx:   ctx,
		mutex: new(sync.Mutex),
	}

}

func (r *Response) BadRequest(data interface{}) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.ctx.JSON(http.StatusBadRequest, data)
}

func (r *Response) Forbidden(data interface{}) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.ctx.JSON(http.StatusForbidden, data)
}


func (r *Response) InternalServerError(data interface{}) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.ctx.JSON(http.StatusInternalServerError, data)
}

func (r *Response) Success(data interface{}) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	responseData := &ResponseData{
		data,
		time.Now().Format("2006-01-02 15:04:05"),
		common.API_VERSION,
	}
	r.ctx.JSON(http.StatusOK, responseData)
}
