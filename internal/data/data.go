package data

import (
	"context"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"slacker/internal/conf"
	"slacker/internal/data/ent"
	"slacker/internal/pkg/util/errutil"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	NewRecordRepo,
)

// Data .
type Data struct {
	DBClient *ent.Client
	WXClient *miniProgram.MiniProgram
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	// 数据库
	dbClient, err := newDBClient(c)
	if err != nil {
		return nil, nil, err
	}

	// 微信小程序
	wxClient, err := newWXClient(c)
	if err != nil {
		return nil, nil, err
	}

	return &Data{
		WXClient: wxClient,
		DBClient: dbClient,
	}, cleanup, nil
}

func newDBClient(c *conf.Data) (*ent.Client, error) {
	client, err := ent.Open(c.Database.GetDriver(), c.Database.GetSource())
	if err != nil {
		return nil, errutil.WithStack(err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.TODO()); err != nil {
		return nil, errutil.Wrap(err, "creating schema resources failed")
	}

	return client, nil
}

func newWXClient(c *conf.Data) (*miniProgram.MiniProgram, error) {
	client, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:  c.Wx.GetId(),     // 小程序app id
		Secret: c.Wx.GetSecret(), // 小程序app secret
	})
	if err != nil {
		return nil, errutil.Wrap(err, "new wx mini program client failed")
	}

	return client, nil
}
