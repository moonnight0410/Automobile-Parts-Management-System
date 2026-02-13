/**
 * 示例单元测试
 * 测试基本功能
 */

import { describe, it, expect } from 'vitest'

describe('示例测试', () => {
  it('应该正确计算加法', () => {
    expect(1 + 1).toBe(2)
  })

  it('应该正确处理字符串', () => {
    const str = 'Hello World'
    expect(str).toContain('Hello')
    expect(str.length).toBe(11)
  })

  it('应该正确处理数组', () => {
    const arr = [1, 2, 3]
    expect(arr).toHaveLength(3)
    expect(arr).toContain(2)
  })
})
