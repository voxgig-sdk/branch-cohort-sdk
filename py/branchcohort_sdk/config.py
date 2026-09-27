# BranchCohort SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "BranchCohort",
            "slug": "branch-cohort",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "debug": {
        "options": {
          "active": False,
          "max": 100,
          "redact": [
            "authorization",
            "cookie",
            "set-cookie",
            "api-key",
            "apikey",
            "x-api-key",
            "idempotency-key",
          ],
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "onEntry": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "idempotency": {
        "options": {
          "active": False,
          "header": "Idempotency-Key",
          "methods": [
            "POST",
            "PUT",
            "PATCH",
            "DELETE",
          ],
          "ops": [
            "create",
            "update",
            "remove",
          ],
        },
        "optspec": {
          "keygen": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "metrics": {
        "options": {
          "active": False,
        },
        "optspec": {
          "now": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "paging": {
        "options": {
          "active": False,
          "afterVar": "after",
          "cursorParam": "cursor",
          "firstVar": "first",
          "limitParam": "limit",
          "pageParam": "page",
          "startPage": 1,
        },
        "optspec": {
          "limit": "`$NUMBER`",
          "ops": "`$LIST`",
        },
        "strict": False,
        "transport": "none",
      },
            "ratelimit": {
        "options": {
          "active": False,
          "burst": 5,
          "rate": 5,
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "retry": {
        "options": {
          "active": False,
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
            504,
          ],
        },
        "optspec": {
          "jitter": "`$BOOLEAN`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "test": {
        "options": {
          "active": False,
        },
        "optspec": {
          "entity": "`$MAP`",
          "net": "`$MAP`",
        },
        "strict": False,
        "transport": "base",
      },
            "timeout": {
        "options": {
          "active": False,
          "ms": 30000,
        },
        "optspec": {
          "clearTimer": "`$FUNCTION`",
          "setTimer": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
        },
        "options": {
            "base": "https://api2.branch.io/v2",
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "analytics": {},
            },
        },
        "entity": {
      "analytics": {
        "fields": [
          {
            "name": "code",
            "title": "Code",
            "type": "`$STRING`",
            "short": "HTTP code representing the outcome of the status request.",
          },
          {
            "name": "cumulative",
            "title": "Cumulative",
            "type": "`$BOOLEAN`",
            "short": "If true, sum across bands so that a given band value is the sum of all preceding values plus the band value.",
          },
          {
            "name": "data_source",
            "title": "Data Source",
            "type": "`$STRING`",
            "req": True,
            "short": "A string value representing the cohort type",
          },
          {
            "name": "dimensions",
            "title": "Dimensions",
            "type": "`$ARRAY`",
            "short": "An array representing dimension(s) to group by.",
          },
          {
            "name": "enable_install_recalculation",
            "title": "Enable Install Recalculation",
            "type": "`$BOOLEAN`",
            "short": "If true, then Branch will de-dupe unattributed installs caused by duplicate events from non-opt-in users coming from paid ads.",
          },
          {
            "name": "end_date",
            "title": "End Date",
            "type": "`$STRING`",
            "req": True,
            "short": "The end of the interval time range represented as an ISO-8601 complete date.",
            "format": "date",
          },
          {
            "name": "error_message",
            "title": "Error Message",
            "type": "`$STRING`",
            "short": "Error message if the query failed",
          },
          {
            "name": "filter",
            "title": "Filter",
            "type": "`$OBJECT`",
            "short": "\"Keys are same as dimensions.",
          },
          {
            "name": "granularity",
            "title": "Granularity",
            "type": "`$STRING`",
            "short": "The time granularity that each band value will represent.",
          },
          {
            "name": "granularity_band_count",
            "title": "Granularity Band Count",
            "type": "`$INTEGER`",
            "req": True,
            "short": "Number of time units since the cohort event to return to the user.",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
          },
          {
            "name": "job_id",
            "title": "Job Id",
            "type": "`$STRING`",
            "short": "Unique identifier used to retrieve job status and data.",
          },
          {
            "name": "measures",
            "title": "Measures",
            "type": "`$ARRAY`",
            "req": True,
            "short": "The cohort measures to return.",
          },
          {
            "name": "ordered",
            "title": "Ordered",
            "type": "`$STRING`",
            "short": "Order of response based on ordered_by value.",
          },
          {
            "name": "ordered_by",
            "title": "Ordered By",
            "type": "`$STRING`",
            "short": "The dimension used for sorting",
          },
          {
            "name": "per_user",
            "title": "Per User",
            "type": "`$BOOLEAN`",
            "short": "If true, divide each band value by the user count.",
          },
          {
            "name": "response_url",
            "title": "Response Url",
            "type": "`$STRING`",
            "short": "S3 url for downloading the response data.",
          },
          {
            "name": "start_date",
            "title": "Start Date",
            "type": "`$STRING`",
            "req": True,
            "short": "The start of the interval time range represented as an ISO-8601 complete date.",
            "format": "date",
          },
          {
            "name": "status",
            "title": "Status",
            "type": "`$STRING`",
            "short": "Status of the query.",
          },
          {
            "name": "status_url",
            "title": "Status Url",
            "type": "`$STRING`",
            "short": "More information about the subscription's status.",
          },
          {
            "name": "unique",
            "title": "Unique",
            "type": "`$BOOLEAN`",
            "short": "Whether or not to return unique values.",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
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
                    "lit": "analytics",
                  },
                ],
                "parts": [
                  "analytics",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "query": [
                    {
                      "name": "app_id",
                      "orig": "app_id",
                      "type": "`$STRING`",
                      "kind": "query",
                      "reqd": True,
                    },
                    {
                      "name": "format",
                      "orig": "format",
                      "type": "`$STRING`",
                      "kind": "query",
                      "reqd": True,
                      "example": "csv",
                    },
                    {
                      "name": "limit",
                      "orig": "limit",
                      "type": "`$INTEGER`",
                      "kind": "query",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "app_id",
                    "format",
                    "limit",
                  ],
                },
              },
            ],
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
                    "lit": "analytics",
                  },
                  {
                    "var": "id",
                  },
                ],
                "parts": [
                  "analytics",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "job_id": "id",
                  },
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "job_id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                      "example": "0000-XXxx",
                    },
                  ],
                  "query": [
                    {
                      "name": "app_id",
                      "orig": "app_id",
                      "type": "`$STRING`",
                      "kind": "query",
                      "reqd": True,
                    },
                    {
                      "name": "format",
                      "orig": "format",
                      "type": "`$STRING`",
                      "kind": "query",
                      "reqd": True,
                      "example": "csv",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "app_id",
                    "format",
                    "id",
                  ],
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
