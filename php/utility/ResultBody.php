<?php
declare(strict_types=1);

// BranchCohort SDK utility: result_body

class BranchCohortResultBody
{
    public static function call(BranchCohortContext $ctx): ?BranchCohortResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
