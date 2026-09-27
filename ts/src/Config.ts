
import { BaseFeature } from './feature/base/BaseFeature'
import { DebugFeature } from './feature/debug/DebugFeature'
import { IdempotencyFeature } from './feature/idempotency/IdempotencyFeature'
import { MetricsFeature } from './feature/metrics/MetricsFeature'
import { PagingFeature } from './feature/paging/PagingFeature'
import { RatelimitFeature } from './feature/ratelimit/RatelimitFeature'
import { RetryFeature } from './feature/retry/RetryFeature'
import { TestFeature } from './feature/test/TestFeature'
import { TimeoutFeature } from './feature/timeout/TimeoutFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   debug: DebugFeature,
 idempotency: IdempotencyFeature,
 metrics: MetricsFeature,
 paging: PagingFeature,
 ratelimit: RatelimitFeature,
 retry: RetryFeature,
 test: TestFeature,
 timeout: TimeoutFeature,

}


const FEATURE_PLUGINS: Record<string, any[]> = {
  
}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    return fi
  }

  // False for a feature added at runtime via options.extend (station's
  // adopt path) - the constructor uses this to skip makeFeature for names
  // no generated class backs.
  hasFeature(this: any, fn: string) {
    return null != FEATURE_CLASS[fn]
  }


  main = {
    name: 'BranchCohort',
        slug: "branch-cohort",
    version: "0.0.1",
    target: "ts",

  }


  feature = {
     debug:     {
      "options": {
        "active": false,
        "max": 100,
        "redact": [
          "authorization",
          "cookie",
          "set-cookie",
          "api-key",
          "apikey",
          "x-api-key",
          "idempotency-key"
        ]
      },
      "optspec": {
        "now": "`$FUNCTION`",
        "onEntry": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "none"
    },
 idempotency:     {
      "options": {
        "active": false,
        "header": "Idempotency-Key",
        "methods": [
          "POST",
          "PUT",
          "PATCH",
          "DELETE"
        ],
        "ops": [
          "create",
          "update",
          "remove"
        ]
      },
      "optspec": {
        "keygen": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "none"
    },
 metrics:     {
      "options": {
        "active": false
      },
      "optspec": {
        "now": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "none"
    },
 paging:     {
      "options": {
        "active": false,
        "afterVar": "after",
        "cursorParam": "cursor",
        "firstVar": "first",
        "limitParam": "limit",
        "pageParam": "page",
        "startPage": 1
      },
      "optspec": {
        "limit": "`$NUMBER`",
        "ops": "`$LIST`"
      },
      "strict": false,
      "transport": "none"
    },
 ratelimit:     {
      "options": {
        "active": false,
        "burst": 5,
        "rate": 5
      },
      "optspec": {
        "now": "`$FUNCTION`",
        "sleep": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },
 retry:     {
      "options": {
        "active": false,
        "factor": 2,
        "maxDelay": 2000,
        "minDelay": 50,
        "retries": 2,
        "statuses": [
          408,
          425,
          429,
          500,
          502,
          503,
          504
        ]
      },
      "optspec": {
        "jitter": "`$BOOLEAN`",
        "sleep": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },
 test:     {
      "options": {
        "active": false
      },
      "optspec": {
        "entity": "`$MAP`",
        "net": "`$MAP`"
      },
      "strict": false,
      "transport": "base"
    },
 timeout:     {
      "options": {
        "active": false,
        "ms": 30000
      },
      "optspec": {
        "clearTimer": "`$FUNCTION`",
        "setTimer": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },

  }


  options = {
    base: "https://api2.branch.io/v2",

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
        analytics: {
        },
  
    }
  }


  entity = {
    "analytics": {
      "fields": [
        {
          "name": "code",
          "title": "Code",
          "type": "`$STRING`",
          "short": "HTTP code representing the outcome of the status request."
        },
        {
          "name": "cumulative",
          "title": "Cumulative",
          "type": "`$BOOLEAN`",
          "short": "If true, sum across bands so that a given band value is the sum of all preceding values plus the band value."
        },
        {
          "name": "data_source",
          "title": "Data Source",
          "type": "`$STRING`",
          "req": true,
          "short": "A string value representing the cohort type"
        },
        {
          "name": "dimensions",
          "title": "Dimensions",
          "type": "`$ARRAY`",
          "short": "An array representing dimension(s) to group by."
        },
        {
          "name": "enable_install_recalculation",
          "title": "Enable Install Recalculation",
          "type": "`$BOOLEAN`",
          "short": "If true, then Branch will de-dupe unattributed installs caused by duplicate events from non-opt-in users coming from paid ads."
        },
        {
          "name": "end_date",
          "title": "End Date",
          "type": "`$STRING`",
          "req": true,
          "short": "The end of the interval time range represented as an ISO-8601 complete date.",
          "format": "date"
        },
        {
          "name": "error_message",
          "title": "Error Message",
          "type": "`$STRING`",
          "short": "Error message if the query failed"
        },
        {
          "name": "filter",
          "title": "Filter",
          "type": "`$OBJECT`",
          "short": "\"Keys are same as dimensions."
        },
        {
          "name": "granularity",
          "title": "Granularity",
          "type": "`$STRING`",
          "short": "The time granularity that each band value will represent."
        },
        {
          "name": "granularity_band_count",
          "title": "Granularity Band Count",
          "type": "`$INTEGER`",
          "req": true,
          "short": "Number of time units since the cohort event to return to the user."
        },
        {
          "name": "id",
          "title": "Id",
          "type": "`$STRING`"
        },
        {
          "name": "job_id",
          "title": "Job Id",
          "type": "`$STRING`",
          "short": "Unique identifier used to retrieve job status and data."
        },
        {
          "name": "measures",
          "title": "Measures",
          "type": "`$ARRAY`",
          "req": true,
          "short": "The cohort measures to return."
        },
        {
          "name": "ordered",
          "title": "Ordered",
          "type": "`$STRING`",
          "short": "Order of response based on ordered_by value."
        },
        {
          "name": "ordered_by",
          "title": "Ordered By",
          "type": "`$STRING`",
          "short": "The dimension used for sorting"
        },
        {
          "name": "per_user",
          "title": "Per User",
          "type": "`$BOOLEAN`",
          "short": "If true, divide each band value by the user count."
        },
        {
          "name": "response_url",
          "title": "Response Url",
          "type": "`$STRING`",
          "short": "S3 url for downloading the response data."
        },
        {
          "name": "start_date",
          "title": "Start Date",
          "type": "`$STRING`",
          "req": true,
          "short": "The start of the interval time range represented as an ISO-8601 complete date.",
          "format": "date"
        },
        {
          "name": "status",
          "title": "Status",
          "type": "`$STRING`",
          "short": "Status of the query."
        },
        {
          "name": "status_url",
          "title": "Status Url",
          "type": "`$STRING`",
          "short": "More information about the subscription's status."
        },
        {
          "name": "unique",
          "title": "Unique",
          "type": "`$BOOLEAN`",
          "short": "Whether or not to return unique values."
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
      "name": "analytics",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "kind": "http",
              "method": "POST",
              "orig": "/analytics",
              "segments": [
                {
                  "lit": "analytics"
                }
              ],
              "parts": [
                "analytics"
              ],
              "rename": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "query": [
                  {
                    "name": "app_id",
                    "orig": "app_id",
                    "type": "`$STRING`",
                    "kind": "query",
                    "reqd": true
                  },
                  {
                    "name": "format",
                    "orig": "format",
                    "type": "`$STRING`",
                    "kind": "query",
                    "reqd": true,
                    "example": "csv"
                  },
                  {
                    "name": "limit",
                    "orig": "limit",
                    "type": "`$INTEGER`",
                    "kind": "query",
                    "reqd": true
                  }
                ]
              },
              "select": {
                "exist": [
                  "app_id",
                  "format",
                  "limit"
                ]
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "kind": "http",
              "method": "GET",
              "orig": "/analytics/{job_id}",
              "segments": [
                {
                  "lit": "analytics"
                },
                {
                  "var": "id"
                }
              ],
              "parts": [
                "analytics",
                "{id}"
              ],
              "rename": {
                "param": {
                  "job_id": "id"
                }
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "args": {
                "params": [
                  {
                    "name": "id",
                    "orig": "job_id",
                    "type": "`$STRING`",
                    "kind": "param",
                    "reqd": true,
                    "example": "0000-XXxx"
                  }
                ],
                "query": [
                  {
                    "name": "app_id",
                    "orig": "app_id",
                    "type": "`$STRING`",
                    "kind": "query",
                    "reqd": true
                  },
                  {
                    "name": "format",
                    "orig": "format",
                    "type": "`$STRING`",
                    "kind": "query",
                    "reqd": true,
                    "example": "csv"
                  }
                ]
              },
              "select": {
                "exist": [
                  "app_id",
                  "format",
                  "id"
                ]
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    }
  }
}


const config = new Config()

export {
  config,
  FEATURE_PLUGINS,
}

