// MIT License
//
// Copyright (c) 2016-2018 GACHAIN
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package consts

import (
	"time"
)

// VERSION is current version
const VERSION = "0.1.6b13"

// NETWORK_ID is id of network
const NETWORK_ID = 1

// DEFAULT_TCP_PORT used when port number missed in host addr
const DEFAULT_TCP_PORT = 7078

// AddressLength is length of address
const AddressLength = 20

// PubkeySizeLength is pubkey length
const PubkeySizeLength = 64

// BlockSize is size of block
const BlockSize = 16

// HashSize is size of hash
const HashSize = 32


const (
	TxTypeFirstBlock  = 1
	TxTypeStopNetwork = 2

	TxTypeParserFirstBlock  = "FirstBlock"
	TxTypeParserStopNetwork = "StopNetwork"
)

// TxTypes is the list of the embedded transactions
var TxTypes = map[int]string{
	TxTypeFirstBlock:  TxTypeParserFirstBlock,
	TxTypeStopNetwork: TxTypeParserStopNetwork,
}

// ApiPath is the beginning of the api url
var ApiPath = `/api/v2/`

// DefaultConfigFile name of config file (toml format)
const DefaultConfigFile = "config.toml"

// DefaultWorkdirName name of working directory
const DefaultWorkdirName = "gachain-data"

// DefaultPidFilename is default filename of pid file
const DefaultPidFilename = "go-gachain.pid"

// DefaultLockFilename is default filename of lock file
const DefaultLockFilename = "go-gachain.lock"

// FirstBlockFilename name of first block binary file
const FirstBlockFilename = "1block"

// PrivateKeyFilename name of wallet private key file
const PrivateKeyFilename = "PrivateKey"

// PublicKeyFilename name of wallet public key file
const PublicKeyFilename = "PublicKey"

// NodePrivateKeyFilename name of node private key file
const NodePrivateKeyFilename = "NodePrivateKey"

// NodePublicKeyFilename name of node public key file
const NodePublicKeyFilename = "NodePublicKey"

// KeyIDFilename generated KeyID
const KeyIDFilename = "KeyID"

// RollbackResultFilename rollback result file
const RollbackResultFilename = "rollback_result"

// FromToPerDayLimit day limit token transfer between accounts
const FromToPerDayLimit = 10000

// TokenMovementQtyPerBlockLimit block limit token transfer
const TokenMovementQtyPerBlockLimit = 100

// TCPConnTimeout timeout of tcp connection
const TCPConnTimeout = 5 * time.Second

// TxRequestExpire is expiration time for request of transaction
const TxRequestExpire = 1 * time.Minute

// DefaultTempDirName is default name of temporary directory
const DefaultTempDirName = "gachain-temp"
