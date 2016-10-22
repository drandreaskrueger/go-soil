// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"compress/gzip"
	"encoding/base64"
	"io"
	"strings"
)

func NewDefaultGenesisReader() (io.Reader, error) {
	return gzip.NewReader(base64.NewDecoder(base64.StdEncoding, strings.NewReader(defaultGenesisBlock)))
}

// Testnet Genesis

//{
    //"nonce": "0x0000000000000042",
    //"timestamp": "0x0",
    //"parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    //"extraData": "0x4b4a5c0ed4c351b1c6459a8a7f2bdf6074434c78",
    //"gasLimit": "0xffffffff",
    //"difficulty": "0x400",
    //"mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    //"coinbase": "0x0000000000000000000000000000000000000000",
	//"alloc": {}   
//}

const defaultGenesisBlock = "H4sIAAAJbogA/6yRTWvEIBCGz8mvEM97UDNq7LmHHvonZvzoCvlYNhZSlvz3piShFLaFhb5HnecZfL3VbA0fxsFH/sS4mMWPgOKnbaTkPk4F+8s+dpxf8BqH8oLT+R7/eA5vnMsVn7HgpgUC1F7EAL7RkqQ3oB22aJOikIywAA142x74G06vuc9lo9Oe4zbklLJ/78rHbv/e2+f5/O+P8WMeCKe7Ff9JVxy7bvQrd6uramVtAtuQJpIYk44OE1mhrHONExS0EDI5oX34Ihgn7HD/Wfn7GracNnljFARjoXUalAEnUotq7TmGgEKTdEYbDOEh+VJXC2P1Un8CAAD//wEAAP//LgHeTmsC"
