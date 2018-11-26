// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package productsearch

// [START vision_product_search_create_product]

import (
	"context"
	"fmt"
	"io"

	vision "cloud.google.com/go/vision/apiv1"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

func createProduct(w io.Writer, projectId string, location string, productId string, productDisplayName string, productCategory string) error {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		return err
	}

	req := &visionpb.CreateProductRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s", projectId, location),
		ProductId: productId,
		Product: &visionpb.Product{
			DisplayName: productDisplayName,
			ProductCategory: productCategory,
		},
	}

	resp, err := c.CreateProduct(ctx, req)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, "Product name:", resp.Name)

	return nil
}

// [END vision_product_search_create_product]
