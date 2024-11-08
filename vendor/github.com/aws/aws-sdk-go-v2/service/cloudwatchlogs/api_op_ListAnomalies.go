// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of anomalies that log anomaly detectors have found. For details
// about the structure format of each anomaly object that is returned, see the
// example in this section.
func (c *Client) ListAnomalies(ctx context.Context, params *ListAnomaliesInput, optFns ...func(*Options)) (*ListAnomaliesOutput, error) {
	if params == nil {
		params = &ListAnomaliesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnomalies", params, optFns, c.addOperationListAnomaliesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnomaliesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAnomaliesInput struct {

	// Use this to optionally limit the results to only the anomalies found by a
	// certain anomaly detector.
	AnomalyDetectorArn *string

	// The maximum number of items to return. If you don't specify a value, the
	// default maximum value of 50 items is used.
	Limit *int32

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// You can specify this parameter if you want to the operation to return only
	// anomalies that are currently either suppressed or unsuppressed.
	SuppressionState types.SuppressionState

	noSmithyDocumentSerde
}

type ListAnomaliesOutput struct {

	// An array of structures, where each structure contains information about one
	// anomaly that a log anomaly detector has found.
	Anomalies []types.Anomaly

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAnomaliesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAnomalies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAnomalies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAnomalies"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAnomalies(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListAnomaliesAPIClient is a client that implements the ListAnomalies operation.
type ListAnomaliesAPIClient interface {
	ListAnomalies(context.Context, *ListAnomaliesInput, ...func(*Options)) (*ListAnomaliesOutput, error)
}

var _ ListAnomaliesAPIClient = (*Client)(nil)

// ListAnomaliesPaginatorOptions is the paginator options for ListAnomalies
type ListAnomaliesPaginatorOptions struct {
	// The maximum number of items to return. If you don't specify a value, the
	// default maximum value of 50 items is used.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAnomaliesPaginator is a paginator for ListAnomalies
type ListAnomaliesPaginator struct {
	options   ListAnomaliesPaginatorOptions
	client    ListAnomaliesAPIClient
	params    *ListAnomaliesInput
	nextToken *string
	firstPage bool
}

// NewListAnomaliesPaginator returns a new ListAnomaliesPaginator
func NewListAnomaliesPaginator(client ListAnomaliesAPIClient, params *ListAnomaliesInput, optFns ...func(*ListAnomaliesPaginatorOptions)) *ListAnomaliesPaginator {
	if params == nil {
		params = &ListAnomaliesInput{}
	}

	options := ListAnomaliesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAnomaliesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAnomaliesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAnomalies page.
func (p *ListAnomaliesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAnomaliesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListAnomalies(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListAnomalies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAnomalies",
	}
}
