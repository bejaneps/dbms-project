package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/bejaneps/dbms-project/auth"

	"github.com/bejaneps/dbms-project/handlers"

	"github.com/gin-gonic/gin"
)

var (
	logger   *os.File
	recovery *os.File
)

func init() {
	//creating files for logs and recovery
	var err error

	logger, err = os.Create("logs/log.txt")
	if err != nil {
		log.Fatal(err)
	}

	recovery, err = os.Create("logs/recovery.txt")
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode("debug")
}

func main() {
	//assigning multiplexer and router for server
	router := gin.New()
	router.Use(gin.LoggerWithWriter(logger), gin.RecoveryWithWriter(recovery))

	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.IndexHandler)
	router.GET("/login", handlers.LoginHandler)
	router.GET("/login/submit", handlers.LoginSubmitHandler)

	router.GET("/register", handlers.RegisterHandler)
	router.GET("/register/submit", handlers.RegisterSubmitHandler)

	router.GET("/dashboard", handlers.DashboardHandler)

	router.GET("/dashboard/cdr/course", handlers.CdrCourseHandler)
	router.GET("/dashboard/cdr/teacher", handlers.CdrTeacherHandler)
	router.GET("/dashboard/cdr/student", handlers.CdrStudentHandler)
	router.GET("/dashboard/cdr/assign/course", handlers.CdrAssignCourseHandler)
	router.GET("/dashboard/cdr/assign/instructor", handlers.CdrAssignInstructorHandler)

	router.GET("/dashboard/tcr/courses", handlers.TcrListCoursesHandler)

	router.GET("/dashboard/std/courses", handlers.StdListCoursesHandler)

	router.GET("/logout", handlers.LogoutHandler)

	//starting a server and serving on tcp
	var server = &http.Server{
		Addr:    ":5050",
		Handler: router,
	}

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//graceful shutdown part
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()

	logger.Close()
	recovery.Close()
}
