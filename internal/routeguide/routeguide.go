package routeguide

import (
	"context"

	"github.com/smantic/cannonical/proto"
)

type RouteGuide struct {
	proto.UnimplementedRouteGuideServer
}

func (r *RouteGuide) GetFeature(context.Context, *proto.Point) (*proto.Feature, error) {
	return new(proto.Feature), nil
}
