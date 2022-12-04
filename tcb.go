package tcb

import (
	"github.com/dongzeli95/tcb-openapi-go/component"
	"github.com/dongzeli95/tcb-openapi-go/component/database"
	"github.com/dongzeli95/tcb-openapi-go/component/functions"
	"github.com/dongzeli95/tcb-openapi-go/component/storage"
	"github.com/dongzeli95/tcb-openapi-go/config"
	"github.com/dongzeli95/tcb-openapi-go/context"
	"github.com/dongzeli95/tcb-openapi-go/sts"
	"github.com/dongzeli95/tcb-openapi-go/util"
	"github.com/gogf/gf/database/gredis"
	"github.com/sirupsen/logrus"
	"os"
)

/*
Tcb 实例
*/
type Tcb struct {
	context *context.Context
	core    *component.Core
}

/*
创建实例
*/
func NewTcb(config *config.Config) *Tcb {
	//上下文
	ctx := &context.Context{
		Config: config,
		Logger: &logrus.Logger{
			Out:          os.Stdout,
			Formatter:    &util.CustomerFormatter{Prefix: config.LogPrefix},
			Level:        logrus.DebugLevel,
			ExitFunc:     os.Exit,
			ReportCaller: true,
		},
	}
	//cam
	client := sts.NewStsClient(&config.StsConfig, gredis.New(&config.RedisConfig), ctx.Logger)
	return &Tcb{ctx, component.NewCore(ctx, client)}
}

//接入数据库
func (t *Tcb) GetDatabase() *database.Database {
	return database.NewDatabase(t.context, t.core)
}

//接入云函数
func (t *Tcb) GetFunction() *functions.Function {
	return functions.NewFunction(t.context, t.core)
}

//接入云存储
func (t *Tcb) GetStorage() *storage.Storage {
	return storage.NewStorage(t.context, t.core)
}
