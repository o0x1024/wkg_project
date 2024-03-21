package global

import "wkg-agent/model/request"

var (
	AttackQueue = make(chan request.UploadReq, 50)
)
