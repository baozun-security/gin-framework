package setting

// server run mode
const (
	Development RunMode = "dev"
	Production  RunMode = "prd"
)

// RunMode defines app run mode
type RunMode string

// IsValid return tru if mode is valid
func (mode RunMode) IsValid() bool {
	switch mode {
	case Development, Production:
		return true
	}
	return false
}

// IsDevelopment returns true if mode is development
func (mode RunMode) IsDevelopment() bool {
	return mode == Development
}

// IsProduction returns true if mode equals to production
func (mode RunMode) IsProduction() bool {
	return mode == Production
}

func (mode RunMode) String() string {
	if mode.IsDevelopment() {
		return "development"
	}
	if mode.IsProduction() {
		return "production"
	}
	return string(mode)
}
