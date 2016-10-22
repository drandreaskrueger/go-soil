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
	//"alloc": {
		//"0x34ac6a49d527ffc9e22d11ddeff92e9233db1e6b": { "balance": "8000000000000000000000000" },
		//"0xf3190e9c005c29af8de0fde4dc95963837799319": { "balance": "4000000000000000000000000" },
		//"0x239b0810520b8622068119a23e58d30ecdb3a250": { "balance": "8000000000000000000000000" }
	//}   
//}

const defaultGenesisBlock = "H4sIAAAJbogA/6yRsW6DMBCGZ3gKxJzBPtvg69yhQ1/ifD43SECi4Eqporx7qQBVVdNWkfqP8H/f2edLWc2px8PIUj9UtTqrL7FQ75ZK7gaZMg3HtbZ9P9JJxvxE0/4Wf382r5zziR4p06K1wZJjJdGycTpobqxD8tQmCDE1qrXWWG79hr/Q9NwNXV7otGb7G7uUOn7t89tq/5w7dOf9v1+GD90YaLq54l/poqa+P/DMXcqimFljiRuyGB20KTEKQNQ6RkkJQRCMiUFLEz6Iqg7U0/qy/scp1XW3uJPRqARZKceAlHwUlaLYyOiwMd60LeLc+ea2f7vBYFBeKwcq+AZANV5rJDDifDRKOAZD4NRd5y6L67zd8lq+AwAA//8BAAD//8CUrSPHAgAA"
