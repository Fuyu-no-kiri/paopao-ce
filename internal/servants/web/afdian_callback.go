package web

import (
	"fmt"
	"github.com/alimy/mir/v4"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
)

var (
	_ api.AfdianCallback = (*afdianCallbackSrv)(nil)
)

type afdianCallbackSrv struct {
	api.UnimplementedAfdianCallbackServant
	*base.DaoServant
}

func (srv *afdianCallbackSrv) Callback(req *web.AfdianCallbackReq) (*web.AfdianCallbackResp, mir.Error) {
	fmt.Println(req.Data)

	remark := req.Data.Order.Remark
	if remark == "" {
		return &web.AfdianCallbackResp{
			EC: 502,
			EM: "未填写用户UID, 请联系管理员解决!",
		}, nil
	}

	user, err := srv.Ds.GetUserByUsername(remark)
	amount := req.Data.Order.TotalAmount
	recharge, err := srv.Ds.CreateRecharge(user.ID, int64(amount))
	if err != nil {
		return &web.AfdianCallbackResp{
			EC: 502,
			EM: "充值失败, 请联系管理员!",
		}, nil
	}
	err = srv.Ds.HandleRechargeSuccess(recharge, "0")
	if err != nil {
		return &web.AfdianCallbackResp{
			EC: 502,
			EM: "充值失败, 请联系管理员!",
		}, nil
	}

	return &web.AfdianCallbackResp{EC: 200, EM: ""}, nil
}

func newAfdianSrv(s *base.DaoServant) api.AfdianCallback {
	return &afdianCallbackSrv{
		DaoServant: s,
	}
}
