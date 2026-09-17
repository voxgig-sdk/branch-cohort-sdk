-- BranchCohort SDK exists test

local sdk = require("branch-cohort_sdk")

describe("BranchCohortSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
