package client

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
)

func TestClient_GetInsightWorkflow(t *testing.T) {
	ctx := context.Background()
	cl := NewClient("xxx")

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "https://circleci.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "GET",
								Path:   "/api/v2/insights/gh/suzuki-shunsuke/go-circleci-v2-api-client/workflows/test",
								Query: url.Values{
									"circle-token": []string{"xxx"},
									"branch":       []string{"develop"},
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: `{
  "items": [
    {
      "id": "00000000-0000-0000-0000-000000000000",
      "status": "success",
      "duration": 783,
      "created_at": "2020-02-01T13:35:55.854Z",
      "stopped_at": "2020-02-01T13:48:59.316Z",
      "credits_used": 2000
    },
    {
      "id": "00000000-0000-0000-0000-000000000001",
      "status": "success",
      "duration": 703,
      "created_at": "2020-02-01T13:35:22.154Z",
      "stopped_at": "2020-02-01T13:47:05.458Z",
      "credits_used": 1900
    }
  ]
}`,
							},
						},
					},
				},
			},
		},
	})
	body, resp, err := cl.GetInsightWorkflow(ctx, "gh/suzuki-shunsuke/go-circleci-v2-api-client", "test", "develop")
	require.Nil(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, &InsightWorkflow{
		Items: []InsightWorkflowItem{
			{
				ID:          "00000000-0000-0000-0000-000000000000",
				Status:      "success",
				Duration:    783,
				CreatedAt:   "2020-02-01T13:35:55.854Z",
				StoppedAt:   "2020-02-01T13:48:59.316Z",
				CreditsUsed: 2000,
			},
			{
				ID:          "00000000-0000-0000-0000-000000000001",
				Status:      "success",
				Duration:    703,
				CreatedAt:   "2020-02-01T13:35:22.154Z",
				StoppedAt:   "2020-02-01T13:47:05.458Z",
				CreditsUsed: 1900,
			},
		},
	}, body)
}

func TestClient_GetInsightWorkflows(t *testing.T) {
	ctx := context.Background()
	cl := NewClient("xxx")

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "https://circleci.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "GET",
								Path:   "/api/v2/insights/gh/suzuki-shunsuke/go-circleci-v2-api-client/workflows",
								Query: url.Values{
									"circle-token": []string{"xxx"},
									"branch":       []string{"develop"},
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: `{
  "next_page_token": null,
  "items": [
    {
      "name": "test",
      "metrics": {
        "success_rate": 1,
        "total_runs": 3,
        "failed_runs": 0,
        "successful_runs": 3,
        "throughput": 1,
        "mttr": 0,
        "duration_metrics": {
          "min": 180,
          "max": 220,
          "median": 200,
          "mean": 200,
          "p95": 215,
          "standard_deviation": 20
        },
        "total_credits_used": 90
      },
      "window_start": "2020-01-31T00:33:01.964Z",
      "window_end": "2020-02-02T00:36:25.999Z"
    },
    {
      "name": "deploy",
      "metrics": {
        "success_rate": 1,
        "total_runs": 45,
        "failed_runs": 0,
        "successful_runs": 45,
        "throughput": 15,
        "mttr": 0,
        "duration_metrics": {
          "min": 180,
          "max": 500,
          "median": 255,
          "mean": 275,
          "p95": 450,
          "standard_deviation": 71
        },
        "total_credits_used": 4000
      },
      "window_start": "2020-01-30T09:33:59.482Z",
      "window_end": "2020-02-01T13:39:13.723Z"
    }
  ]
}`,
							},
						},
					},
				},
			},
		},
	})
	body, resp, err := cl.GetInsightWorkflows(ctx, "gh/suzuki-shunsuke/go-circleci-v2-api-client", "develop")
	require.Nil(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, &InsightWorkflows{
		Items: []InsightWorkflowsItem{
			{
				Name: "test",
				Metrics: InsightWorkflowsMetrics{
					SuccessRate:    1,
					TotalRuns:      3,
					FailedRuns:     0,
					SuccessfulRuns: 3,
					Throughput:     1,
					Mttr:           0,
					DurationMetrics: InsightWorkflowsDurationMetrics{
						Min:               180,
						Max:               220,
						Median:            200,
						Mean:              200,
						P95:               215,
						StandardDeviation: 20,
					},
					TotalCreditsUsed: 90,
				},
				WindowStart: "2020-01-31T00:33:01.964Z",
				WindowEnd:   "2020-02-02T00:36:25.999Z",
			},
			{
				Name: "deploy",
				Metrics: InsightWorkflowsMetrics{
					SuccessRate:    1,
					TotalRuns:      45,
					FailedRuns:     0,
					SuccessfulRuns: 45,
					Throughput:     15,
					Mttr:           0,
					DurationMetrics: InsightWorkflowsDurationMetrics{
						Min:               180,
						Max:               500,
						Median:            255,
						Mean:              275,
						P95:               450,
						StandardDeviation: 71,
					},
					TotalCreditsUsed: 4000,
				},
				WindowStart: "2020-01-30T09:33:59.482Z",
				WindowEnd:   "2020-02-01T13:39:13.723Z",
			},
		},
	}, body)
}
