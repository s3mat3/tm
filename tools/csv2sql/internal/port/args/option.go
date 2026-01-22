/*
 option.go
 Copyright © 2026 s3mat3
 This code is licensed under the MIT License, see the LICENSE file for details
 Author s3mat3
*/
// Package args Command line arguments
package args

// Options command line options
type Options struct {
	// In input csv file with path
	In string
	// Out output file with path
	Out string
	// Mode genarate mode
	Mode string
	// First drop
	Drop bool
}

func NewOptions(i *string, o *string, m *string, w *bool) *Options {
	return &Options{
		In:		*i,
		Out:	*o,
		Mode:	*m,
		Drop:	*w,
	}
}


//<-- option.go ends here.
