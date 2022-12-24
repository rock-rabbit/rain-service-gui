package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

type RainService struct {
	Host string
}

type RPCSend struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	ID     interface{}   `json:"id"`
}

// Row 下载任务
type Row struct {
	UUID   string
	Status string

	URI     string
	Outdir  string
	Outname string
	Header  string

	Stat struct {
		// Status 状态
		Status int
		// TotalLength 文件总大小
		TotalLength int64
		// CompletedLength 已下载的文件大小
		CompletedLength int64
		// DownloadSpeed 每秒下载字节数
		DownloadSpeed int64
		// EstimatedTime 预计下载完成还需要的时间
		EstimatedTime time.Duration
		// Progress 下载进度, 长度为 100
		Progress int
		// Outpath 文件输出路径
		Outpath string
	}

	Error    string
	Progress int

	CreateTime time.Time
	UpdateTime time.Time

	SyncTime time.Time
}

type Uri struct {
	Uri     string
	Outdir  string
	Outname string
}

func NewRainService(host string) *RainService {
	return &RainService{Host: host}
}

func (rs *RainService) GetRow(status string) ([]Row, error) {
	var v struct {
		Error  string `json:"error"`
		Result []Row  `json:"result"`
		ID     int    `json:"id"`
	}
	err := rs.Send("Service.GetRows", &v, struct{ Status string }{Status: status})
	if err != nil {
		return nil, err
	}
	if v.Error != "" {
		return nil, errors.New(v.Error)
	}
	return v.Result, nil
}

func (rs *RainService) AddUri(u *Uri) (*Row, error) {
	var v struct {
		Error  string `json:"error"`
		Result Row    `json:"result"`
		ID     int    `json:"id"`
	}
	err := rs.Send("Service.AddUri", &v, u)
	if err != nil {
		return nil, err
	}
	if v.Error != "" {
		return nil, errors.New(v.Error)
	}
	return &v.Result, nil
}

func (rs *RainService) Start(uuid string) (*Row, error) {
	var v struct {
		Error  string `json:"error"`
		Result Row    `json:"result"`
		ID     int    `json:"id"`
	}
	err := rs.Send("Service.Start", &v, uuid)
	if err != nil {
		return nil, err
	}
	if v.Error != "" {
		return nil, errors.New(v.Error)
	}
	return &v.Result, nil
}

func (rs *RainService) Delete(uuid string) (*Row, error) {
	var v struct {
		Error  string `json:"error"`
		Result Row    `json:"result"`
		ID     int    `json:"id"`
	}
	err := rs.Send("Service.Delete", &v, uuid)
	if err != nil {
		return nil, err
	}
	if v.Error != "" {
		return nil, errors.New(v.Error)
	}
	return &v.Result, nil
}

func (rs *RainService) Pause(uuid string) (*Row, error) {
	var v struct {
		Error  string `json:"error"`
		Result Row    `json:"result"`
		ID     int    `json:"id"`
	}
	err := rs.Send("Service.Pause", &v, uuid)
	if err != nil {
		return nil, err
	}
	if v.Error != "" {
		return nil, errors.New(v.Error)
	}
	return &v.Result, nil
}

func (rs *RainService) Send(method string, v any, params ...interface{}) error {
	data, err := json.Marshal(&RPCSend{
		Method: method,
		Params: params,
		ID:     0,
	})
	if err != nil {
		return err
	}
	res, err := http.Post(rs.Host, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resData, v)
	if err != nil {
		return err
	}
	return nil
}
