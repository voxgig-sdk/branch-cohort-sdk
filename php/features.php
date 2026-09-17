<?php
declare(strict_types=1);

// BranchCohort SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/DebugFeature.php';
require_once __DIR__ . '/feature/IdempotencyFeature.php';
require_once __DIR__ . '/feature/MetricsFeature.php';
require_once __DIR__ . '/feature/PagingFeature.php';
require_once __DIR__ . '/feature/RatelimitFeature.php';
require_once __DIR__ . '/feature/RetryFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';
require_once __DIR__ . '/feature/TimeoutFeature.php';


class BranchCohortFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new BranchCohortBaseFeature();
            case "debug":
                return new BranchCohortDebugFeature();
            case "idempotency":
                return new BranchCohortIdempotencyFeature();
            case "metrics":
                return new BranchCohortMetricsFeature();
            case "paging":
                return new BranchCohortPagingFeature();
            case "ratelimit":
                return new BranchCohortRatelimitFeature();
            case "retry":
                return new BranchCohortRetryFeature();
            case "test":
                return new BranchCohortTestFeature();
            case "timeout":
                return new BranchCohortTimeoutFeature();
            default:
                return new BranchCohortBaseFeature();
        }
    }

    /**
     * Does a generated feature class back this name? False for a name only
     * an options extend instance can supply (the station adopt path) - the
     * constructor uses this to skip make_feature for such names instead of
     * adding a stray BaseFeature.
     */
    public static function has_feature(string $name): bool
    {
        switch ($name) {
            case "base":
            case "debug":
            case "idempotency":
            case "metrics":
            case "paging":
            case "ratelimit":
            case "retry":
            case "test":
            case "timeout":
                return true;
            default:
                return false;
        }
    }
}
