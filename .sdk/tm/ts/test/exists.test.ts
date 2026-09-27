
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { BranchCohortSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = BranchCohortSDK.test()
    equal(testsdk instanceof BranchCohortSDK, true,
      'BranchCohortSDK.test() must return a client synchronously')
  })

})
