package main

import (
	"errors"
	"mobile-ecommerce/command"
	"mobile-ecommerce/config"
	property_postgres "mobile-ecommerce/internal/Properties/repository/postgres"
	auth_http "mobile-ecommerce/internal/auth/transport/http"
	auth_usecase "mobile-ecommerce/internal/auth/usecase"
	brand_postgres "mobile-ecommerce/internal/brands/repository/postgres"
	brand_http "mobile-ecommerce/internal/brands/transport/http"
	brand_usecase "mobile-ecommerce/internal/brands/usecase"
	cart_postgres "mobile-ecommerce/internal/carts/repository/postgres"
	cart_http "mobile-ecommerce/internal/carts/transport/http"
	cart_usecase "mobile-ecommerce/internal/carts/usercase"
	category_postgres "mobile-ecommerce/internal/categories/repository/postgres"
	category_http "mobile-ecommerce/internal/categories/transport/http"
	category_usecase "mobile-ecommerce/internal/categories/usecase"
	common_postgres "mobile-ecommerce/internal/common/repository/postgres"
	product_postgres "mobile-ecommerce/internal/products/repository/postgres"
	product_http "mobile-ecommerce/internal/products/transport/http"
	product_usecase "mobile-ecommerce/internal/products/usecase"
	user_postgres "mobile-ecommerce/internal/users/repository/postgresql"
	user_http "mobile-ecommerce/internal/users/transport/http"
	user_usecase "mobile-ecommerce/internal/users/usecase"
	"mobile-ecommerce/internal/util/token"
	"mobile-ecommerce/pkg/middleware"
	"net/http"
	"time"

	_ "mobile-ecommerce/docs"

	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title mobile-ecommerce project by Nguyễn Khánh Duy
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	err := run()
	if err != nil {
		logrus.Fatalf("error when run main: %s", err.Error())
	}

}

func run() error {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal("get error when load env file")
	}

	cfg := config.Load()

	cmd := &cobra.Command{
		Use: "server",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "run server",
		Run: func(cmd *cobra.Command, args []string) {
			serveServer(cfg)
		},
	}, command.MigrationCommand(cfg.PostGres))

	return cmd.Execute()

}

func serveServer(cfg *config.Config) {
	// Connect DB
	gormDB, err := connectGormDB(cfg.PostGres)
	if err != nil {
		logrus.Fatal("get error when connect database:", err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.CORSMiddleware())
	r.Static("/static", "./public")
	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello world!")
	})

	engine := r.Group("api/v1")

	tokenMaker := token.NewJwtMaker(string(token.SECRET_KEY), time.Duration(10*time.Minute))
	commonRepo := common_postgres.NewCommonRepository(gormDB)
	userRepo := user_postgres.NewUserRepository(gormDB)
	brandRepo := brand_postgres.NewBrandRepository(gormDB)
	categoryRepo := category_postgres.NewCategoryRepo(gormDB)
	cartRepo := cart_postgres.NewCartRepository(gormDB)
	productRepo := product_postgres.NewProductRepository(gormDB)
	propertyRepo := property_postgres.NewPropertyRepository(gormDB)

	userUsecase := user_usecase.NewUserUsecase(userRepo)
	authUsecase := auth_usecase.NewAuthUsecase(userRepo, tokenMaker)
	brandUsecase := brand_usecase.NewBrandUsecase(brandRepo, userRepo)
	categoryUsecase := category_usecase.NewCategoryUsecase(categoryRepo, userRepo)
	cartUsecase := cart_usecase.NewCartUsecase(cartRepo, userRepo)
	productUseCase := product_usecase.NewProductUsecase(productRepo, userRepo, propertyRepo, commonRepo)

	user_http.NewUserHandler(engine, userUsecase, tokenMaker)
	auth_http.NewAuthHandler(engine, authUsecase)
	brand_http.NewBrandHandler(engine, brandUsecase, tokenMaker)
	category_http.NewCategoryHandler(engine, categoryUsecase, tokenMaker)
	cart_http.NewCartHandler(engine, cartUsecase, tokenMaker)
	product_http.NewProductHandler(engine, productUseCase, tokenMaker)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler))
	logrus.Info("view docs api: http://localhost:8080/swagger/index.html")
	r.Run()
}

func connectGormDB(cfg config.DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.Source), &gorm.Config{})
	if err != nil {
		return nil, errors.New("get error when connect database:" + err.Error())
	}

	return db, nil
}
