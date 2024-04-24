package services

import (
	"fmt"
	pb "repository_lock/protobufs"
	"sync"
)

type RepositoryLockService struct {
	pb.UnimplementedRepositoryLockServer
}

var cache = make(map[string]map[string][]bool)
var mutex = sync.Mutex{}

func (s *RepositoryLockService) HandleSaveRawSuccess(stream pb.RepositoryLock_HandleSaveRawSuccessServer) error {
	fmt.Println("new stream")

	// lock cache variable memory with mutex
	mutex.Lock()
	defer mutex.Unlock()

	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		// initialize file chunk metadata to redis if it doesn't already exist
		if _, userIdExists := cache[req.UserId]; !userIdExists {
			cache[req.UserId] = make(map[string][]bool)
		}

		if _, filePathExists := cache[req.UserId][req.FilePath]; !filePathExists {
			cache[req.UserId][req.FilePath] = []bool{}
			for i := 0; i < (int)(req.NumTotalChunks); i++ {
				cache[req.UserId][req.FilePath] = append(cache[req.UserId][req.FilePath], false)
			}
		}

		// update file chunk metadata to redis
		cache[req.UserId][req.FilePath][req.ContentChunkIndex] = true

		// determine if this was the last chunk
		allChunksSave := true
		for i := 0; i < len(cache[req.UserId][req.FilePath]); i++ {
			allChunksSave = allChunksSave && cache[req.UserId][req.FilePath][i]
		}

		err = stream.Send(&pb.ReportSaveRawSuccessResponse{FilePath: req.FilePath, AllChunksSaved: allChunksSave})
		if err != nil {
			return err
		}
	}
}
