package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tehdev/summoner-rift-api/config"
	"github.com/tehdev/summoner-rift-api/databases"

)

type echoServer struct {
	app  *echo.Echo     //จัดการ พวก route
	db   databases.Database       //ที่เก็บวัตถุดิบไรงี้
	conf *config.Config //การตั้งค่าต่าง ๆ ของโปรเจค
}

var (
	once   sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config, db databases.Database) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	//ไม่ว่าจะเรียกฟังก์ชันนี้กี่ครั้งก็ตาม ให้สร้างเซิร์ฟเวอร์ขึ้นมาแค่ครั้งแรกครั้งเดียวเท่านั้น
	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			db:   db,
			conf: conf,
		}

	})

	return server
}

// Start Server
func (s *echoServer) Star() {

	corMiddleware := getCORSMiddleware(s.conf.Server.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.conf.Server.BodyLimit)
	timeoutMiddleware := getTimeOutMiddleware(s.conf.Server.Timeout)

	//Call Middleware
	s.app.Use(middleware.Recover()) //ดักจับข้อผิดพลาดร้ายแรง (panic)

	s.app.Use(corMiddleware)
	s.app.Use(bodyLimitMiddleware)
	s.app.Use(timeoutMiddleware)

	s.app.Use(middleware.Logger())

	s.app.GET("/v1/health", s.healthCheck) //"กำหนดเส้นทาง"


	s.initCardShopRouter()
	s.initCardManagingRouter()
	

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefullyShutdown(quitCh)

	s.httpListening() //เปิดทำงาน
}

// จัดการพวกคำขอของ HTTP ที่เป็น infinity loop
func (s *echoServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port) // อ่านค่า Port ที่จะเปิดใช้งานมาจากการตั้งค่าconf

	//ถ้าเกิด Error จะปิดไปเลยน้ิองง
	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

// Gracefully
func (s *echoServer) gracefullyShutdown(quitCh chan os.Signal) {
	ctx := context.Background()

	<-quitCh
	s.app.Logger.Info("Shutting down server...")

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error : %s", err.Error())

	}

}

// Respone"OK" และสถานะ 200 ซึ่งเป็นการบอกว่า "เซิร์ฟเวอร์ยังทำงานปกติดีอยู่!"
func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// กรณี client ส่ง Request มานานนเกินกำหนด
func getTimeOutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Request Timeout",
		Timeout:      timeout * time.Second, //แปลงเป็นวินาทีงับ
	})

}

// กัน Client แปลก ๆ ยิงเข้ามา
func getCORSMiddleware(allowOrigins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: allowOrigins,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}

// กัน payload เกินนน
func getBodyLimitMiddleware(bodylimit string) echo.MiddlewareFunc {
	return middleware.BodyLimit(bodylimit)
}
