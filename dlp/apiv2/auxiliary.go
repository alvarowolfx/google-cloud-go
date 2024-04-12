// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package dlp

import (
	dlppb "cloud.google.com/go/dlp/apiv2/dlppb"
	"google.golang.org/api/iterator"
)

// ColumnDataProfileIterator manages a stream of *dlppb.ColumnDataProfile.
type ColumnDataProfileIterator struct {
	items    []*dlppb.ColumnDataProfile
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dlppb.ColumnDataProfile, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ColumnDataProfileIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ColumnDataProfileIterator) Next() (*dlppb.ColumnDataProfile, error) {
	var item *dlppb.ColumnDataProfile
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ColumnDataProfileIterator) bufLen() int {
	return len(it.items)
}

func (it *ColumnDataProfileIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// DeidentifyTemplateIterator manages a stream of *dlppb.DeidentifyTemplate.
type DeidentifyTemplateIterator struct {
	items    []*dlppb.DeidentifyTemplate
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dlppb.DeidentifyTemplate, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *DeidentifyTemplateIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *DeidentifyTemplateIterator) Next() (*dlppb.DeidentifyTemplate, error) {
	var item *dlppb.DeidentifyTemplate
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *DeidentifyTemplateIterator) bufLen() int {
	return len(it.items)
}

func (it *DeidentifyTemplateIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// DiscoveryConfigIterator manages a stream of *dlppb.DiscoveryConfig.
type DiscoveryConfigIterator struct {
	items    []*dlppb.DiscoveryConfig
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dlppb.DiscoveryConfig, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *DiscoveryConfigIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *DiscoveryConfigIterator) Next() (*dlppb.DiscoveryConfig, error) {
	var item *dlppb.DiscoveryConfig
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *DiscoveryConfigIterator) bufLen() int {
	return len(it.items)
}

func (it *DiscoveryConfigIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// DlpJobIterator manages a stream of *dlppb.DlpJob.
type DlpJobIterator struct {
	items    []*dlppb.DlpJob
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dlppb.DlpJob, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *DlpJobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *DlpJobIterator) Next() (*dlppb.DlpJob, error) {
	var item *dlppb.DlpJob
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *DlpJobIterator) bufLen() int {
	return len(it.items)
}

func (it *DlpJobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// InspectTemplateIterator manages a stream of *dlppb.InspectTemplate.
type InspectTemplateIterator struct {
	items    []*dlppb.InspectTemplate
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dlppb.InspectTemplate, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *InspectTemplateIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *InspectTemplateIterator) Next() (*dlppb.InspectTemplate, error) {
	var item *dlppb.InspectTemplate
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *InspectTemplateIterator) bufLen() int {
	return len(it.items)
}

func (it *InspectTemplateIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// JobTriggerIterator manages a stream of *dlppb.JobTrigger.
type JobTriggerIterator struct {
	items    []*dlppb.JobTrigger
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dlppb.JobTrigger, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *JobTriggerIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *JobTriggerIterator) Next() (*dlppb.JobTrigger, error) {
	var item *dlppb.JobTrigger
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *JobTriggerIterator) bufLen() int {
	return len(it.items)
}

func (it *JobTriggerIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ProjectDataProfileIterator manages a stream of *dlppb.ProjectDataProfile.
type ProjectDataProfileIterator struct {
	items    []*dlppb.ProjectDataProfile
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dlppb.ProjectDataProfile, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ProjectDataProfileIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ProjectDataProfileIterator) Next() (*dlppb.ProjectDataProfile, error) {
	var item *dlppb.ProjectDataProfile
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ProjectDataProfileIterator) bufLen() int {
	return len(it.items)
}

func (it *ProjectDataProfileIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// StoredInfoTypeIterator manages a stream of *dlppb.StoredInfoType.
type StoredInfoTypeIterator struct {
	items    []*dlppb.StoredInfoType
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dlppb.StoredInfoType, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *StoredInfoTypeIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *StoredInfoTypeIterator) Next() (*dlppb.StoredInfoType, error) {
	var item *dlppb.StoredInfoType
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *StoredInfoTypeIterator) bufLen() int {
	return len(it.items)
}

func (it *StoredInfoTypeIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TableDataProfileIterator manages a stream of *dlppb.TableDataProfile.
type TableDataProfileIterator struct {
	items    []*dlppb.TableDataProfile
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*dlppb.TableDataProfile, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TableDataProfileIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TableDataProfileIterator) Next() (*dlppb.TableDataProfile, error) {
	var item *dlppb.TableDataProfile
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TableDataProfileIterator) bufLen() int {
	return len(it.items)
}

func (it *TableDataProfileIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
