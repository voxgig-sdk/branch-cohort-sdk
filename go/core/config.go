package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "BranchCohort",
			"slug": "branch-cohort",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"debug": map[string]any{
				"options": map[string]any{
					"active": false,
					"max": 100,
					"redact": []any{
						"authorization",
						"cookie",
						"set-cookie",
						"api-key",
						"apikey",
						"x-api-key",
						"idempotency-key",
					},
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"onEntry": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"idempotency": map[string]any{
				"options": map[string]any{
					"active": false,
					"header": "Idempotency-Key",
					"methods": []any{
						"POST",
						"PUT",
						"PATCH",
						"DELETE",
					},
					"ops": []any{
						"create",
						"update",
						"remove",
					},
				},
				"optspec": map[string]any{
					"keygen": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"metrics": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"paging": map[string]any{
				"options": map[string]any{
					"active": false,
					"afterVar": "after",
					"cursorParam": "cursor",
					"firstVar": "first",
					"limitParam": "limit",
					"pageParam": "page",
					"startPage": 1,
				},
				"optspec": map[string]any{
					"limit": "`$NUMBER`",
					"ops": "`$LIST`",
				},
				"strict": false,
				"transport": "none",
			},
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://api2.branch.io/v2",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"analytics": map[string]any{},
			},
		},
		"entity": map[string]any{
			"analytics": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "code",
						"short": "HTTP code representing the outcome of the status request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cumulative",
						"short": "If true, sum across bands so that a given band value is the sum of all preceding values plus the band value.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "data_source",
						"req": true,
						"short": "A string value representing the cohort type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "dimensions",
						"short": "An array representing dimension(s) to group by.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "enable_install_recalculation",
						"short": "If true, then Branch will de-dupe unattributed installs caused by duplicate events from non-opt-in users coming from paid ads.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"format": "date",
						"name": "end_date",
						"req": true,
						"short": "The end of the interval time range represented as an ISO-8601 complete date.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "error_message",
						"short": "Error message if the query failed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "filter",
						"short": "\"Keys are same as dimensions.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "granularity",
						"short": "The time granularity that each band value will represent.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "granularity_band_count",
						"req": true,
						"short": "Number of time units since the cohort event to return to the user.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "job_id",
						"short": "Unique identifier used to retrieve job status and data.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "measures",
						"req": true,
						"short": "The cohort measures to return.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "ordered",
						"short": "Order of response based on ordered_by value.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ordered_by",
						"short": "The dimension used for sorting",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "per_user",
						"short": "If true, divide each band value by the user count.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "response_url",
						"short": "S3 url for downloading the response data.",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date",
						"name": "start_date",
						"req": true,
						"short": "The start of the interval time range represented as an ISO-8601 complete date.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "Status of the query.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status_url",
						"short": "More information about the subscription's status.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "unique",
						"short": "Whether or not to return unique values.",
						"type": "`$BOOLEAN`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "analytics",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "app_id",
											"orig": "app_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "csv",
											"kind": "query",
											"name": "format",
											"orig": "format",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/analytics",
								"segments": []any{
									map[string]any{
										"lit": "analytics",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"app_id",
										"format",
										"limit",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"analytics",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "0000-XXxx",
											"kind": "param",
											"name": "id",
											"orig": "job_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "app_id",
											"orig": "app_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "csv",
											"kind": "query",
											"name": "format",
											"orig": "format",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/analytics/{job_id}",
								"rename": map[string]any{
									"param": map[string]any{
										"job_id": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "analytics",
									},
									map[string]any{
										"var": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"app_id",
										"format",
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"analytics",
									"{id}",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "debug":
		if NewDebugFeatureFunc != nil {
			return NewDebugFeatureFunc()
		}
	case "idempotency":
		if NewIdempotencyFeatureFunc != nil {
			return NewIdempotencyFeatureFunc()
		}
	case "metrics":
		if NewMetricsFeatureFunc != nil {
			return NewMetricsFeatureFunc()
		}
	case "paging":
		if NewPagingFeatureFunc != nil {
			return NewPagingFeatureFunc()
		}
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
