# Typed models for the BranchCohort SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class AnalyticsRequired(TypedDict):
    data_source: str
    end_date: str
    granularity_band_count: int
    measures: list
    start_date: str


class Analytics(AnalyticsRequired, total=False):
    code: str
    cumulative: bool
    dimensions: list
    enable_install_recalculation: bool
    error_message: str
    filter: dict
    granularity: str
    id: str
    job_id: str
    ordered: str
    ordered_by: str
    per_user: bool
    response_url: str
    status: str
    status_url: str
    unique: bool


class AnalyticsLoadMatch(TypedDict):
    id: str
    app_id: str
    format: str


class AnalyticsCreateDataRequired(TypedDict):
    app_id: str
    format: str
    limit: int
    data_source: str
    end_date: str
    granularity_band_count: int
    measures: list
    start_date: str


class AnalyticsCreateData(AnalyticsCreateDataRequired, total=False):
    code: str
    cumulative: bool
    dimensions: list
    enable_install_recalculation: bool
    error_message: str
    filter: dict
    granularity: str
    id: str
    job_id: str
    ordered: str
    ordered_by: str
    per_user: bool
    response_url: str
    status: str
    status_url: str
    unique: bool
