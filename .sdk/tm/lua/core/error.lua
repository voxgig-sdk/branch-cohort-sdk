-- BranchCohort SDK error

local BranchCohortError = {}
BranchCohortError.__index = BranchCohortError


function BranchCohortError.new(code, msg, ctx)
  local self = setmetatable({}, BranchCohortError)
  self.is_sdk_error = true
  self.sdk = "BranchCohort"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function BranchCohortError:error()
  return self.msg
end


function BranchCohortError:__tostring()
  return self.msg
end


return BranchCohortError
