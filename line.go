package tfl

import (
	"context"
	"fmt"

	"github.com/jamesalexatkin/tfl-golang/internal/conv"
)

// GetLineStatusByMode gets the line status of all lines for the given modes.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Line/Line_StatusByMode
func (c *Client) GetLineStatusByMode(ctx context.Context, modes []string) ([]Status, error) {
	path := fmt.Sprintf("/Line/Mode/%s/Status", conv.StringSliceToString(modes))

	statuses := []Status{}
	err := c.get(ctx, path, &statuses)
	if err != nil {
		return nil, err
	}

	return statuses, err
}

package tfl

import (
	"context"
	"fmt"

	"github.com/jamesalexatkin/tfl-golang/internal/conv"
)

// GetLineStatusByMode gets the line status of all lines for the given modes.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Line/Line_StatusByMode
func (c *Client) GetLineStatusByMode(ctx context.Context, modes []string) ([]Status, error) {
	path := fmt.Sprintf("/Line/Mode/%s/Status", conv.StringSliceToString(modes))

	statuses := []Status{}
	err := c.get(ctx, path, &statuses)
	if err != nil {
		return nil, err
	}

	return statuses, err
}

// GetLineArrivals gets the list of arrival predictions for given line ids based at the given stop.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Line/Line_Arrivals
func (c *Client) GetLineArrivals(ctx context.Context, lines []string, stopid string) ([]Prediction, error) {
	path := fmt.Sprintf("/Line/%s/Arrivals/%s", conv.StringSliceToString(lines), stopid)

	predictions := []Prediction{}
	err := c.get(ctx, path, &predictions)
	if err != nil {
		return nil, err
	}

	return predictions, err
}
