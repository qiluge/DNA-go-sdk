// SPDX-License-Identifier: LGPL-3.0-or-later
// Copyright 2019 DNA Dev team
//
/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */
package client

import (
	"encoding/json"
	"fmt"
	sdkcom "github.com/DNAProject/DNA-go-sdk/common"
	"github.com/DNAProject/DNA-go-sdk/utils"
	"github.com/DNAProject/DNA/common"
	"github.com/DNAProject/DNA/core/types"
	"sync/atomic"
	"time"
)

type ClientMgr struct {
	rpc       *RpcClient  //Rpc client used the rpc api of dna
	rest      *RestClient //Rest client used the rest api of dna
	ws        *WSClient   //Web socket client used the web socket api of dna
	defClient DNAClient
	qid       uint64
}

func (this *ClientMgr) NewRpcClient() *RpcClient {
	this.rpc = NewRpcClient()
	return this.rpc
}

func (this *ClientMgr) GetRpcClient() *RpcClient {
	return this.rpc
}

func (this *ClientMgr) NewRestClient() *RestClient {
	this.rest = NewRestClient()
	return this.rest
}

func (this *ClientMgr) GetRestClient() *RestClient {
	return this.rest
}

func (this *ClientMgr) NewWebSocketClient() *WSClient {
	wsClient := NewWSClient()
	this.ws = wsClient
	return wsClient
}

func (this *ClientMgr) GetWebSocketClient() *WSClient {
	return this.ws
}

func (this *ClientMgr) SetDefaultClient(client DNAClient) {
	this.defClient = client
}

func (this *ClientMgr) GetCurrentBlockHeight() (uint32, error) {
	client := this.getClient()
	if client == nil {
		return 0, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getCurrentBlockHeight(this.getNextQid())
	if err != nil {
		return 0, err
	}
	return utils.GetUint32(data)
}

func (this *ClientMgr) GetCurrentBlockHash() (common.Uint256, error) {
	client := this.getClient()
	if client == nil {
		return common.UINT256_EMPTY, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getCurrentBlockHash(this.getNextQid())
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return utils.GetUint256(data)
}

func (this *ClientMgr) GetBlockByHeight(height uint32) (*types.Block, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getBlockByHeight(this.getNextQid(), height)
	if err != nil {
		return nil, err
	}
	return utils.GetBlock(data)
}

func (this *ClientMgr) GetBlockInfoByHeight(height uint32) ([]byte, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getBlockInfoByHeight(this.getNextQid(), height)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *ClientMgr) GetBlockByHash(blockHash string) (*types.Block, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getBlockByHash(this.getNextQid(), blockHash)
	if err != nil {
		return nil, err
	}
	return utils.GetBlock(data)
}

func (this *ClientMgr) GetTransaction(txHash string) (*types.Transaction, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getRawTransaction(this.getNextQid(), txHash)
	if err != nil {
		return nil, err
	}
	return utils.GetTransaction(data)
}

func (this *ClientMgr) GetBlockHash(height uint32) (common.Uint256, error) {
	client := this.getClient()
	if client == nil {
		return common.UINT256_EMPTY, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getBlockHash(this.getNextQid(), height)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return utils.GetUint256(data)
}

func (this *ClientMgr) GetBlockHeightByTxHash(txHash string) (uint32, error) {
	client := this.getClient()
	if client == nil {
		return 0, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getBlockHeightByTxHash(this.getNextQid(), txHash)
	if err != nil {
		return 0, err
	}
	return utils.GetUint32(data)
}

func (this *ClientMgr) GetBlockTxHashesByHeight(height uint32) (*sdkcom.BlockTxHashes, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getBlockTxHashesByHeight(this.getNextQid(), height)
	if err != nil {
		return nil, err
	}
	return utils.GetBlockTxHashes(data)
}

func (this *ClientMgr) GetStorage(contractAddress string, key []byte) ([]byte, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getStorage(this.getNextQid(), contractAddress, key)
	if err != nil {
		return nil, err
	}
	return utils.GetStorage(data)
}

func (this *ClientMgr) GetSmartContract(contractAddress string) (*sdkcom.SmartContract, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getSmartContract(this.getNextQid(), contractAddress)
	if err != nil {
		return nil, err
	}
	deployCode, err := utils.GetSmartContract(data)
	if err != nil {
		return nil, err
	}
	sm := sdkcom.SmartContract(*deployCode)
	return &sm, nil
}

func (this *ClientMgr) GetSmartContractEvent(txHash string) (*sdkcom.SmartContactEvent, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getSmartContractEvent(this.getNextQid(), txHash)
	if err != nil {
		return nil, err
	}
	return utils.GetSmartContractEvent(data)
}

func (this *ClientMgr) GetSmartContractEventByBlock(height uint32) ([]*sdkcom.SmartContactEvent, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getSmartContractEventByBlock(this.getNextQid(), height)
	if err != nil {
		return nil, err
	}
	return utils.GetSmartContactEvents(data)
}

func (this *ClientMgr) GetMerkleProof(txHash string) (*sdkcom.MerkleProof, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getMerkleProof(this.getNextQid(), txHash)
	if err != nil {
		return nil, err
	}
	return utils.GetMerkleProof(data)
}

func (this *ClientMgr) GetMemPoolTxState(txHash string) (*sdkcom.MemPoolTxState, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getMemPoolTxState(this.getNextQid(), txHash)
	if err != nil {
		return nil, err
	}
	return utils.GetMemPoolTxState(data)
}

func (this *ClientMgr) GetMemPoolTxCount() (*sdkcom.MemPoolTxCount, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getMemPoolTxCount(this.getNextQid())
	if err != nil {
		return nil, err
	}
	return utils.GetMemPoolTxCount(data)
}

func (this *ClientMgr) GetVersion() (string, error) {
	client := this.getClient()
	if client == nil {
		return "", fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getVersion(this.getNextQid())
	if err != nil {
		return "", err
	}
	return utils.GetVersion(data)
}

func (this *ClientMgr) GetNetworkId() (uint32, error) {
	client := this.getClient()
	if client == nil {
		return 0, fmt.Errorf("don't have available client of dna")
	}
	data, err := client.getNetworkId(this.getNextQid())
	if err != nil {
		return 0, err
	}
	return utils.GetUint32(data)
}

func (this *ClientMgr) SendTransaction(mutTx *types.MutableTransaction) (common.Uint256, error) {
	client := this.getClient()
	if client == nil {
		return common.UINT256_EMPTY, fmt.Errorf("don't have available client of dna")
	}
	tx, err := mutTx.IntoImmutable()
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	data, err := client.sendRawTransaction(this.getNextQid(), tx, false)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	return utils.GetUint256(data)
}

func (this *ClientMgr) PreExecTransaction(mutTx *types.MutableTransaction) (*sdkcom.PreExecResult, error) {
	client := this.getClient()
	if client == nil {
		return nil, fmt.Errorf("don't have available client of dna")
	}
	tx, err := mutTx.IntoImmutable()
	if err != nil {
		return nil, err
	}
	data, err := client.sendRawTransaction(this.getNextQid(), tx, true)
	if err != nil {
		return nil, err
	}
	preResult := &sdkcom.PreExecResult{}
	err = json.Unmarshal(data, &preResult)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal PreExecResult:%s error:%s", data, err)
	}
	return preResult, nil
}

//WaitForGenerateBlock Wait dna generate block. Default wait 2 blocks.
//return timeout error when there is no block generate in some time.
func (this *ClientMgr) WaitForGenerateBlock(timeout time.Duration, blockCount ...uint32) (bool, error) {
	count := uint32(2)
	if len(blockCount) > 0 && blockCount[0] > 0 {
		count = blockCount[0]
	}
	blockHeight, err := this.GetCurrentBlockHeight()
	if err != nil {
		return false, fmt.Errorf("GetCurrentBlockHeight error:%s", err)
	}
	secs := int(timeout / time.Second)
	if secs <= 0 {
		secs = 1
	}
	for i := 0; i < secs; i++ {
		time.Sleep(time.Second)
		curBlockHeigh, err := this.GetCurrentBlockHeight()
		if err != nil {
			continue
		}
		if curBlockHeigh-blockHeight >= count {
			return true, nil
		}
	}
	return false, fmt.Errorf("timeout after %d (s)", secs)
}

func (this *ClientMgr) getClient() DNAClient {
	if this.defClient != nil {
		return this.defClient
	}
	if this.rpc != nil {
		return this.rpc
	}
	if this.rest != nil {
		return this.rest
	}
	if this.ws != nil {
		return this.ws
	}
	return nil
}

func (this *ClientMgr) getNextQid() string {
	return fmt.Sprintf("%d", atomic.AddUint64(&this.qid, 1))
}
