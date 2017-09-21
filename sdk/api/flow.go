package api

/**
 * @author dzh
 * @date 20/09/2017 20:49
 * @since 0.0.1
 */

type FlowApi interface {
    YunpianApi
}

type FlowApiImpl struct {
    YunpianApiOption
}

func NewFlow() FlowApi {
    flow := &FlowApiImpl{}
    flow.name = Flow
    return flow
}
