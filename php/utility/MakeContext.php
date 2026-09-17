<?php
declare(strict_types=1);

// BranchCohort SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class BranchCohortMakeContext
{
    public static function call(array $ctxmap, ?BranchCohortContext $basectx): BranchCohortContext
    {
        return new BranchCohortContext($ctxmap, $basectx);
    }
}
