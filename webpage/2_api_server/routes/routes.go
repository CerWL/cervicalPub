package routes

import (
	"net/http"
	"time"

	ctr "github.com/paulxiong/cervical/webpage/2_api_server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Router 注册路由
func Router() *gin.Engine {
	r := gin.New()
	// 图片服务器API,放在这里目的是不打印log，因为这个是动态路由
	r.GET("/imgs/*any", ctr.ImageAPI)
	// swager,放在这里目的是不打印log，因为这个是动态路由
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "RELEASE"))

	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/api1/ping", "/api1/job"}}))
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.StaticFile("/favicon.ico", "./web/dist/favicon.ico")
	r.StaticFile("/", "./web/dist/")
	r.StaticFS("/static", http.Dir("./web/dist/static"))
	r.NoRoute(func(c *gin.Context) { c.JSON(404, gin.H{"text": "Not Found."}) })

	corsObject := cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:    []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With", "testHeader"},
		MaxAge:          12 * time.Hour,
	})
	r.Use(corsObject)

	r.Use(ctr.History)

	/* 用户相关的API */
	user := r.Group("/user")
	user.POST("/register", ctr.RegisterUser)
	user.POST("/login", ctr.LoginUser)
	user.POST("/emailcode", ctr.GetEmailCode)
	user.Use(ctr.CheckAuth)
	{
		user.GET("/info", ctr.GetUser)
		user.GET("/logout", ctr.LogoutUser)
		user.GET("/userinfo", ctr.GetUser)
		user.GET("/accesslog", ctr.GetAccessLog)
		user.GET("/lists", ctr.GetUserLists)
		user.POST("/updateinfo", ctr.UpdateUserInfo)
	}

	api1 := r.Group("/api1")
	api1.GET("/ping", ctr.Pong) //不需要登录就能ping的API
	// 任务
	api1.POST("/job", ctr.GetOneJob)
	api1.POST("/jobresult", ctr.SetJobResult)

	api1.Use(ctr.CheckAuth)
	{
		// api1.GET("/refresh_token", ctr.AuthMiddleware.RefreshHandler) // Refresh time can be longer than token timeout
		api1.GET("/authping", ctr.AuthPong)
		// 数据
		api1.GET("/dtinfo", ctr.AllInfo)
		api1.GET("/batchinfo", ctr.GetBatchInfo)
		api1.GET("/medicalidinfo", ctr.GetMedicalIDInfo)
		api1.GET("/categoryinfo", ctr.GetCategoryInfo)
		api1.POST("/imglistsofwanted", ctr.GetImgListOfWanted)
		api1.GET("/imglistsonebyone", ctr.GetImgListOneByOne)
		api1.POST("/getimgnptypebymids", ctr.GetImagesNPTypeByMedicalID)
		api1.GET("/getimgbymid", ctr.GetImgListOfMedicalID)
		// api1.GET("/getimgtree", ctr.GetImgTree)

		// 标注
		api1.POST("/updatelabelsofimage", ctr.UpdateLabelsOfImage)
		api1.GET("/getLabelbyimageid", ctr.GetLabelByImageID)

		// 数据集
		api1.POST("/createdataset", ctr.CreateDataset)
		api1.GET("/listdatasets", ctr.ListDatasets)

		// 项目
		api1.POST("/createproject", ctr.CreateProject)
		api1.GET("/listprojects", ctr.ListProjects)

		// 任务
		api1.GET("/jobresult", ctr.GetJobResult)
		api1.GET("/jobpercent", ctr.GetJobPercent)
		api1.GET("/joblog", ctr.GetJobLog)
		api1.GET("/trainresult", ctr.GetTrainResult)
		api1.GET("/predictresult", ctr.GetPredictResult)
		// 模型
		api1.GET("/listmodel", ctr.GetModelLists)
		api1.POST("/savemodel", ctr.SaveModelInfo)

		// 文件操作
		api1.GET("/zipdownload", ctr.FileDownload)
		api1.POST("/uploads", ctr.UploadsHandler)
		api1.POST("/upload", ctr.UploadDatasetHandler)
		api1.POST("/uploadimg", ctr.UploadImgHandler)

		// 系统
		api1.POST("/errorlog", ctr.CreateErrorLog)
		api1.GET("/errorlog", ctr.GetErrorLog)
		api1.POST("/registeremailcfg", ctr.RegisterEmailCfg)
		api1.GET("/registeremailcfg", ctr.GetRegisterEmailCfg)

		// 医疗报告
		api1.GET("/predictresult2", ctr.GetPredictResult2)
		api1.GET("/datasetimgs", ctr.GetPredictImges)
		api1.POST("/updatepredict", ctr.UpdatePredict)
		api1.POST("/reviewpredict", ctr.ReviewPredict)
	}

	return r
}
