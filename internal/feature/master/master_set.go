package master

import (
	"lapakita-backend/internal/feature/master/area"
	businesstype "lapakita-backend/internal/feature/master/business_type"
	"lapakita-backend/internal/feature/master/public"

	"github.com/google/wire"
)

var MasterFeatureSet = wire.NewSet(
	area.AreaFeatureSet,
	businesstype.BusinessTypeFeatureSet,
	public.PublicFeatureSet,
)
