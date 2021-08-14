package options

import (
	cliflag "github.com/daz2yy/go-base/pkg/cli/flag"
	"github.com/daz2yy/go-base/pkg/log"

	genericoptions "github.com/daz2yy/go-base/internal/pkg/options"
)

// Options runs a api server.
type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions `json:"server" mapstructure:"server"`

	InsecureServing *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServing   *genericoptions.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	MySQLOptions    *genericoptions.MySQLOptions           `json:"mysql"    mapstructure:"mysql"`
	Log             *log.Options                           `json:"log"      mapstructure:"log"`
	FeatureOptions  *genericoptions.FeatureOptions         `json:"feature"  mapstructure:"feature"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	o := Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		InsecureServing:         genericoptions.NewInsecureServingOptions(),
		SecureServing:           genericoptions.NewSecureServingOptions(),
		MySQLOptions:            genericoptions.NewMySQLOptions(),
		Log:                     log.NewOptions(),
		FeatureOptions:          genericoptions.NewFeatureOptions(),
	}

	return &o
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))
	o.InsecureServing.AddFlags(fss.FlagSet("insecure"))
	o.SecureServing.AddFlags(fss.FlagSet("secure"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))
	o.Log.AddFlags(fss.FlagSet("logs"))
	o.FeatureOptions.AddFlags(fss.FlagSet("feature"))

	return fss
}

// Validate checks Options and return a slice of found errs.
func (o *Options) Validate() []error {
	var errs []error

	errs = append(errs, o.GenericServerRunOptions.Validate()...)
	errs = append(errs, o.InsecureServing.Validate()...)
	errs = append(errs, o.SecureServing.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)
	errs = append(errs, o.Log.Validate()...)
	errs = append(errs, o.FeatureOptions.Validate()...)

	return errs
}
