package enum

type AssetStatus string

const (
	AssetStatusAvailable AssetStatus = "available"
	AssetStatusBooked    AssetStatus = "booked"
	AssetStatusBroken    AssetStatus = "broken"
	AssetStatusRepair    AssetStatus = "repair"
)

func (s AssetStatus) IsValid() bool {
	switch s {
	case AssetStatusAvailable, AssetStatusBooked, AssetStatusBroken, AssetStatusRepair:
		return true
	default:
		return false
	}
}