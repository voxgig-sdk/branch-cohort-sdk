// Typed models for the BranchCohort SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Analytics {
  code?: string
  cumulative?: boolean
  data_source: string
  dimensions?: any[]
  enable_install_recalculation?: boolean
  end_date: string
  error_message?: string
  filter?: Record<string, any>
  granularity?: string
  granularity_band_count: number
  id?: string
  job_id?: string
  measures: any[]
  ordered?: string
  ordered_by?: string
  per_user?: boolean
  response_url?: string
  start_date: string
  status?: string
  status_url?: string
  unique?: boolean
}

export interface AnalyticsLoadMatch {
  id: string
  app_id: string
  format: string
}

export interface AnalyticsCreateData {
  app_id: string
  format: string
  limit: number
  code?: string
  cumulative?: boolean
  data_source: string
  dimensions?: any[]
  enable_install_recalculation?: boolean
  end_date: string
  error_message?: string
  filter?: Record<string, any>
  granularity?: string
  granularity_band_count: number
  id?: string
  job_id?: string
  measures: any[]
  ordered?: string
  ordered_by?: string
  per_user?: boolean
  response_url?: string
  start_date: string
  status?: string
  status_url?: string
  unique?: boolean
}

