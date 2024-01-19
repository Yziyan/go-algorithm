// @Author: Ciusyan 1/16/24

package reflect_module

import (
	"context"
	"sync"
)

type ValueITEPostRank struct {
	NickName           string  `default_value:"ITE计算"`
	Description        string  `default_value:"流量价值场景，计算ite作为物料打分,ite=ctr*(ctcvr-cvr)"`
	WeightFeatureName  string  // 加权权重特征名
	WeightDefaultValue float64 `default_value:"1"` // 权重默认值
}

func (s *ValueITEPostRank) DoAction(ctx context.Context, sc *StrategyContext, ctr map[string]map[string]*ItemPredictResp,
	cvr map[string]map[string]*ItemPredictResp, ctcvr map[string]map[string]string) (map[string]map[string]*ItemPredictResp, bool) {

	return nil, false
}

type StrategyContext struct {
	ModuleResultMap sync.Map          //module返回数据，存储的是reflect.value(即是地址)
	TimeMap         sync.Map          //module的耗时
	SkipModuleMap   sync.Map          //记录导致后续worfklow不执行的module
	CountMap        map[string]*int32 //引擎执行时，统计当前执行的gorutine
	ErrMap          sync.Map          //引擎执行时，记录当前执行module过程中的panic
	ContextMap      sync.Map          //引擎执行时，传递上下文信息
}

type ItemPredictResp struct {
	ID           string
	Score        float64
	ModelName    string
	ModelVersion string
	Feature      map[string]string
}

type BusinessResultModule struct {
	NickName    string `default_value:"结果组装模块"`
	Description string `default_value:"流量价值场景,对之前流程的结果进行组装，并返回结果"`
}

func (s *BusinessResultModule) DoAction(ctx context.Context, sc *StrategyContext,
	strategy map[string]map[string]*ItemPredictResp) (*RespAndFeature, bool) {

	return nil, false
}

type RespAndFeature struct {
	Response interface{}
	Feature  map[string]map[string]map[string]string
	OtherMap map[string]map[string]string
}
