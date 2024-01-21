package sdk

import "context"

type policyReference struct {
	client *Client
}

func (v *policyReference) GetForEntity(ctx context.Context, request *GetForEntityPolicyReferenceRequest) ([]PolicyReference, error) {
	opts := request.toOpts()
	dbRows, err := validateAndQuery[policyReferenceDBRow](v.client, ctx, opts)
	if err != nil {
		return nil, err
	}

	resultList := convertRows[policyReferenceDBRow, PolicyReference](dbRows)

	return resultList, nil
}

func (request *GetForEntityPolicyReferenceRequest) toOpts() *getForEntityPolicyReferenceOptions {
	return &getForEntityPolicyReferenceOptions{
		select_:       true,
		asterisk:      true,
		from:          true,
		tableFunction: true,
		arguments: []policyReferenceFunctionArgument{
			{Key: "ref_entity_name", Value: request.RefEntityName},
			{Key: "ref_entity_domain", Value: request.RefEntityDomain},
		},
		endTableFunction: true,
	}
}