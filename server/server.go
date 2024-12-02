package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/swagger"
)

type Server struct {
	ServerReady chan bool
	APP_LOGGER  bool
	APP_PORT    string
}

// func (s *Server) connectMongoDBTest(ctx context.Context) *mongo.Database {
// 	var (
// 		client *mongo.Client
// 		err    error
// 	)
// 	if client, err = mongo.Connect(ctx, options.Client().ApplyURI(s.MONGODB_CONNECTION_URI)); err != nil {
// 		panic(err)
// 	}
// 	if err = client.Ping(ctx, nil); err != nil {
// 		panic(err)
// 	}
// 	return client.Database("aec-form")
// }

// func (s *Server) connectMongoDB(ctx context.Context) *mongo.Database {
// 	var (
// 		logg   = options.Logger().SetComponentLevel(options.LogComponentCommand, 0)
// 		client *mongo.Client
// 		err    error
// 	)
// 	if client, err = mongo.Connect(ctx, options.Client().ApplyURI(s.MONGODB_CONNECTION_URI).SetLoggerOptions(logg)); err != nil {
// 		panic(err)
// 	}
// 	if err = client.Ping(ctx, nil); err != nil {
// 		panic(err)
// 	}
// 	if err = utils.CreateFormIndex(client.Database("aec-form")); err != nil {
// 		panic(err)
// 	}
// 	return client.Database("aec-form")
// }

// func (s *Server) connectMinio() *minio.Client {
// 	var (
// 		ctx  context.Context = context.Background()
// 		opts *minio.Options  = &minio.Options{
// 			Creds:  credentials.NewStaticV4(s.MINIO_ACCESS_KEY, s.MINIO_SECRET_KEY, ""),
// 			Secure: s.MINIO_SSL,
// 			Region: s.MINIO_REGION,
// 		}
// 		existsBucket bool = false
// 		client       *minio.Client
// 		err          error
// 	)
// 	if client, err = minio.New(s.MINIO_ENDPOINT, opts); err != nil {
// 		logrus.Errorln(err)
// 	}

// 	if existsBucket, err = client.BucketExists(ctx, s.MINIO_DEFAULT_BUCKET); err != nil {
// 		logrus.Errorln(err)
// 	}
// 	if !existsBucket {
// 		var opts minio.MakeBucketOptions = minio.MakeBucketOptions{
// 			Region: s.MINIO_REGION,
// 		}
// 		if err = client.MakeBucket(ctx, s.MINIO_DEFAULT_BUCKET, opts); err != nil {
// 			logrus.Errorln(err)
// 		}
// 		if err = client.SetBucketPolicy(ctx, s.MINIO_DEFAULT_BUCKET, constant.POLICY_PUBLIC(s.MINIO_DEFAULT_BUCKET)); err != nil {
// 			logrus.Errorln(err)
// 		}
// 	}
// 	return client
// }

// func (s *Server) StartGrpcServer(grpcServ *grpc.Server) {
// 	var (
// 		listen net.Listener
// 		err    error
// 	)
// 	if listen, err = net.Listen("tcp", ":"+s.GRPC_PORT); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("gRPC server is running on port:", s.GRPC_PORT)
// 	if err = grpcServ.Serve(listen); err != nil {
// 		panic(err)
// 	}
// }

func (s *Server) Start() {
	var (
		// ctx, cancel                 = context.WithTimeout(context.Background(), 10*time.Second)
		// minioClient *minio.Client   = s.connectMinio()
		// mdb         *mongo.Database = s.connectMongoDB(ctx)
		app *fiber.App
		// grpcServ *grpc.Server
		port string
		err  error
	)
	// defer cancel()
	// minioClient :=
	// Inititalize the http server
	app = fiber.New()
	if s.APP_LOGGER {
		app.Use(logger.New())
	}

	// Initialize the gRPC server
	// grpcServ = grpc.NewServer()
	// defer grpcServ.GracefulStop()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "POST, GET, PUT, DELETE",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//==============================================================
	// Repository
	//==============================================================

	//==============================================================
	// Usecase
	//==============================================================

	//==============================================================
	// Handler
	//==============================================================

	//==============================================================
	// Grpc
	//==============================================================

	//==============================================================
	// Grpc Route
	//==============================================================

	//==============================================================
	// Fiber Router
	//==============================================================

	app.Get("/swagger/*", swagger.HandlerDefault)

	route := routes.NewRoute(app)

	// go func() {
	// 	if r := recover(); r != nil {
	// 		s.ServerReady <- false
	// 		panic(r.(error))
	// 	} else {
	// 		s.StartGrpcServer(grpcServ)
	// 	}
	// }()

	port = ":" + s.APP_PORT
	if err = app.Listen(port); err != nil {
		panic(err)
	}
}