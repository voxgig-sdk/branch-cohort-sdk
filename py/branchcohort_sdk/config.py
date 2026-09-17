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
            "short": "HTTP code representing the outcome of the status request.",
            "type": "`$STRING`",
          },
          {
            "name": "cumulative",
            "short": "If true, sum across bands so that a given band value is the sum of all preceding values plus the band value.",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "data_source",
            "req": True,
            "short": "A string value representing the cohort type",
            "type": "`$STRING`",
          },
          {
            "name": "dimensions",
            "short": "An array representing dimension(s) to group by.",
            "type": "`$ARRAY`",
          },
          {
            "name": "enable_install_recalculation",
            "short": "If true, then Branch will de-dupe unattributed installs caused by duplicate events from non-opt-in users coming from paid ads.",
            "type": "`$BOOLEAN`",
          },
          {
            "format": "date",
            "name": "end_date",
            "req": True,
            "short": "The end of the interval time range represented as an ISO-8601 complete date.",
            "type": "`$STRING`",
          },
          {
            "name": "error_message",
            "short": "Error message if the query failed",
            "type": "`$STRING`",
          },
          {
            "name": "filter",
            "short": "\"Keys are same as dimensions.",
            "type": "`$OBJECT`",
          },
          {
            "name": "granularity",
            "short": "The time granularity that each band value will represent.",
            "type": "`$STRING`",
          },
          {
            "name": "granularity_band_count",
            "req": True,
            "short": "Number of time units since the cohort event to return to the user.",
            "type": "`$INTEGER`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "job_id",
            "short": "Unique identifier used to retrieve job status and data.",
            "type": "`$STRING`",
          },
          {
            "name": "measures",
            "req": True,
            "short": "The cohort measures to return.",
            "type": "`$ARRAY`",
          },
          {
            "name": "ordered",
            "short": "Order of response based on ordered_by value.",
            "type": "`$STRING`",
          },
          {
            "name": "ordered_by",
            "short": "The dimension used for sorting",
            "type": "`$STRING`",
          },
          {
            "name": "per_user",
            "short": "If true, divide each band value by the user count.",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "response_url",
            "short": "S3 url for downloading the response data.",
            "type": "`$STRING`",
          },
          {
            "format": "date",
            "name": "start_date",
            "req": True,
            "short": "The start of the interval time range represented as an ISO-8601 complete date.",
            "type": "`$STRING`",
          },
          {
            "name": "status",
            "short": "Status of the query.",
            "type": "`$STRING`",
          },
          {
            "name": "status_url",
            "short": "More information about the subscription's status.",
            "type": "`$STRING`",
          },
          {
            "name": "unique",
            "short": "Whether or not to return unique values.",
            "type": "`$BOOLEAN`",
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
                "args": {
                  "query": [
                    {
                      "kind": "query",
                      "name": "app_id",
                      "orig": "app_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "csv",
                      "kind": "query",
                      "name": "format",
                      "orig": "format",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "kind": "query",
                      "name": "limit",
                      "orig": "limit",
                      "reqd": True,
                      "type": "`$INTEGER`",
                    },
                  ],
                },
                "kind": "http",
                "method": "POST",
                "orig": "/analytics",
                "segments": [
                  {
                    "lit": "analytics",
                  },
                ],
                "select": {
                  "exist": [
                    "app_id",
                    "format",
                    "limit",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "analytics",
                ],
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "0000-XXxx",
                      "kind": "param",
                      "name": "id",
                      "orig": "job_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "kind": "query",
                      "name": "app_id",
                      "orig": "app_id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "csv",
                      "kind": "query",
                      "name": "format",
                      "orig": "format",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/analytics/{job_id}",
                "rename": {
                  "param": {
                    "job_id": "id",
                  },
                },
                "segments": [
                  {
                    "lit": "analytics",
                  },
                  {
                    "var": "id",
                  },
                ],
                "select": {
                  "exist": [
                    "app_id",
                    "format",
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "analytics",
                  "{id}",
                ],
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
