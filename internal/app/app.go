package app

import (
	"embed"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"github.com/9688101/HX/common"
	"github.com/9688101/HX/common/client"
	"github.com/9688101/HX/common/config"
	"github.com/9688101/HX/common/i18n"
	"github.com/9688101/HX/common/logger"
	"github.com/9688101/HX/internal/controller/http"
	"github.com/9688101/HX/internal/controller/http/middleware"
	v1 "github.com/9688101/HX/internal/controller/http/v1"
	"github.com/9688101/HX/internal/entity"
	"github.com/9688101/HX/relay/adaptor/openai"
)

// go:embed ../web/build/*
var buildFS embed.FS

func Run() {
	
}
