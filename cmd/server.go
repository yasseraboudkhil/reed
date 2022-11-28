package cmd

import (
	"context"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/yasseraboudkhil/cloud-hsm/mykms"
	grpchandle "github.com/yasseraboudkhil/reed/handler/grpc"
	grpc "github.com/yasseraboudkhil/reed/interface/grpc"

	"github.com/yasseraboudkhil/reed/conf"
)

func init() {

	serverCmd := &cobra.Command{
		Use:   "reed-svc",
		Short: "Starts the reed server",
		Long:  `Starts REED (i.e Remote Envelope Encryption Dictionary) to act as an isolated and central encrypt/decrypt service`,
		Run:   RunServer,
	}
	rootCmd.AddCommand(serverCmd)
}

func RunServer(cmd *cobra.Command, args []string) {

	//ctx context.Context, toDoSvc v1.ToDoServiceServer, port string

	log.Println("INSIDE cmd.server.RunServer()")

	ctx := context.Background()

	kmsClient, err := mykms.New(conf.C.Aws.Region, conf.C.Aws.AccessKey, conf.C.Aws.SecretKey, conf.C.Kms.KeyId)

	if err != nil {
		log.Fatal("Could not create a KMS client")
	}

	reedSvcServer := grpchandle.NewReedServiceServer(kmsClient)

	if err := grpc.RunServer(ctx, reedSvcServer, strconv.Itoa(conf.C.Server.Port)); err != nil {
		log.Fatal("Could not server for this reason: " + err.Error())
	}

}
