// Copyright 2019 Graham Clark. All rights reserved.  Use of this source
// code is governed by the MIT license that can be found in the LICENSE
// file.

// +build !darwin,!android,!windows

package termshark

var CopyToClipboard = []string{"xsel", "-i", "-b"}

var OpenURL = []string{"xdg-open"}
