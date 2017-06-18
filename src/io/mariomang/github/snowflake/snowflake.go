package snowflake

import (
	"fmt"
	"io/mariomang/github/consts"
	"sync"
	"time"
)

type SnowFlake struct {
	baseTimeStamp int64
	workIDBits    int64
	machineIdbits int64
	sequenceBits  int64
	maxWorkID     int64
	maxMachineID  int64
	maxSequence   int64
	TimeStamp     int64
	WorkID        int64
	MachineID     int64
	Sequence      int64
	lastTimeStamp int64
	lock          sync.Mutex
}

type SnowFlakePool struct {
	snowFlakes map[int64]*SnowFlake
}

var pool *SnowFlakePool

func init() {
	pool = new(SnowFlakePool)
	pool.snowFlakes = make(map[int64]*SnowFlake)
}

func NewSnowFlake(workID int64, machineId int64) *SnowFlake {

	if s, ok := pool.snowFlakes[workID]; ok {
		return s
	}
	snowflake := &SnowFlake{
		baseTimeStamp: consts.BaseTime,
		workIDBits:    consts.WorkIDBits,
		machineIdbits: consts.MachineIDBits,
		sequenceBits:  consts.SequenceBits,
		lastTimeStamp: -1,
		WorkID:        workID,
		MachineID:     machineId,
	}
	snowflake.maxWorkID = -1 ^ (-1 << uint(snowflake.workIDBits))
	snowflake.maxMachineID = -1 ^ (-1 << uint(snowflake.machineIdbits))
	snowflake.maxSequence = -1 ^ (-1 << uint(snowflake.sequenceBits))

	if workID > snowflake.maxWorkID || machineId > snowflake.maxMachineID {
		return nil
	}

	pool.snowFlakes[workID] = snowflake

	return snowflake
}

func (s *SnowFlake) GetID() int64 {
	s.lock.Lock()
	s.TimeStamp = time.Now().Unix()
	if s.TimeStamp < s.lastTimeStamp {
		fmt.Println("Clock moved backwards.")
		return -1
	}

	if s.TimeStamp == s.lastTimeStamp {
		s.Sequence++
		if s.Sequence > s.maxSequence {
			s.TimeStamp = gotoNextMills(s.lastTimeStamp)
		}
	} else {
		s.Sequence = 0
	}

	s.lastTimeStamp = s.TimeStamp

	timestamp := (s.TimeStamp - s.baseTimeStamp) << uint(s.workIDBits+s.machineIdbits+s.sequenceBits)
	workdID := s.WorkID << uint(s.machineIdbits+s.sequenceBits)
	machineID := s.MachineID << uint(s.sequenceBits)

	id := timestamp + workdID + machineID + s.Sequence
	s.lock.Unlock()
	fmt.Printf("timestamp: %d workId: %d machineId: %d SequenceID: %d ID: %b\n ", s.TimeStamp, s.WorkID, s.MachineID, s.Sequence, id)
	return id
}

func gotoNextMills(lastTimeStamp int64) int64 {
	timestamp := time.Now().Unix()
	for timestamp <= lastTimeStamp {
		timestamp = time.Now().Unix()
	}
	return timestamp
}
