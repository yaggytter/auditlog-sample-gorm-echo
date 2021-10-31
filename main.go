package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/yaggytter/auditlog-sample-gorm-echo/auditlog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type authContext struct {
	echo.Context
	UserName string
}

func main() {
	connectDB()
	router := newRouter()
	log.Fatal(router.Start(":8080"))
}

func newRouter() *echo.Echo {
	e := echo.New()
	//	e.Use(debugMiddleware)
	e.Use(authMiddleware)

	e.GET("/logtest", logtest)
	return e
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// いろいろ認可の処理をしたあと
		fmt.Println("auth ok!")
		// 例えばユーザ名を入れる
		username := "User1"
		// コンテキストにユーザ名を入れる
		ctx := context.WithValue(c.Request().Context(), "UserName", username)
		r := c.Request().WithContext(ctx)
		c.SetRequest(r)

		return next(c)
	}
}

func debugMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// デバッグ用途
		log.Println("before action")
		log.Printf("cookie = %v", c.Cookies())
		log.Printf("ref = %v", c.Request().Referer())
		if err := next(c); err != nil {
			c.Error(err)
		}
		log.Println("after action")
		return nil
	}
}

// DB
type user struct {
	Id           int    `json:"id" param:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Secret       string `json:"secret"`
	Personalinfo string `json:"personalinfo"`
}

var db *gorm.DB

func connectDB() {
	user := "mysqluser"
	password := "mysqlpass"
	host := "db"
	port := "3306"
	database_name := "auditlogtest"

	var err error
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: auditlog.Default.LogMode(logger.Info), // for Audit logs
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func logtest(c echo.Context) error {
	ctxdb := db.WithContext(c.Request().Context())
	sqlDB, _ := ctxdb.DB()
	defer sqlDB.Close()
	err := sqlDB.Ping()

	users := []user{}
	ctxdb.Find(&users)
	//	fmt.Printf("%v\n", users)

	if err != nil {
		return c.String(http.StatusInternalServerError, "can not connect to database")
	} else {
		return c.JSON(http.StatusOK, users)
	}
}
