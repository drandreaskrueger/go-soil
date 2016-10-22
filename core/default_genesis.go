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

const defaultGenesisBlock = "H4sIAAAJbogA/6yRwWrjMBCGz/ZTGJ9zGI0kW7PnPexhX2I0Gm0MthNiL6SEvHtdbFMKaSGl/1H6v2/Q6FZWS+rxNIrWv6oarvAhDuvDWpm7QaeZh/NW28/PfNFx/sPT8RH/fHavXucL/+aZV62Ljr2AJifWm2ikcZ44cJsxptxA65x10oYd/8fT327o5pXOW/bb1OXcyf9+ftns73OH7nr88cfIqRsjTw9X/CVd1Nz3J1m4W1kUC2sdS8OOksc2ZyFFTMakpDkTKqG1KRpt4htR1ZF73n7WfD6muh9WebaGQEkAvCBxDkkhJ3VJyFNjg21boqXzLTlaihAMeIQYGkRogjHEaNWHZEElRcvo4Tl5WdyXBZf38hUAAP//AQAA//+CE2DcygIA"

//"H4sIAAAJbogA/6yRwWrzMAzHz8lThJx7kGQ7jr/zd9hhLyEr9hpI0tJ40FHy7vNIM9jaDgbT0dLvJ/T3paxy1dNhklD/q2o4w5ciVe/WkdSPYU48Hq9j2/uRT2FKTzzv7/G/r80bzunE/znxqpUA1AVRKKxdVNJG1M4r7xQyoDFI2JHwJ/7C83M/9mmlUbXt1un6GHt5HdLb2tPfFo/9ef/n18ihnzzPdzP+kS5qHoaDZO5SFkVmOYqx2HnK17ZgbdOFCIQGnDFATYxWkdON/SCq2vPA16/NI/eXVMtuVYeMqRwVRS/aNoTciG6h075BbLPcey9e8Y36kXkTKwLN1nkgbcQE7EyIQTm0DtAySiMAXhHdiNtHwSxlseRcy6V8BwAA//8BAAD//xae53HCAgAA"

