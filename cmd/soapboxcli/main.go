package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/adhocteam/soapbox/buildinfo"
	pb "github.com/adhocteam/soapbox/proto"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"google.golang.org/grpc"
)

var dryRun bool

func main() {
	serverAddr := flag.String("server", "127.0.0.1:9090", "host:port of server")
	printVersion := flag.Bool("V", false, "print version and exit")

	flag.BoolVar(&dryRun, "dry-run", false, "do not perform any changes")
	flag.Parse()

	if *printVersion {
		fmt.Printf("        version: %s\n", buildinfo.Version)
		fmt.Printf("     git commit: %s\n", buildinfo.GitCommit)
		fmt.Printf("     build time: %s\n", buildinfo.BuildTime)
		return
	}

	if flag.NArg() < 1 {
		usage()
		os.Exit(1)
	}

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect to server %s: %v", *serverAddr, err)
	}
	defer conn.Close()

	client := pb.NewApplicationsClient(conn)
	ctx := context.Background()

	type command func(context.Context, pb.ApplicationsClient, []string) error
	var cmd command

	switch flag.Arg(0) {
	case "create-application":
		cmd = createApplication
	case "list-applications":
		cmd = listApplications
	case "get-application":
		cmd = getApplication
	case "get-version":
		if err := getVersion(ctx, pb.NewVersionClient(conn), nil); err != nil {
			log.Fatalf("getting version: %v", err)
		}
		return
	case "cleanup":
		cmd = deployCleanup
	default:
		log.Fatalf("unknown command %q", flag.Arg(0))
	}

	if err := cmd(ctx, client, flag.Args()); err != nil {
		log.Fatalf("error executing command %s: %v", flag.Arg(0), err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <command>\n", filepath.Base(os.Args[0]))
}

func createApplication(ctx context.Context, client pb.ApplicationsClient, args []string) error {
	args = args[1:]
	if len(args) < 4 {
		return fmt.Errorf("4 arguments are required: name, description, github repo URL, and type (server, cronjob)")
	}

	var appType pb.ApplicationType
	switch args[3] {
	case "server":
		appType = pb.ApplicationType_SERVER
	case "cronjob":
		appType = pb.ApplicationType_CRONJOB
	default:
		return fmt.Errorf("unknown app type %q, expecting either 'server' or 'cronjob'", args[3])
	}

	req := &pb.Application{
		Name:          args[0],
		Description:   args[1],
		GithubRepoUrl: args[2],
		Type:          appType,
	}
	app, err := client.CreateApplication(ctx, req)
	if err != nil {
		return fmt.Errorf("error creating application: %v", err)
	}

	fmt.Printf("created application %q, ID %d", args[0], app.GetId())
	return nil
}

func listApplications(ctx context.Context, client pb.ApplicationsClient, args []string) error {
	apps, err := client.ListApplications(ctx, new(pb.Empty))
	if err != nil {
		return fmt.Errorf("error listing applications: %v", err)
	}
	for _, app := range apps.Applications {
		fmt.Printf("%d\t%s\n", app.Id, app.Name)
	}
	return nil
}

func getApplication(ctx context.Context, client pb.ApplicationsClient, args []string) error {
	args = args[1:]
	if len(args) < 1 {
		return fmt.Errorf("1 argument required: ID of application")
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid ID: %v", err)
	}
	req := &pb.GetApplicationRequest{Id: int32(id)}
	app, err := client.GetApplication(ctx, req)
	if err != nil {
		return fmt.Errorf("getting application: %v", err)
	}
	fmt.Printf("name:                %s\n", app.Name)
	fmt.Printf("ID:                  %d\n", app.Id)
	fmt.Printf("type:                %s\n", pb.ApplicationType_name[int32(app.Type)])
	fmt.Printf("created at:          %s\n", app.CreatedAt)
	fmt.Printf("external DNS:        %s\n", app.ExternalDns)
	fmt.Printf("Dockerfile path:     %s\n", app.DockerfilePath)
	fmt.Printf("entrypoint override: %s\n", app.EntrypointOverride)
	fmt.Printf("description:\n%s\n", app.Description)
	return nil
}

func getVersion(ctx context.Context, client pb.VersionClient, args []string) error {
	resp, err := client.GetVersion(ctx, &pb.Empty{})
	if err != nil {
		return fmt.Errorf("getting version: %v", err)
	}
	fmt.Println("Soapbox API")
	fmt.Println("-----------")
	fmt.Printf("    version: %s\n", resp.Version)
	fmt.Printf(" git commit: %s\n", resp.GitCommit)
	fmt.Printf(" build time: %s\n", resp.BuildTime)
	return nil
}

func deployCleanup(ctx context.Context, client pb.ApplicationsClient, args []string) error {
	args = args[1:]
	if len(args) < 1 {
		return fmt.Errorf("1 argument is required: application name")
	}

	appName := args[0]

	sess, _ := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String("us-east-1")},
	})
	svc := autoscaling.New(sess)

	descAsgInput := autoscaling.DescribeAutoScalingGroupsInput{}
	asgRes, err := svc.DescribeAutoScalingGroups(&descAsgInput)
	if err != nil {
		return fmt.Errorf("describing autoscaling groups: %s", err)
	}

	var inUseLcs map[string]bool
	inUseLcs = make(map[string]bool)
	for _, asg := range asgRes.AutoScalingGroups {
		inUseLcs[*asg.LaunchConfigurationName] = true
	}

	descLcInput := autoscaling.DescribeLaunchConfigurationsInput{}
	lcRes, err := svc.DescribeLaunchConfigurations(&descLcInput)
	if err != nil {
		return fmt.Errorf("describing autoscaling groups: %s", err)
	}

	var unusedLcs []string
	for _, lc := range lcRes.LaunchConfigurations {
		if !inUseLcs[*lc.LaunchConfigurationName] {
			unusedLcs = append(unusedLcs, *lc.LaunchConfigurationName)
		}
	}

	for _, lc := range unusedLcs {
		if strings.HasPrefix(lc, appName) {
			if !dryRun {
				fmt.Println("deleting:", lc)
				delLcInput := autoscaling.DeleteLaunchConfigurationInput{
					LaunchConfigurationName: aws.String(lc),
				}
				_, err := svc.DeleteLaunchConfiguration(&delLcInput)
				if err != nil {
					return fmt.Errorf("deleting launch config: %s", err)
				}
			} else {
				fmt.Println("would delete:", lc)
			}
		}
	}
	return nil
}
