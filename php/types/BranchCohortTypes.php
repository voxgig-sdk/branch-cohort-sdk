<?php
declare(strict_types=1);

// Typed models for the BranchCohort SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Analytics entity data model. */
class Analytics
{
    public ?string $code = null;
    public ?bool $cumulative = null;
    public string $data_source;
    public ?array $dimensions = null;
    public ?bool $enable_install_recalculation = null;
    public string $end_date;
    public ?string $error_message = null;
    public ?array $filter = null;
    public ?string $granularity = null;
    public int $granularity_band_count;
    public ?string $id = null;
    public ?string $job_id = null;
    public array $measures;
    public ?string $ordered = null;
    public ?string $ordered_by = null;
    public ?bool $per_user = null;
    public ?string $response_url = null;
    public string $start_date;
    public ?string $status = null;
    public ?string $status_url = null;
    public ?bool $unique = null;
}

/** Request payload for Analytics#load. */
class AnalyticsLoadMatch
{
    public string $id;
    public string $app_id;
    public string $format;
}

/** Request payload for Analytics#create. */
class AnalyticsCreateData
{
    public string $app_id;
    public string $format;
    public int $limit;
    public ?string $code = null;
    public ?bool $cumulative = null;
    public string $data_source;
    public ?array $dimensions = null;
    public ?bool $enable_install_recalculation = null;
    public string $end_date;
    public ?string $error_message = null;
    public ?array $filter = null;
    public ?string $granularity = null;
    public int $granularity_band_count;
    public ?string $id = null;
    public ?string $job_id = null;
    public array $measures;
    public ?string $ordered = null;
    public ?string $ordered_by = null;
    public ?bool $per_user = null;
    public ?string $response_url = null;
    public string $start_date;
    public ?string $status = null;
    public ?string $status_url = null;
    public ?bool $unique = null;
}

