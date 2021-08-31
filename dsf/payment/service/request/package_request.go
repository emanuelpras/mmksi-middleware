package request

type HeaderPackageRequest struct {
	ApplicationName string `json:"applicationName"`
}

type PackageRequest struct {
	Brand        string `json:"Brand"`
	Model        string `json:"Model"`
	Variant      string `json:"Variant"`
	Province     string `json:"Province"`
	City         string `json:"City"`
	PackageName  string `json:"PackageName"`
	CarCondition string `json:"CarCondition"`
}
