package utils

import (
	"log"
	"testing"
)

func TestAverageTool(t *testing.T) {
	result := CalculateAverage("../resource_allocation/withCache_mem_1record_totalAllocation.txt")
	if result == 0 {
		t.Fail()
	}

	log.Println(result)

	result = CalculateAverage("../resource_allocation/withCache_mem_1record_totalHeapSystem.txt")
	if result == 0 {
		t.Fail()
	}

	log.Println(result)
}
