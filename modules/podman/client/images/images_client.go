// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new images API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for images API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	LibpodBuildImage(params *LibpodBuildImageParams) (*LibpodBuildImageOK, error)

	LibpodChangesImages(params *LibpodChangesImagesParams) (*LibpodChangesImagesOK, error)

	LibpodExportImage(params *LibpodExportImageParams, writer io.Writer) (*LibpodExportImageOK, error)

	LibpodExportImages(params *LibpodExportImagesParams, writer io.Writer) (*LibpodExportImagesOK, error)

	LibpodImageExists(params *LibpodImageExistsParams) (*LibpodImageExistsNoContent, error)

	LibpodImageHistory(params *LibpodImageHistoryParams) (*LibpodImageHistoryOK, error)

	LibpodImageTree(params *LibpodImageTreeParams) (*LibpodImageTreeOK, error)

	LibpodImagesImport(params *LibpodImagesImportParams) (*LibpodImagesImportOK, error)

	LibpodImagesLoad(params *LibpodImagesLoadParams) (*LibpodImagesLoadOK, error)

	LibpodImagesPull(params *LibpodImagesPullParams) (*LibpodImagesPullOK, error)

	LibpodImagesRemove(params *LibpodImagesRemoveParams) (*LibpodImagesRemoveOK, error)

	LibpodInspectImage(params *LibpodInspectImageParams) (*LibpodInspectImageOK, error)

	LibpodListImages(params *LibpodListImagesParams) (*LibpodListImagesOK, error)

	LibpodPruneImages(params *LibpodPruneImagesParams) (*LibpodPruneImagesOK, error)

	LibpodPushImage(params *LibpodPushImageParams, writer io.Writer) (*LibpodPushImageOK, error)

	LibpodRemoveImage(params *LibpodRemoveImageParams) (*LibpodRemoveImageOK, error)

	LibpodSearchImages(params *LibpodSearchImagesParams) (*LibpodSearchImagesOK, error)

	LibpodTagImage(params *LibpodTagImageParams) (*LibpodTagImageCreated, error)

	LibpodUntagImage(params *LibpodUntagImageParams) (*LibpodUntagImageCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  LibpodBuildImage creates image

  Build an image from the given Dockerfile(s)
*/
func (a *Client) LibpodBuildImage(params *LibpodBuildImageParams) (*LibpodBuildImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodBuildImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodBuildImage",
		Method:             "POST",
		PathPattern:        "/libpod/build",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodBuildImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodBuildImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodBuildImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodChangesImages reports on changes to images s filesystem adds deletes or modifications

  Returns which files in a images's filesystem have been added, deleted, or modified. The Kind of modification can be one of:

0: Modified
1: Added
2: Deleted

*/
func (a *Client) LibpodChangesImages(params *LibpodChangesImagesParams) (*LibpodChangesImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodChangesImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodChangesImages",
		Method:             "GET",
		PathPattern:        "/libpod/images/{name}/changes",
		ProducesMediaTypes: []string{"application/json", "text/html", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodChangesImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodChangesImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodChangesImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodExportImage exports an image

  Export an image
*/
func (a *Client) LibpodExportImage(params *LibpodExportImageParams, writer io.Writer) (*LibpodExportImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodExportImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodExportImage",
		Method:             "GET",
		PathPattern:        "/libpod/images/{name:.*}/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodExportImageReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodExportImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodExportImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodExportImages exports multiple images

  Export multiple images into a single object. Only `docker-archive` is currently supported.
*/
func (a *Client) LibpodExportImages(params *LibpodExportImagesParams, writer io.Writer) (*LibpodExportImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodExportImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodExportImages",
		Method:             "GET",
		PathPattern:        "/libpod/images/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodExportImagesReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodExportImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodExportImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodImageExists images exists

  Check if image exists in local store
*/
func (a *Client) LibpodImageExists(params *LibpodImageExistsParams) (*LibpodImageExistsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodImageExistsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodImageExists",
		Method:             "GET",
		PathPattern:        "/libpod/images/{name:.*}/exists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodImageExistsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodImageExistsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodImageExists: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodImageHistory histories of an image

  Return parent layers of an image.
*/
func (a *Client) LibpodImageHistory(params *LibpodImageHistoryParams) (*LibpodImageHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodImageHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodImageHistory",
		Method:             "GET",
		PathPattern:        "/libpod/images/{name:.*}/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodImageHistoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodImageHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodImageHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodImageTree images tree

  Retrieve the image tree for the provided image name or ID
*/
func (a *Client) LibpodImageTree(params *LibpodImageTreeParams) (*LibpodImageTreeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodImageTreeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodImageTree",
		Method:             "GET",
		PathPattern:        "/libpod/images/{name:.*}/tree",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodImageTreeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodImageTreeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodImageTree: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodImagesImport imports image

  Import a previously exported tarball as an image.
*/
func (a *Client) LibpodImagesImport(params *LibpodImagesImportParams) (*LibpodImagesImportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodImagesImportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodImagesImport",
		Method:             "POST",
		PathPattern:        "/libpod/images/import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodImagesImportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodImagesImportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodImagesImport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodImagesLoad loads image

  Load an image (oci-archive or docker-archive) stream.
*/
func (a *Client) LibpodImagesLoad(params *LibpodImagesLoadParams) (*LibpodImagesLoadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodImagesLoadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodImagesLoad",
		Method:             "POST",
		PathPattern:        "/libpod/images/load",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodImagesLoadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodImagesLoadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodImagesLoad: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodImagesPull pulls images

  Pull one or more images from a container registry.
*/
func (a *Client) LibpodImagesPull(params *LibpodImagesPullParams) (*LibpodImagesPullOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodImagesPullParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodImagesPull",
		Method:             "POST",
		PathPattern:        "/libpod/images/pull",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodImagesPullReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodImagesPullOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodImagesPull: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodImagesRemove removes one or more images from the storage

  Remove one or more images from the storage.
*/
func (a *Client) LibpodImagesRemove(params *LibpodImagesRemoveParams) (*LibpodImagesRemoveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodImagesRemoveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodImagesRemove",
		Method:             "DELETE",
		PathPattern:        "/libpod/images/remove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodImagesRemoveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodImagesRemoveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodImagesRemove: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodInspectImage inspects an image

  Obtain low-level information about an image
*/
func (a *Client) LibpodInspectImage(params *LibpodInspectImageParams) (*LibpodInspectImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodInspectImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodInspectImage",
		Method:             "GET",
		PathPattern:        "/libpod/images/{name:.*}/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodInspectImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodInspectImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodInspectImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodListImages lists images

  Returns a list of images on the server
*/
func (a *Client) LibpodListImages(params *LibpodListImagesParams) (*LibpodListImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodListImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodListImages",
		Method:             "GET",
		PathPattern:        "/libpod/images/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodListImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodListImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodListImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodPruneImages prunes unused images

  Remove images that are not being used by a container
*/
func (a *Client) LibpodPruneImages(params *LibpodPruneImagesParams) (*LibpodPruneImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodPruneImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodPruneImages",
		Method:             "POST",
		PathPattern:        "/libpod/images/prune",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodPruneImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodPruneImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodPruneImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodPushImage pushes image

  Push an image to a container registry
*/
func (a *Client) LibpodPushImage(params *LibpodPushImageParams, writer io.Writer) (*LibpodPushImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodPushImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodPushImage",
		Method:             "POST",
		PathPattern:        "/libpod/images/{name:.*}/push",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodPushImageReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodPushImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodPushImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodRemoveImage removes an image from the local storage

  Remove an image from the local storage.
*/
func (a *Client) LibpodRemoveImage(params *LibpodRemoveImageParams) (*LibpodRemoveImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodRemoveImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodRemoveImage",
		Method:             "DELETE",
		PathPattern:        "/libpod/images/{name:.*}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodRemoveImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodRemoveImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodRemoveImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodSearchImages searches images

  Search registries for images
*/
func (a *Client) LibpodSearchImages(params *LibpodSearchImagesParams) (*LibpodSearchImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodSearchImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodSearchImages",
		Method:             "GET",
		PathPattern:        "/libpod/images/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodSearchImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodSearchImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodSearchImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodTagImage tags an image

  Tag an image so that it becomes part of a repository.
*/
func (a *Client) LibpodTagImage(params *LibpodTagImageParams) (*LibpodTagImageCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodTagImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodTagImage",
		Method:             "POST",
		PathPattern:        "/libpod/images/{name:.*}/tag",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodTagImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodTagImageCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodTagImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LibpodUntagImage untags an image

  Untag an image. If not repo and tag are specified, all tags are removed from the image.
*/
func (a *Client) LibpodUntagImage(params *LibpodUntagImageParams) (*LibpodUntagImageCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLibpodUntagImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "libpodUntagImage",
		Method:             "POST",
		PathPattern:        "/libpod/images/{name:.*}/untag",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LibpodUntagImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LibpodUntagImageCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for libpodUntagImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
