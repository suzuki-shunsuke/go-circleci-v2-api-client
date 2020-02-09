package client

import (
	"context"
	"errors"
	"net/http"
	"net/url"
)

type (
	InsightWorkflow struct {
		Items []InsightWorkflowItem `json:"items"`
	}

	InsightWorkflowItem struct {
		ID          string `json:"id"`
		Status      string `json:"status"`
		Duration    int    `json:"duration"`
		CreatedAt   string `json:"created_at"`
		StoppedAt   string `json:"stopped_at"`
		CreditsUsed int    `json:"credits_used"`
	}

	InsightWorkflows struct {
		// NextPageToken string `json:"next_page_token"`
		Items []InsightWorkflowsItem `json:"items"`
	}

	InsightWorkflowsItem struct {
		Name        string                  `json:"name"`
		Metrics     InsightWorkflowsMetrics `json:"metrics"`
		WindowStart string                  `json:"window_start"`
		WindowEnd   string                  `json:"window_end"`
	}

	InsightWorkflowsMetrics struct {
		SuccessRate      float64                         `json:"success_rate"`
		Throughput       float64                         `json:"throughput"`
		TotalRuns        int                             `json:"total_runs"`
		FailedRuns       int                             `json:"failed_runs"`
		SuccessfulRuns   int                             `json:"successful_runs"`
		Mttr             int                             `json:"mttr"`
		TotalCreditsUsed int                             `json:"total_credits_used"`
		DurationMetrics  InsightWorkflowsDurationMetrics `json:"duration_metrics"`
	}

	InsightWorkflowsDurationMetrics struct {
		Min               int     `json:"min"`
		Max               int     `json:"max"`
		Median            int     `json:"median"`
		Mean              int     `json:"mean"`
		P95               int     `json:"p95"`
		StandardDeviation float64 `json:"standard_deviation"`
	}
)

func (c *Client) GetInsightWorkflowResp(ctx context.Context, projectSlug, workflow, branch string) (*http.Response, error) {
	// GET /insights/:project-slug/workflows/:workflow-name?branch=":branch-name"
	query := url.Values{}
	if workflow == "" {
		return nil, errors.New("workflow name is required")
	}
	if branch != "" {
		query.Add("branch", branch)
	}
	return c.getResp(ctx, "GET", "/insights/"+projectSlug+"/workflows/"+workflow, nil, query)
}

// The returned response body is closed.
func (c *Client) GetInsightWorkflow(
	ctx context.Context, projectSlug, workflow, branch string,
) (*InsightWorkflow, *http.Response, error) {
	// GET /insights/:project-slug/workflows/:workflow-name?branch=":branch-name"
	resp, err := c.GetInsightWorkflowResp(ctx, projectSlug, workflow, branch)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	body := &InsightWorkflow{}
	return body, resp, c.parseResponse(resp, body)
}

func (c *Client) GetInsightWorkflowsResp(ctx context.Context, projectSlug, branch string) (*http.Response, error) {
	// GET /insights/:project-slug/workflows?branch=":branch-name"
	query := url.Values{}
	if branch != "" {
		query.Add("branch", branch)
	}
	return c.getResp(ctx, "GET", "/insights/"+projectSlug+"/workflows", nil, query)
}

// The returned response body is closed.
func (c *Client) GetInsightWorkflows(
	ctx context.Context, projectSlug, branch string,
) (*InsightWorkflows, *http.Response, error) {
	// GET /insights/:project-slug/workflows?branch=":branch-name"
	resp, err := c.GetInsightWorkflowsResp(ctx, projectSlug, branch)
	if err != nil {
		return nil, resp, err
	}
	defer resp.Body.Close()
	body := &InsightWorkflows{}
	return body, resp, c.parseResponse(resp, body)
}
