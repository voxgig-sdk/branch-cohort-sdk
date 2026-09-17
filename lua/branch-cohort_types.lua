-- Typed models for the BranchCohort SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Analytics
---@field code? string
---@field cumulative? boolean
---@field data_source string
---@field dimensions? table
---@field enable_install_recalculation? boolean
---@field end_date string
---@field error_message? string
---@field filter? table
---@field granularity? string
---@field granularity_band_count number
---@field id? string
---@field job_id? string
---@field measures table
---@field ordered? string
---@field ordered_by? string
---@field per_user? boolean
---@field response_url? string
---@field start_date string
---@field status? string
---@field status_url? string
---@field unique? boolean

---@class AnalyticsLoadMatch
---@field id string
---@field app_id string
---@field format string

---@class AnalyticsCreateData
---@field app_id string
---@field format string
---@field limit number
---@field code? string
---@field cumulative? boolean
---@field data_source string
---@field dimensions? table
---@field enable_install_recalculation? boolean
---@field end_date string
---@field error_message? string
---@field filter? table
---@field granularity? string
---@field granularity_band_count number
---@field id? string
---@field job_id? string
---@field measures table
---@field ordered? string
---@field ordered_by? string
---@field per_user? boolean
---@field response_url? string
---@field start_date string
---@field status? string
---@field status_url? string
---@field unique? boolean

local M = {}

return M
