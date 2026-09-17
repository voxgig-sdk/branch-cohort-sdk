
const { test, describe } = require('node:test')
const { equal } = require('node:assert')


const { BranchCohortSDK } = require('..')


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await BranchCohortSDK.test()
    equal(null !== testsdk, true)
  })

})
