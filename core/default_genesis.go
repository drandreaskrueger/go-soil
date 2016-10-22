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

const defaultGenesisBlock = "H4sIAAAJbogA/6yRzWrrMBBG1/ZTGK2zkOQZybrru+iiLzGjn0ZgOyFWISX43etim1KaFAr9ltKcM5rRrW6WiPE0+ij+NUJe5ZeAFoe1pOQhToWG81a2n5/pEsfyRNPxHv/77N54LRf6T4VWLTAQehkD+BYVK28AHXVkk+aQjLQALXjb7fgLTc95yGWl05b9NuSUsn/ty9tm/+w75Ovxz4fxpzwyTXdX/CNdCer7k1+4W11VC2sT2JaRWVFMGB0ltlJb51onOaCUKjmJPnwQjWDqaftZhQ+aNPNhVbdGQzAWOoegDTiZOtLLlmMIJJGVM2gohO/q7tH757qal/HruX4HAAD//wEAAP///YpnBmgC"
