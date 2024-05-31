package utils

import (
	"strconv"
	"time"
)

func ConvertEpochToHuman(convertEpoch string) string {
	epoch, err := strconv.ParseInt(convertEpoch, 10, 64)
	if err != nil {
		return "Erro ao converter epoch"
	}
	timeHuman := time.Unix(epoch, 0)
	return timeHuman.Format("02/01/2006 15:04:05")
}
