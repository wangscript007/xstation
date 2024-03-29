package device

import (
	"errors"
	"xstation/pkg/ctx"

	"github.com/gin-gonic/gin"
	"github.com/wlgd/xproto"
)

type devInfo struct {
	deviceId string
}

func checkParam(c *gin.Context, param interface{}) (*devInfo, error) {
	dvr := &devInfo{}
	dvr.deviceId = c.Query("deviceId")
	if dvr.deviceId == "" {
		return nil, errors.New("deviceId invalid")
	}
	if param == nil {
		return dvr, nil
	}
	return dvr, c.ShouldBind(&param)
}

// DvrQuery 查询
func QueryHandler(c *gin.Context) {
	var param xproto.Query
	i, err := checkParam(c, &param)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var resp []xproto.QueryResult
	if err := xproto.SyncSendToDevice(xproto.REQ_Query, param, &resp, i.deviceId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// DvrParameters 参数管理 获取/设置
func ParametersHandler(c *gin.Context) {
	var param interface{}
	i, err := checkParam(c, &param)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var resp interface{}
	if err := xproto.SyncSendToDevice(xproto.REQ_Parameters, param, &resp, i.deviceId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// DvrControl 设置控制
func ControlPTZHandler(c *gin.Context) {
	var param xproto.Control
	i, err := checkParam(c, &param.Data)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	param.Type = xproto.CTRL_PTZ
	var resp interface{}
	if err := xproto.SyncSendToDevice(xproto.REQ_Control, param, &resp, i.deviceId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// ControlRebootHandler 设置控制
func ControlRebootHandler(c *gin.Context) {
	i, err := checkParam(c, nil)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	param := &xproto.Control{Type: xproto.CTRL_Reboot}
	var resp interface{}
	if err := xproto.SyncSendToDevice(xproto.REQ_Control, param, &resp, i.deviceId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// ControlCaptureHandler 设置控制
func ControlCaptureHandler(c *gin.Context) {
	var param xproto.Control
	i, err := checkParam(c, &param.Data)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	param.Type = xproto.CTRL_Capture
	var resp interface{}
	if err := xproto.SyncSendToDevice(xproto.REQ_Control, param, &resp, i.deviceId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// ControlOsdHandler 设置控制
func ControlOsdHandler(c *gin.Context) {
	var param xproto.Control
	i, err := checkParam(c, &param.Data)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	param.Type = xproto.CTRL_OsdSpeed
	var resp interface{}
	if err := xproto.SyncSendToDevice(xproto.REQ_Control, param, &resp, i.deviceId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// DvrLiveStream 实时视频请求
// return url 直接打开播放
func LiveStreamHandler(c *gin.Context) {
	var param xproto.LiveStream
	i, err := checkParam(c, &param)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var resp interface{}
	if err := xproto.SyncSendToDevice(xproto.REQ_LiveStream, param, &resp, i.deviceId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// VoiceHandler 语音业务
// return url 直接打开播放
func VoiceHandler(c *gin.Context) {
	var param xproto.Voice
	i, err := checkParam(c, &param)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var resp interface{}
	if err := xproto.SyncSendToDevice(xproto.REQ_Voice, param, &resp, i.deviceId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// DvrPlayback 录像回放
func PlaybackHandler(c *gin.Context) {
	var param xproto.Playback
	i, err := checkParam(c, &param)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var resp interface{}
	if err := xproto.SyncSendToDevice(xproto.REQ_Playback, param, &resp, i.deviceId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// DvrSerialTransparent 串口透传设置
func SerialTransparentHandler(c *gin.Context) {
	var param xproto.SerialTransparent
	i, err := checkParam(c, &param)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var resp interface{}
	if err := xproto.SyncSendToDevice(xproto.REQ_SerialTransparent, param, &resp, i.deviceId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// DvrSerialTransfer 串口透传数据
func SerialTransferHandler(c *gin.Context) {
	session := c.Query("session")
	if session == "" {
		ctx.JSONWriteError(xproto.ErrParam, c)
		return
	}
	var param xproto.RawData
	i, err := checkParam(c, &param)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var resp interface{}
	if err := xproto.SyncSendToDevice(xproto.REQ_SerialTransfer, param, &resp, i.deviceId, session); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// DvrFileTransfer 文件传输
func FileTransferHandler(c *gin.Context) {
	var param xproto.FileTransfer
	i, err := checkParam(c, &param)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var resp interface{}
	if err := xproto.SyncSendToDevice(xproto.REQ_FileTransfer, param, &resp, i.deviceId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// DvrFtpTransfer ftp文件传输
func FtpTransferHandler(c *gin.Context) {
	var param xproto.FtpTransfer
	i, err := checkParam(c, &param)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var resp interface{}
	if err := xproto.SyncSendToDevice(xproto.REQ_FtpTransfer, param, &resp, i.deviceId); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(gin.H{"data": resp}, c)
}

// DvrCloseLink 关闭链路
func CloseLinkHandler(c *gin.Context) {
	i, err := checkParam(c, nil)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	sessionID := c.Query("session")
	if err := xproto.SyncSendToDevice(xproto.REQ_Close, nil, nil, i.deviceId, sessionID); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}
