package csv

import (
	"context"
	"encoding/csv"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"math"
	"csvreader/internal/pb"
	"csvreader/pkg/apierror"
	"runtime"
	"strconv"
	"strings"
)

type UserProcessor struct {
	client    pb.UserServiceClient
	batchSize int64
}

func NewUserProcessor(conn *grpc.ClientConn, bs int64) UserProcessor {
	return UserProcessor{
		client:    pb.NewUserServiceClient(conn),
		batchSize: bs,
	}
}

func (u UserProcessor) Process(ctx context.Context, file io.Reader) (errRet apierror.ErrorMessage) {

	defer func() {
		if !errRet.R.HasErrors() {
			errRet.Message = "completed successfully"
		}
	}()

	// number of workers from number of CPU cores - minus one for the main goroutine
	numWorkers := int(math.Max(1.0, float64(runtime.NumCPU()-1)))

	// prepare channels for communicating parsed data and termination
	batches, users, done := make(chan [][]string, numWorkers), make(chan *pb.UserBatchReq, numWorkers), make(chan int, numWorkers)

	// start the number of workers determined by numWorkers
	fmt.Printf("Starting %v workers...\n", numWorkers)
	for i := 0; i < numWorkers; i++ {
		go prepareBatch(ctx, i, batches, users, done)
	}

	go u.readCSV(ctx, file, &errRet, batches)

	waits := numWorkers

	for {
		select {
		case userBatch, ok := <-users:
			if !ok {
				return
			}
			// send to CRM Integrator via gRPC
			_, err := u.client.ProcessUserBatch(ctx, userBatch)
			if err != nil {
				log.Printf("error sending batch to crmIntegrator %v\n", err)
				errRet.R.AddError(err)
			} else {
				log.Printf("sent batch from %d to %d crmIntegrator\n",
					userBatch.Users[0].Id, userBatch.Users[len(userBatch.Users)-1].Id)
			}

		case <-done:
			waits--
			if waits == 0 {
				close(users)
			}
		// expect cancel
		case <-ctx.Done():
			errRet.Message = "cancelled by caller"
			return
		}
	}
}

func (u UserProcessor) readCSV(ctx context.Context, file io.Reader, errRet *apierror.ErrorMessage, batches chan [][]string) {
	r := csv.NewReader(file)
	// skip the header
	if _, err := r.Read(); err != nil {
		log.Println(err)
		errRet.R.AddError(err)
		return
	}
	defer close(batches)

	batch := make([][]string, 0)

	for {
		// expect cancel
		select {
		case <-ctx.Done():
			return
		default:
			// read one row from .csv file
			row, err := r.Read()
			if err != nil {
				if err != io.EOF {
					log.Println(err)
					errRet.R.AddError(err)
				}
				return
			}

			batch = append(batch, row)
			// if batch is not full
			if len(batch) < int(u.batchSize) {
				// keep adding to it
				continue
			}
			// otherwise, send the batch for further processing
			batches <- batch
			batch = nil
		}
	}
}

func prepareBatch(ctx context.Context, id int, batches <-chan [][]string, usersReq chan<- *pb.UserBatchReq, done chan<- int) {
	proc := 0 // how many batches did we process?

	for batch := range batches {
		users := make([]*pb.User, 0)
		for _, val := range batch {
			// expect cancel
			select {
			case <-ctx.Done():
				return
			default:

				if len(val) < 5 {
					continue
				}

				id, err := strconv.ParseInt(val[0], 10, 64)
				if err != nil {
					log.Println(err)
				}

				user := &pb.User{
					Id:        id,
					FirstName: escape(val[1]),
					LastName:  escape(val[2]),
					Email:     escape(val[3]),
					Phone:     escape(val[4]),
				}
				users = append(users, user)
			}
		}
		usersReq <- &pb.UserBatchReq{
			Users: users,
		}

		proc++

	}
	log.Printf("worker %d finished after processing %d batches\n", id, proc)
	done <- id // send goroutine identifier to done channel
}

func escape(str string) string {
	return strings.Replace(str, "'", "''", -1)
}
