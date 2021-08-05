// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package beta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func (r *Brand) validate() error {

	return nil
}

func brandGetURL(userBasePath string, r *Brand) (string, error) {
	params := map[string]interface{}{
		"project": dcl.ValueOrEmptyString(r.Project),
		"name":    dcl.ValueOrEmptyString(r.Name),
	}
	return dcl.URL("projects/{{project}}/brands/{{name}}", "https://iap.googleapis.com/v1/", userBasePath, params), nil
}

func brandListURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/brands", "https://iap.googleapis.com/v1/", userBasePath, params), nil

}

func brandCreateURL(userBasePath, project string) (string, error) {
	params := map[string]interface{}{
		"project": project,
	}
	return dcl.URL("projects/{{project}}/brands", "https://iap.googleapis.com/v1/", userBasePath, params), nil

}

// brandApiOperation represents a mutable operation in the underlying REST
// API such as Create, Update, or Delete.
type brandApiOperation interface {
	do(context.Context, *Brand, *Client) error
}

func (c *Client) listBrandRaw(ctx context.Context, project, pageToken string, pageSize int32) ([]byte, error) {
	u, err := brandListURL(c.Config.BasePath, project)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if pageToken != "" {
		m["pageToken"] = pageToken
	}

	if pageSize != BrandMaxPage {
		m["pageSize"] = fmt.Sprintf("%v", pageSize)
	}

	u, err = dcl.AddQueryParams(u, m)
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	return ioutil.ReadAll(resp.Response.Body)
}

type listBrandOperation struct {
	Brands []map[string]interface{} `json:"brands"`
	Token  string                   `json:"nextPageToken"`
}

func (c *Client) listBrand(ctx context.Context, project, pageToken string, pageSize int32) ([]*Brand, string, error) {
	b, err := c.listBrandRaw(ctx, project, pageToken, pageSize)
	if err != nil {
		return nil, "", err
	}

	var m listBrandOperation
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, "", err
	}

	var l []*Brand
	for _, v := range m.Brands {
		res, err := unmarshalMapBrand(v, c)
		if err != nil {
			return nil, m.Token, err
		}
		res.Project = &project
		l = append(l, res)
	}

	return l, m.Token, nil
}

// Create operations are similar to Update operations, although they do not have
// specific request objects. The Create request object is the json encoding of
// the resource, which is modified by res.marshal to form the base request body.
type createBrandOperation struct {
	response map[string]interface{}
}

func (op *createBrandOperation) FirstResponse() (map[string]interface{}, bool) {
	return op.response, len(op.response) > 0
}

func (op *createBrandOperation) do(ctx context.Context, r *Brand, c *Client) error {
	c.Config.Logger.Infof("Attempting to create %v", r)

	project := r.createFields()
	u, err := brandCreateURL(c.Config.BasePath, project)

	if err != nil {
		return err
	}

	req, err := r.marshal(c)
	if err != nil {
		return err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "POST", u, bytes.NewBuffer(req), c.Config.RetryProvider)
	if err != nil {
		return err
	}

	o, err := dcl.ResponseBodyAsJSON(resp)
	if err != nil {
		return fmt.Errorf("error decoding response body into JSON: %w", err)
	}
	op.response = o

	// Include Name in URL substitution for initial GET request.
	name, ok := op.response["name"].(string)
	if !ok {
		return fmt.Errorf("expected name to be a string, was %T", name)
	}
	r.Name = &name

	if _, err := c.GetBrand(ctx, r.URLNormalized()); err != nil {
		c.Config.Logger.Warningf("get returned error: %v", err)
		return err
	}

	return nil
}

func (c *Client) getBrandRaw(ctx context.Context, r *Brand) ([]byte, error) {

	u, err := brandGetURL(c.Config.BasePath, r.URLNormalized())
	if err != nil {
		return nil, err
	}
	resp, err := dcl.SendRequest(ctx, c.Config, "GET", u, &bytes.Buffer{}, c.Config.RetryProvider)
	if err != nil {
		return nil, err
	}
	defer resp.Response.Body.Close()
	b, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) brandDiffsForRawDesired(ctx context.Context, rawDesired *Brand, opts ...dcl.ApplyOption) (initial, desired *Brand, diffs []*dcl.FieldDiff, err error) {
	c.Config.Logger.Info("Fetching initial state...")
	// First, let us see if the user provided a state hint.  If they did, we will start fetching based on that.
	var fetchState *Brand
	if sh := dcl.FetchStateHint(opts); sh != nil {
		if r, ok := sh.(*Brand); !ok {
			c.Config.Logger.Warningf("Initial state hint was of the wrong type; expected Brand, got %T", sh)
		} else {
			fetchState = r
		}
	}
	if fetchState == nil {
		fetchState = rawDesired
	}

	if fetchState.Name == nil {
		// We cannot perform a get because of lack of information. We have to assume
		// that this is being created for the first time.
		desired, err := canonicalizeBrandDesiredState(rawDesired, nil)
		return nil, desired, nil, err
	}
	// 1.2: Retrieval of raw initial state from API
	rawInitial, err := c.GetBrand(ctx, fetchState.URLNormalized())
	if rawInitial == nil {
		if !dcl.IsNotFound(err) {
			c.Config.Logger.Warningf("Failed to retrieve whether a Brand resource already exists: %s", err)
			return nil, nil, nil, fmt.Errorf("failed to retrieve Brand resource: %v", err)
		}
		c.Config.Logger.Info("Found that Brand resource did not exist.")
		// Perform canonicalization to pick up defaults.
		desired, err = canonicalizeBrandDesiredState(rawDesired, rawInitial)
		return nil, desired, nil, err
	}
	c.Config.Logger.Infof("Found initial state for Brand: %v", rawInitial)
	c.Config.Logger.Infof("Initial desired state for Brand: %v", rawDesired)

	// 1.3: Canonicalize raw initial state into initial state.
	initial, err = canonicalizeBrandInitialState(rawInitial, rawDesired)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized initial state for Brand: %v", initial)

	// 1.4: Canonicalize raw desired state into desired state.
	desired, err = canonicalizeBrandDesiredState(rawDesired, rawInitial, opts...)
	if err != nil {
		return nil, nil, nil, err
	}
	c.Config.Logger.Infof("Canonicalized desired state for Brand: %v", desired)

	// 2.1: Comparison of initial and desired state.
	diffs, err = diffBrand(c, desired, initial, opts...)
	return initial, desired, diffs, err
}

func canonicalizeBrandInitialState(rawInitial, rawDesired *Brand) (*Brand, error) {
	// TODO(magic-modules-eng): write canonicalizer once relevant traits are added.
	return rawInitial, nil
}

/*
* Canonicalizers
*
* These are responsible for converting either a user-specified config or a
* GCP API response to a standard format that can be used for difference checking.
* */

func canonicalizeBrandDesiredState(rawDesired, rawInitial *Brand, opts ...dcl.ApplyOption) (*Brand, error) {

	if rawInitial == nil {
		// Since the initial state is empty, the desired state is all we have.
		// We canonicalize the remaining nested objects with nil to pick up defaults.

		return rawDesired, nil
	}
	canonicalDesired := &Brand{}
	if dcl.StringCanonicalize(rawDesired.ApplicationTitle, rawInitial.ApplicationTitle) {
		canonicalDesired.ApplicationTitle = rawInitial.ApplicationTitle
	} else {
		canonicalDesired.ApplicationTitle = rawDesired.ApplicationTitle
	}
	if dcl.IsZeroValue(rawDesired.Name) {
		canonicalDesired.Name = rawInitial.Name
	} else {
		canonicalDesired.Name = rawDesired.Name
	}
	if dcl.StringCanonicalize(rawDesired.SupportEmail, rawInitial.SupportEmail) {
		canonicalDesired.SupportEmail = rawInitial.SupportEmail
	} else {
		canonicalDesired.SupportEmail = rawDesired.SupportEmail
	}
	if dcl.NameToSelfLink(rawDesired.Project, rawInitial.Project) {
		canonicalDesired.Project = rawInitial.Project
	} else {
		canonicalDesired.Project = rawDesired.Project
	}

	return canonicalDesired, nil
}

func canonicalizeBrandNewState(c *Client, rawNew, rawDesired *Brand) (*Brand, error) {

	if dcl.IsEmptyValueIndirect(rawNew.ApplicationTitle) && dcl.IsEmptyValueIndirect(rawDesired.ApplicationTitle) {
		rawNew.ApplicationTitle = rawDesired.ApplicationTitle
	} else {
		if dcl.StringCanonicalize(rawDesired.ApplicationTitle, rawNew.ApplicationTitle) {
			rawNew.ApplicationTitle = rawDesired.ApplicationTitle
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.Name) && dcl.IsEmptyValueIndirect(rawDesired.Name) {
		rawNew.Name = rawDesired.Name
	} else {
	}

	if dcl.IsEmptyValueIndirect(rawNew.OrgInternalOnly) && dcl.IsEmptyValueIndirect(rawDesired.OrgInternalOnly) {
		rawNew.OrgInternalOnly = rawDesired.OrgInternalOnly
	} else {
		if dcl.BoolCanonicalize(rawDesired.OrgInternalOnly, rawNew.OrgInternalOnly) {
			rawNew.OrgInternalOnly = rawDesired.OrgInternalOnly
		}
	}

	if dcl.IsEmptyValueIndirect(rawNew.SupportEmail) && dcl.IsEmptyValueIndirect(rawDesired.SupportEmail) {
		rawNew.SupportEmail = rawDesired.SupportEmail
	} else {
		if dcl.StringCanonicalize(rawDesired.SupportEmail, rawNew.SupportEmail) {
			rawNew.SupportEmail = rawDesired.SupportEmail
		}
	}

	rawNew.Project = rawDesired.Project

	return rawNew, nil
}

// The differ returns a list of diffs, along with a list of operations that should be taken
// to remedy them. Right now, it does not attempt to consolidate operations - if several
// fields can be fixed with a patch update, it will perform the patch several times.
// Diffs on some fields will be ignored if the `desired` state has an empty (nil)
// value. This empty value indicates that the user does not care about the state for
// the field. Empty fields on the actual object will cause diffs.
// TODO(magic-modules-eng): for efficiency in some resources, add batching.
func diffBrand(c *Client, desired, actual *Brand, opts ...dcl.ApplyOption) ([]*dcl.FieldDiff, error) {
	if desired == nil || actual == nil {
		return nil, fmt.Errorf("nil resource passed to diff - always a programming error: %#v, %#v", desired, actual)
	}

	var fn dcl.FieldName
	var newDiffs []*dcl.FieldDiff
	// New style diffs.
	if ds, err := dcl.Diff(desired.ApplicationTitle, actual.ApplicationTitle, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("ApplicationTitle")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Name, actual.Name, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Name")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.OrgInternalOnly, actual.OrgInternalOnly, dcl.Info{OutputOnly: true, OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("OrgInternalOnly")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.SupportEmail, actual.SupportEmail, dcl.Info{OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("SupportEmail")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	if ds, err := dcl.Diff(desired.Project, actual.Project, dcl.Info{Type: "ReferenceType", OperationSelector: dcl.RequiresRecreate()}, fn.AddNest("Project")); len(ds) != 0 || err != nil {
		if err != nil {
			return nil, err
		}
		newDiffs = append(newDiffs, ds...)
	}

	return newDiffs, nil
}

func (r *Brand) getFields() (string, string) {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project), dcl.ValueOrEmptyString(n.Name)
}

func (r *Brand) createFields() string {
	n := r.URLNormalized()
	return dcl.ValueOrEmptyString(n.Project)
}

func (r *Brand) updateURL(userBasePath, updateName string) (string, error) {
	return "", fmt.Errorf("unknown update name: %s", updateName)
}

// marshal encodes the Brand resource into JSON for a Create request, and
// performs transformations from the resource schema to the API schema if
// necessary.
func (r *Brand) marshal(c *Client) ([]byte, error) {
	m, err := expandBrand(c, r)
	if err != nil {
		return nil, fmt.Errorf("error marshalling Brand: %w", err)
	}

	return json.Marshal(m)
}

// unmarshalBrand decodes JSON responses into the Brand resource schema.
func unmarshalBrand(b []byte, c *Client) (*Brand, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return unmarshalMapBrand(m, c)
}

func unmarshalMapBrand(m map[string]interface{}, c *Client) (*Brand, error) {

	return flattenBrand(c, m), nil
}

// expandBrand expands Brand into a JSON request object.
func expandBrand(c *Client, f *Brand) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if v := f.ApplicationTitle; !dcl.IsEmptyValueIndirect(v) {
		m["applicationTitle"] = v
	}
	if v, err := dcl.DeriveField("projects/%s/brands/%s", f.Name, f.Project, f.Name); err != nil {
		return nil, fmt.Errorf("error expanding Name into name: %w", err)
	} else if v != nil {
		m["name"] = v
	}
	if v := f.OrgInternalOnly; !dcl.IsEmptyValueIndirect(v) {
		m["orgInternalOnly"] = v
	}
	if v := f.SupportEmail; !dcl.IsEmptyValueIndirect(v) {
		m["supportEmail"] = v
	}
	if v, err := dcl.EmptyValue(); err != nil {
		return nil, fmt.Errorf("error expanding Project into project: %w", err)
	} else if v != nil {
		m["project"] = v
	}

	return m, nil
}

// flattenBrand flattens Brand from a JSON request object into the
// Brand type.
func flattenBrand(c *Client, i interface{}) *Brand {
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil
	}
	if len(m) == 0 {
		return nil
	}

	res := &Brand{}
	res.ApplicationTitle = dcl.FlattenString(m["applicationTitle"])
	res.Name = dcl.SelfLinkToName(dcl.FlattenString(m["name"]))
	res.OrgInternalOnly = dcl.FlattenBool(m["orgInternalOnly"])
	res.SupportEmail = dcl.FlattenString(m["supportEmail"])
	res.Project = dcl.FlattenString(m["project"])

	return res
}

// This function returns a matcher that checks whether a serialized resource matches this resource
// in its parameters (as defined by the fields in a Get, which definitionally define resource
// identity).  This is useful in extracting the element from a List call.
func (r *Brand) matcher(c *Client) func([]byte) bool {
	return func(b []byte) bool {
		cr, err := unmarshalBrand(b, c)
		if err != nil {
			c.Config.Logger.Warning("failed to unmarshal provided resource in matcher.")
			return false
		}
		nr := r.URLNormalized()
		ncr := cr.URLNormalized()
		c.Config.Logger.Infof("looking for %v\nin %v", nr, ncr)

		if nr.Project == nil && ncr.Project == nil {
			c.Config.Logger.Info("Both Project fields null - considering equal.")
		} else if nr.Project == nil || ncr.Project == nil {
			c.Config.Logger.Info("Only one Project field is null - considering unequal.")
			return false
		} else if *nr.Project != *ncr.Project {
			return false
		}
		if nr.Name == nil && ncr.Name == nil {
			c.Config.Logger.Info("Both Name fields null - considering equal.")
		} else if nr.Name == nil || ncr.Name == nil {
			c.Config.Logger.Info("Only one Name field is null - considering unequal.")
			return false
		} else if *nr.Name != *ncr.Name {
			return false
		}
		return true
	}
}

type brandDiff struct {
	// The diff should include one or the other of RequiresRecreate or UpdateOp.
	RequiresRecreate bool
	UpdateOp         brandApiOperation
}

func convertFieldDiffsToBrandDiffs(config *dcl.Config, fds []*dcl.FieldDiff, opts []dcl.ApplyOption) ([]brandDiff, error) {
	opNamesToFieldDiffs := make(map[string][]*dcl.FieldDiff)
	// Map each operation name to the field diffs associated with it.
	for _, fd := range fds {
		for _, ro := range fd.ResultingOperation {
			if fieldDiffs, ok := opNamesToFieldDiffs[ro]; ok {
				fieldDiffs = append(fieldDiffs, fd)
				opNamesToFieldDiffs[ro] = fieldDiffs
			} else {
				config.Logger.Infof("%s required due to diff in %q", ro, fd.FieldName)
				opNamesToFieldDiffs[ro] = []*dcl.FieldDiff{fd}
			}
		}
	}
	var diffs []brandDiff
	// For each operation name, create a brandDiff which contains the operation.
	for opName, fieldDiffs := range opNamesToFieldDiffs {
		diff := brandDiff{}
		if opName == "Recreate" {
			diff.RequiresRecreate = true
		} else {
			apiOp, err := convertOpNameToBrandApiOperation(opName, fieldDiffs, opts...)
			if err != nil {
				return diffs, err
			}
			diff.UpdateOp = apiOp
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func convertOpNameToBrandApiOperation(opName string, fieldDiffs []*dcl.FieldDiff, opts ...dcl.ApplyOption) (brandApiOperation, error) {
	switch opName {

	default:
		return nil, fmt.Errorf("no such operation with name: %v", opName)
	}
}
