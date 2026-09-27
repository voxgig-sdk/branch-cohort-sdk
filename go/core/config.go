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
						"title": "Code",
						"type": "`$STRING`",
						"short": "HTTP code representing the outcome of the status request.",
					},
					map[string]any{
						"name": "cumulative",
						"title": "Cumulative",
						"type": "`$BOOLEAN`",
						"short": "If true, sum across bands so that a given band value is the sum of all preceding values plus the band value.",
					},
					map[string]any{
						"name": "data_source",
						"title": "Data Source",
						"type": "`$STRING`",
						"req": true,
						"short": "A string value representing the cohort type",
					},
					map[string]any{
						"name": "dimensions",
						"title": "Dimensions",
						"type": "`$ARRAY`",
						"short": "An array representing dimension(s) to group by.",
					},
					map[string]any{
						"name": "enable_install_recalculation",
						"title": "Enable Install Recalculation",
						"type": "`$BOOLEAN`",
						"short": "If true, then Branch will de-dupe unattributed installs caused by duplicate events from non-opt-in users coming from paid ads.",
					},
					map[string]any{
						"name": "end_date",
						"title": "End Date",
						"type": "`$STRING`",
						"req": true,
						"short": "The end of the interval time range represented as an ISO-8601 complete date.",
						"format": "date",
					},
					map[string]any{
						"name": "error_message",
						"title": "Error Message",
						"type": "`$STRING`",
						"short": "Error message if the query failed",
					},
					map[string]any{
						"name": "filter",
						"title": "Filter",
						"type": "`$OBJECT`",
						"short": "\"Keys are same as dimensions.",
					},
					map[string]any{
						"name": "granularity",
						"title": "Granularity",
						"type": "`$STRING`",
						"short": "The time granularity that each band value will represent.",
					},
					map[string]any{
						"name": "granularity_band_count",
						"title": "Granularity Band Count",
						"type": "`$INTEGER`",
						"req": true,
						"short": "Number of time units since the cohort event to return to the user.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "job_id",
						"title": "Job Id",
						"type": "`$STRING`",
						"short": "Unique identifier used to retrieve job status and data.",
					},
					map[string]any{
						"name": "measures",
						"title": "Measures",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The cohort measures to return.",
					},
					map[string]any{
						"name": "ordered",
						"title": "Ordered",
						"type": "`$STRING`",
						"short": "Order of response based on ordered_by value.",
					},
					map[string]any{
						"name": "ordered_by",
						"title": "Ordered By",
						"type": "`$STRING`",
						"short": "The dimension used for sorting",
					},
					map[string]any{
						"name": "per_user",
						"title": "Per User",
						"type": "`$BOOLEAN`",
						"short": "If true, divide each band value by the user count.",
					},
					map[string]any{
						"name": "response_url",
						"title": "Response Url",
						"type": "`$STRING`",
						"short": "S3 url for downloading the response data.",
					},
					map[string]any{
						"name": "start_date",
						"title": "Start Date",
						"type": "`$STRING`",
						"req": true,
						"short": "The start of the interval time range represented as an ISO-8601 complete date.",
						"format": "date",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "Status of the query.",
					},
					map[string]any{
						"name": "status_url",
						"title": "Status Url",
						"type": "`$STRING`",
						"short": "More information about the subscription's status.",
					},
					map[string]any{
						"name": "unique",
						"title": "Unique",
						"type": "`$BOOLEAN`",
						"short": "Whether or not to return unique values.",
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
								"kind": "http",
								"method": "POST",
								"orig": "/analytics",
								"segments": []any{
									map[string]any{
										"lit": "analytics",
									},
								},
								"parts": []any{
									"analytics",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "app_id",
											"orig": "app_id",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
										},
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
											"example": "csv",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"app_id",
										"format",
										"limit",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/analytics/{job_id}",
								"segments": []any{
									map[string]any{
										"lit": "analytics",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"analytics",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"job_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "job_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "0000-XXxx",
										},
									},
									"query": []any{
										map[string]any{
											"name": "app_id",
											"orig": "app_id",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
										},
										map[string]any{
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
											"example": "csv",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"app_id",
										"format",
										"id",
									},
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
