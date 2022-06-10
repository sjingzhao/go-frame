package crontabx

type Crontab struct {
    name string
    spec string
    fun  func()
}

type CrontabServer struct {
    crontabs []*Crontab
}

func NewCrontabServer() *CrontabServer {
    return &CrontabServer{
        crontabs: make([]*Crontab, 0),
    }
}
