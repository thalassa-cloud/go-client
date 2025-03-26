package iaas

import (
	"context"
	"fmt"

	"github.com/thalassa-cloud/client-go/pkg/client"
)

const (
	SubnetEndpoint = "/v1/subnets"
)

// ListSubnets lists all Subnets for a given organisation.
func (c *Client) ListSubnets(ctx context.Context) ([]Subnet, error) {
	subnets := []Subnet{}
	req := c.R().SetResult(&subnets)

	resp, err := c.Do(ctx, req, client.GET, SubnetEndpoint)
	if err != nil {
		return nil, err
	}

	if err := c.Check(resp); err != nil {
		return subnets, err
	}
	return subnets, nil
}

// GetSubnet retrieves a specific Subnet by its identity.
func (c *Client) GetSubnet(ctx context.Context, identity string) (*Subnet, error) {
	var subnet *Subnet
	req := c.R().SetResult(&subnet)
	resp, err := c.Do(ctx, req, client.GET, fmt.Sprintf("%s/%s", SubnetEndpoint, identity))
	if err != nil {
		return nil, err
	}
	if err := c.Check(resp); err != nil {
		return subnet, err
	}
	return subnet, nil
}

// CreateSubnet creates a new Subnet.
func (c *Client) CreateSubnet(ctx context.Context, create CreateSubnet) (*Subnet, error) {
	var subnet *Subnet
	req := c.R().
		SetBody(create).SetResult(&subnet)

	resp, err := c.Do(ctx, req, client.POST, SubnetEndpoint)
	if err != nil {
		return nil, err
	}
	if err := c.Check(resp); err != nil {
		return subnet, err
	}
	return subnet, nil
}

// UpdateSubnet updates an existing Subnet.
func (c *Client) UpdateSubnet(ctx context.Context, identity string, update UpdateSubnet) (*Subnet, error) {
	var subnet *Subnet
	req := c.R().
		SetBody(update).SetResult(&subnet)

	resp, err := c.Do(ctx, req, client.PUT, fmt.Sprintf("%s/%s", SubnetEndpoint, identity))
	if err != nil {
		return nil, err
	}
	if err := c.Check(resp); err != nil {
		return subnet, err
	}
	return subnet, nil
}

// DeleteSubnet deletes a specific Subnet by its identity.
func (c *Client) DeleteSubnet(ctx context.Context, identity string) error {
	req := c.R()

	resp, err := c.Do(ctx, req, client.DELETE, fmt.Sprintf("%s/%s", SubnetEndpoint, identity))
	if err != nil {
		return err
	}
	if err := c.Check(resp); err != nil {
		return err
	}
	return nil
}
