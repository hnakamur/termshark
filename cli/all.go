// Copyright 2019 Graham Clark. All rights reserved.  Use of this source
// code is governed by the MIT license that can be found in the LICENSE
// file.
//

package cli

import "github.com/jessevdk/go-flags"

//======================================================================

// Used to determine if we should run tshark instead e.g. stdout is not a tty
type Tshark struct {
	PassThru    string `long:"pass-thru" default:"auto" optional:"true" optional-value:"true" choice:"yes" choice:"no" choice:"auto" choice:"true" choice:"false" description:"Run tshark instead (auto => if stdout is not a tty)."`
	PrintIfaces bool   `short:"D" optional:"true" optional-value:"true" description:"Print a list of the interfaces on which termshark can capture."`
}

// Termshark's own command line arguments. Used if we don't pass through to tshark.
type Termshark struct {
	Iface         string         `value-name:"<interface>" short:"i" description:"Interface to read."`
	Pcap          flags.Filename `value-name:"<file>" short:"r" description:"Pcap file to read."`
	DecodeAs      []string       `short:"d" description:"Specify dissection of layer type." value-name:"<layer type>==<selector>,<decode-as protocol>"`
	PrintIfaces   bool           `short:"D" optional:"true" optional-value:"true" description:"Print a list of the interfaces on which termshark can capture."`
	DisplayFilter string         `short:"Y" description:"Apply display filter." value-name:"<displaY filter>"`
	CaptureFilter string         `short:"f" description:"Apply capture filter." value-name:"<capture filter>"`
	PlatformSpecific
	PassThru string `long:"pass-thru" default:"auto" optional:"true" optional-value:"true" choice:"auto" choice:"true" choice:"false" description:"Run tshark instead (auto => if stdout is not a tty)."`
	LogTty   bool   `long:"log-tty" optional:"true" optional-value:"true" choice:"true" choice:"false" description:"Log to the terminal."`
	Debug    string `long:"debug" default:"false" hidden:"true" optional:"true" optional-value:"true" choice:"true" choice:"false" description:"Enable termshark debugging. See https://termshark.io/userguide."`
	Help     bool   `long:"help" short:"h" optional:"true" optional-value:"true" description:"Show this help message."`
	Version  []bool `long:"version" short:"v" optional:"true" optional-value:"true" description:"Show version information."`

	Args struct {
		FilterOrFile string `value-name:"<filter-or-file>" description:"Filter (capture for iface, display for pcap), or pcap file to read."`
	} `positional-args:"yes"`
}

// If args are passed through to tshark (e.g. stdout not a tty), then
// strip these out so tshark doesn't fail.
var TermsharkOnly = []string{"--pass-thru", "--log-tty", "--debug"}

func FlagIsTrue(val string) bool {
	return val == "true" || val == "yes"
}

//======================================================================
// Local Variables:
// mode: Go
// fill-column: 78
// End:
