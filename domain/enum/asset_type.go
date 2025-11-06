package enum

type AssetType string

const (
	AssetTypeIT     AssetType = "it"
	AssetTypeNonIT  AssetType = "non_it"
)

func (t AssetType) IsValid() bool {
	switch t {
	case AssetTypeIT, AssetTypeNonIT:
		return true
	default:
		return false
	}
}