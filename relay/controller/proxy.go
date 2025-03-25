// Package controller is a package for handling the relay controller
package controller

import (
	"fmt"
	"net/http"

	"github.com/9688101/HX/common/logger"
	"github.com/9688101/HX/relay"
	"github.com/9688101/HX/relay/adaptor/openai"
	"github.com/9688101/HX/relay/meta"
	relaymodel "github.com/9688101/HX/relay/model"
	"github.com/gin-gonic/gin"
)

// RelayProxyHelper is a helper function to proxy the request to the upstream service
func RelayProxyHelper(c *gin.Context, relayMode int) *relaymodel.ErrorWithStatusCode {
	ctx := c.Request.Context()
	meta := meta.GetByContext(c)

	adaptor := relay.GetAdaptor(meta.APIType)
	if adaptor == nil {
		return openai.ErrorWrapper(fmt.Errorf("invalid api type: %d", meta.APIType), "invalid_api_type", http.StatusBadRequest)
	}
	adaptor.Init(meta)

	resp, err := adaptor.DoRequest(c, meta, c.Request.Body)
	if err != nil {
		logger.Errorf(ctx, "DoRequest failed: %s", err.Error())
		return openai.ErrorWrapper(err, "do_request_failed", http.StatusInternalServerError)
	}

	// do response
	_, respErr := adaptor.DoResponse(c, resp, meta)
	if respErr != nil {
		logger.Errorf(ctx, "respErr is not nil: %+v", respErr)
		return respErr
	}

	return nil
}
