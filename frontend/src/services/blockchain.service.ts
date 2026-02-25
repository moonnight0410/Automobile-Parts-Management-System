import axios from './axios'

export interface BlockchainInfo {
  channel: string
  status: string
  blockHeight: number
  txCount: number
  nodeCount: number
  channelCount: number
}

export interface BlockInfo {
  blockNumber: number
  txCount: number
  blockHash: string
  timestamp: string
}

export async function getBlockchainInfo(): Promise<{ code: number; data: BlockchainInfo; message: string }> {
  const response = await axios.get('/api/fabric/blockchain/info')
  return response.data
}

export async function getRecentBlocks(limit: number = 10): Promise<{ code: number; data: BlockInfo[]; message: string }> {
  const response = await axios.get('/api/fabric/blockchain/blocks', { params: { limit } })
  return response.data
}
