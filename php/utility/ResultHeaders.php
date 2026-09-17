<?php
declare(strict_types=1);

// BranchCohort SDK utility: result_headers

class BranchCohortResultHeaders
{
    public static function call(BranchCohortContext $ctx): ?BranchCohortResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
