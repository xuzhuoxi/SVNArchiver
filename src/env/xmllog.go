// Create on 2022/7/15
// @author xuzhuoxi
package env

import (
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"time"
)

type ArchLogItem struct {
	Id       string `xml:"id,attr" json:"id"`     // 归档Id
	Code     string `xml:"code,attr" json:"code"` // 归档文件特征码
	FileName string `xml:"name,attr" json:"name"` // 归档文件名称
	FilePath string `xml:",innerxml" json:"path"` // 归档文件路径
}

func (o ArchLogItem) String() string {
	return fmt.Sprintf("{Id=%s, Code=%s, FileName=%s}", o.Id, o.Code, o.FileName)
}

type ArchLog struct {
	Date string         `xml:"date,attr" json:"date"` // 本次归档执行的时间
	Logs []*ArchLogItem `xml:"arch" json:"arch"`      // 归档处理信息列表
}

func (o ArchLog) String() string {
	return fmt.Sprintf("{Date=%s, Logs[%d]=%v}", o.Date, len(o.Logs), o.Logs)
}

func NewOutLogContext(archXmlLog *ArchXmlLog) *OutLogContext {
	return &OutLogContext{ArchXmlLog: archXmlLog, ArchLog: &ArchLog{Date: time.Now().In(time.Local).Format(LocalOutputLayout)}}
}

type OutLogContext struct {
	ArchXmlLog *ArchXmlLog // 归档信息输出配置
	ArchLog    *ArchLog    `xml:"log" json:"log"` // 归档信息文件的根节点
}

func (c *OutLogContext) AppendLog(id string, code string, filePath string) {
	_, fileName := filex.Split(filePath)
	item := &ArchLogItem{Id: id, Code: code, FileName: fileName, FilePath: filePath}
	c.ArchLog.Logs = append(c.ArchLog.Logs, item)
}
