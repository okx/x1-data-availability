package rpcsyncer

import (
	"context"
	"github.com/ethereum/go-ethereum/crypto"
	"time"

	"github.com/0xPolygon/cdk-data-availability/db"
	"github.com/0xPolygon/cdk-data-availability/log"
	"github.com/0xPolygon/cdk-data-availability/types"
)

type RPCSyncer struct {
	l2RPCUrl     string
	maxBatchSize uint64
	intervalTime time.Duration
	db           db.DB

	stop chan struct{}
}

func NewRPCSyncer(l2RPCUrl string, maxBatchSize uint64, intervalTime time.Duration, db db.DB) *RPCSyncer {
	rpcSyncer := &RPCSyncer{
		l2RPCUrl:     l2RPCUrl,
		maxBatchSize: maxBatchSize,
		intervalTime: intervalTime,
		db:           db,
		stop:         make(chan struct{}),
	}

	return rpcSyncer
}

// Start starts the SequencerTracker
func (syncer *RPCSyncer) Start(ctx context.Context) {
	log.Infof("starting rpc syncer")

	start, _ := getStartBlock(syncer.db)

	for {
		select {
		case <-ctx.Done():
			if ctx.Err() != nil && ctx.Err() != context.DeadlineExceeded {
				log.Warnf("context cancelled: %v", ctx.Err())
			}
		default:
			time.Sleep(syncer.intervalTime)
			l2MaxBatch, err := BatchNumber(syncer.l2RPCUrl)
			if err != nil {
				log.Fatal("error getting max batch: %v", err)
			}
			log.Infof("Starting from batch %v, max batch %v", start, l2MaxBatch)
			if start > l2MaxBatch {
				log.Infof("Sleep 30 seconds")
				time.Sleep(30 * time.Second)
			}

			to := start + syncer.maxBatchSize - 1
			if to > l2MaxBatch && to > syncer.maxBatchSize {
				to = l2MaxBatch
				start = to - syncer.maxBatchSize + 1
			}
			log.Infof("Calling ..")
			seqBatches, err := BatchesByNumbers(syncer.l2RPCUrl, start, to)
			if err != nil {
				log.Fatal("error getting batch data: %v", err)
			}

			log.Infof("Hashing ..")
			offChainData := []types.OffChainData{}
			for _, seqBatch := range seqBatches {
				key := crypto.Keccak256Hash(seqBatch.BatchL2Data)

				offChainData = append(offChainData, types.OffChainData{
					Key:   key,
					Value: seqBatch.BatchL2Data,
				})
			}
			log.Infof("Storing ..")
			err = setStoreOffChainData(syncer.db, offChainData)
			if err != nil {
				log.Fatal("error storing off chain data: %v", err)
			}
			log.Infof("Stored data for batchs [%v,%v], size:%v", start, to, len(offChainData))
			if setStartBlock(syncer.db, to) != nil {
				log.Fatal("error setting start block: %v", err)
			}
			start = to + 1
		}
	}
}

// Stop stops the SequencerTracker
func (st *RPCSyncer) Stop() {
	close(st.stop)
}
