package main

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"

	"baseapp/app/cli"
	"baseapp/app/configs"
	"baseapp/app/crons"
	"baseapp/app/database"
	"baseapp/app/utils"
	"baseapp/graph"

	controllers "baseapp/app/http/controllers"
	appMiddleware "baseapp/app/http/middleware"
)

func main() {

	configs.InitConfig()
	database.InitDB()

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "This is a CLI application with Echo server",
	}

	var serverCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start the Echo server",
		Run: func(cmd *cobra.Command, args []string) {
			runEchoServer()
		},
	}

	var taskApp = &cobra.Command{
		Use:   "task",
		Short: "Application Tasks",
		Run: func(cmd *cobra.Command, args []string) {
			cli.CliRunTask(args)
		},
	}

	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(taskApp)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

}

func runEchoServer() {
	//echo start here
	e := echo.New()

	// Middleware
	if utils.InArray(configs.AppConfig.APPLogLevel, []string{"INFO"}) {
		e.Use(middleware.Logger())
	}
	if utils.InArray(configs.AppConfig.APPLogLevel, []string{"INFO", "ERROR"}) {
		e.Use(middleware.Recover())
	}

	e.Use(appMiddleware.TrimMiddleware)
	e.Validator = utils.NewCustomValidator()

	//init cron
	crons.NewAppCron().RunCron()

	//base route
	e.GET("/", controllers.NewWelcomeController().Index)
	e.GET("/health", controllers.NewWelcomeController().Health)

	// GraphQL Handler
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// Playground Handler (for testing in-browser)
	e.GET("/playground", echo.WrapHandler(playground.Handler("GraphQL playground", "/graphql")))
	// GraphQL Query Handler
	e.POST("/graphql", echo.WrapHandler(srv))

	// Start server
	e.Logger.Fatal(e.Start(":3600"))

	printRoutes(e)

	e.Logger.Fatal(e.Start(":" + configs.AppConfig.APPPort))
}

// printRoutes prints all registered routes in the Echo instance
func printRoutes(e *echo.Echo) {
	fmt.Printf("Route APP ==========================================\n")
	for _, route := range e.Routes() {
		fmt.Printf("Method: %s,\t Path: %s,\t\t Name: %s \n", route.Method, route.Path, route.Name)
	}
	fmt.Printf("Route APP ==========================================\n")
}
